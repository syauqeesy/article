package exception

import (
	"net/http"

	"ahmadsyauqi.dev/article/common"
)

var InvalidArticleStatus = common.CreateException(http.StatusBadRequest, "Invalid Article Status")
var ArticleNotFound = common.CreateException(http.StatusNotFound, "Article Not Found")

