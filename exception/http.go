package exception

import (
	"net/http"

	"ahmadsyauqi.dev/article/common"
)

var TooManyRequests = common.CreateException(http.StatusTooManyRequests, http.StatusText(http.StatusTooManyRequests))
