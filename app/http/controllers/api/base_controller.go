// Package api 基层控制器
package api

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

// BaseApiController 基础控制器
type BaseApiController struct {
}

func (ctrl *BaseApiController) Welcome(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "hello goapi!",
	})
}
