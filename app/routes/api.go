/*
 * @Author: yujiajie
 * @Date: 2024-03-13 09:41:27
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-15 18:09:24
 * @FilePath: /stage/app/routes/api.go
 * @Description:
 */
package routes

import (
	"stage/app/controller/user"

	"github.com/gin-gonic/gin"
)

func HandleApi(r *gin.Engine) {
	api := r.Group("/api")

	handleAuth(api)
	handleUser(api)
}

func handleAuth(api *gin.RouterGroup) {
	group := api.Group("/auth")
	group.POST("/login", nil)
}

func handleUser(api *gin.RouterGroup) {
	group := api.Group("/user")
	group.GET("/info", user.Info)
}
