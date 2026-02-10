package exception

import (
	"net/http"

	"ahmadsyauqi.dev/article/common"
)

var InvalidOauthProvider = common.CreateException(http.StatusBadRequest, "Invalid Oauth Provider")

var Unauthorized = common.CreateException(http.StatusUnauthorized, http.StatusText(http.StatusUnauthorized))

var AccountNotFound = common.CreateException(http.StatusNotFound, "Account Not Found")
