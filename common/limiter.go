package common

import (
	"net"
	"net/http"
	"strings"
	"sync"
	"time"

	"golang.org/x/time/rate"
)

type client struct {
	limiter  *rate.Limiter
	lastSeen time.Time
}

type IPRateLimiter struct {
	Mutex        sync.Mutex
	Clients      map[string]*client
	Rate         rate.Limit
	Burst        int
	Ttl          time.Duration
	TrustProxies bool
}

func NewIPRateLimiter(r rate.Limit, burst int, ttl time.Duration, trustProxies bool) *IPRateLimiter {
	rl := &IPRateLimiter{
		Clients:      make(map[string]*client),
		Rate:         r,
		Burst:        burst,
		Ttl:          ttl,
		TrustProxies: trustProxies,
	}

	go rl.cleanupLoop(1 * time.Minute)

	return rl
}

func (rl *IPRateLimiter) getLimiter(ip string) *rate.Limiter {
	now := time.Now()

	rl.Mutex.Lock()
	defer rl.Mutex.Unlock()

	if c, ok := rl.Clients[ip]; ok {
		c.lastSeen = now
		return c.limiter
	}

	lim := rate.NewLimiter(rl.Rate, rl.Burst)
	rl.Clients[ip] = &client{limiter: lim, lastSeen: now}
	return lim
}

func (rl *IPRateLimiter) cleanupLoop(interval time.Duration) {
	t := time.NewTicker(interval)
	defer t.Stop()

	for range t.C {
		cutoff := time.Now().Add(-rl.Ttl)

		rl.Mutex.Lock()
		for ip, c := range rl.Clients {
			if c.lastSeen.Before(cutoff) {
				delete(rl.Clients, ip)
			}
		}
		rl.Mutex.Unlock()
	}
}

func (rl *IPRateLimiter) Middleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ip := rl.clientIP(r)
		lim := rl.getLimiter(ip)

		if !lim.Allow() {
			w.Header().Set("Retry-After", "1")
			HttpErrorHandler(w, CreateException(http.StatusTooManyRequests, http.StatusText(http.StatusTooManyRequests)), nil)
			return
		}

		next.ServeHTTP(w, r)
	})
}

func (rl *IPRateLimiter) clientIP(r *http.Request) string {
	if rl.TrustProxies {
		if xff := r.Header.Get("X-Forwarded-For"); xff != "" {
			parts := strings.Split(xff, ",")
			if len(parts) > 0 {
				ip := strings.TrimSpace(parts[0])
				if net.ParseIP(ip) != nil {
					return ip
				}
			}
		}
		if xrip := strings.TrimSpace(r.Header.Get("X-Real-IP")); xrip != "" {
			if net.ParseIP(xrip) != nil {
				return xrip
			}
		}
	}

	host, _, err := net.SplitHostPort(r.RemoteAddr)
	if err == nil && net.ParseIP(host) != nil {
		return host
	}

	return r.RemoteAddr
}
