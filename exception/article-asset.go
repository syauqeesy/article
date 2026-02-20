package exception

import (
	"net/http"

	"ahmadsyauqi.dev/article/common"
)

var InvalidArticleAssetStatus = common.CreateException(http.StatusBadRequest, "Invalid Status")

var InvalidArticleAssetContentType = common.CreateException(http.StatusBadRequest, "Invalid Content Type")
