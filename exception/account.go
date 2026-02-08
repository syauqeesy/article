package exception

import (
	"net/http"

	"ahmadsyauqi.dev/article/common"
)

var InvalidOauthProvider = common.CreateException(http.StatusBadRequest, "invalid oauth provider")

var Unauthorized = common.CreateException(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))
