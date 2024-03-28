/*
 * @Author: yujiajie
 * @Date: 2024-02-28 17:05:52
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-12 10:11:18
 * @FilePath: /stage/app/controller/base.go
 * @Description:
 */
package controller

import (
	"net/http"

	"github.com/bird-coder/manyo/pkg/response"
	"github.com/gin-gonic/gin"
)

func Error(c *gin.Context, code int, msg string) {
	res := response.Error(code, msg)
	c.AbortWithStatusJSON(code, res)
}

func OK(c *gin.Context, data interface{}, msg string) {
	res := response.OK(data, msg)
	c.AbortWithStatusJSON(http.StatusOK, res)
}
