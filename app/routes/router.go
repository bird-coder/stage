/*
 * @Author: yujiajie
 * @Date: 2024-03-13 09:40:01
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-14 22:39:29
 * @FilePath: /stage/app/routes/router.go
 * @Description:
 */
package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func InitRouter(r *gin.Engine) {
	HandleApi(r)

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK, gin.H{
			"message": "pong",
		})
	})
}
