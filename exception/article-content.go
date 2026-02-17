package exception

import (
	"net/http"

	"ahmadsyauqi.dev/article/common"
)

var SlugAlreadyExists = common.CreateException(http.StatusBadRequest, "Slug Already Exists")
