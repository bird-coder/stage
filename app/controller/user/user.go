package user

import (
	"stage/app/controller"
	"stage/app/dao/user"

	"github.com/gin-gonic/gin"
)

func Info(c *gin.Context) {
	info := user.GetUserInfo(1)
	controller.OK(c, info, "success")
}
