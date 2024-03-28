/*
 * @Author: yujiajie
 * @Date: 2024-03-15 17:59:03
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-15 18:02:33
 * @FilePath: /stage/app/dao/user/user.go
 * @Description:
 */
package user

import (
	"fmt"
	"stage/internal/schema"
	"stage/sdk/core"
)

func GetUserInfo(id int) schema.AdminUser {
	var user schema.AdminUser
	db := core.App.GetDb("alimatch/modadmin")
	if res := db.Where(&schema.AdminUser{ID: id}).First(&user); res.Error != nil {
		fmt.Printf("search user failed, id: %d, err: %s\n", id, res.Error)
	}
	return user
}
