package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/huynhdev/go-gin-structure/pkg/e"
	"github.com/huynhdev/go-gin-structure/pkg/logging"
)

func BindAndValid(c *gin.Context, form interface{}) (int, int) {
	if err := c.ShouldBindJSON(form); err != nil {
		logging.Info(err)
		return http.StatusBadRequest, e.INVALID_PARAMS
	}

	return http.StatusOK, e.SUCCESS
}
