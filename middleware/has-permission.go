package middleware

import (
	"net/http"

	"ahmadsyauqi.dev/article/common"
	"ahmadsyauqi.dev/article/exception"
	"ahmadsyauqi.dev/article/repository"
)

func HasPermission(code string, accountPermissionRepository repository.AccountPermissionRepository) func(http.Handler) http.Handler {
	return func(next http.Handler) http.Handler {
		return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
			subject, ok := GetSubject(r.Context())
			if !ok {
				common.HttpErrorHandler(w, exception.Unauthorized, nil)
				return
			}

			accountPermission, err := accountPermissionRepository.FindByPermissionCodeAndAccountId(r.Context(), code, subject)
			if err != nil {
				common.HttpErrorHandler(w, exception.Unauthorized, nil)
				return
			}
			if accountPermission.Permission == nil {
				common.HttpErrorHandler(w, exception.Unauthorized, nil)
				return
			}

			next.ServeHTTP(w, r.WithContext(r.Context()))
		})
	}
}
