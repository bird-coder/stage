/*
 * @Author: yujiajie
 * @Date: 2024-03-13 09:50:14
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-13 09:56:19
 * @FilePath: /stage/internal/middleware/error.go
 * @Description:
 */
package middleware

import (
	"net/http"

	"github.com/bird-coder/manyo/constant"
	"github.com/gin-gonic/gin"
)

func ErrorHandler() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		defer func() {
			if err := recover(); err != nil {
				ctx.AbortWithStatusJSON(http.StatusInternalServerError, gin.H{
					"code": constant.SERVER_ERROR,
					"msg":  constant.SERVER_ERROR.String(),
					"data": nil,
				})
			}
		}()
		ctx.Next()
	}
}
