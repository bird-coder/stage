/*
 * @Author: yujiajie
 * @Date: 2024-02-23 17:26:14
 * @LastEditors: yujiajie
 * @LastEditTime: 2024-03-12 10:05:13
 * @FilePath: /stage/internal/middleware/header.go
 * @Description:
 */
package middleware

import (
	"net/http"
	"time"

	"github.com/bird-coder/manyo/constant"
	"github.com/gin-gonic/gin"
)

func NoCache(c *gin.Context) {
	c.Header("Cache-Control", "no-cache, no-store, max-age=0, must-revalidate, value")
	c.Header("Expires", "Thu, 01 Jan 1970 00:00:00 GMT")
	c.Header("Last-Modified", time.Now().UTC().Format(http.TimeFormat))
	c.Next()
}

func Options(c *gin.Context) {
	if c.Request.Method != "OPTIONS" {
		c.Next()
	} else {
		c.Header(constant.AllowOrigin, constant.AllOrigins)
		c.Header(constant.AllowMethods, constant.Methods)
		c.Header(constant.AllowHeaders, constant.AllowHeadersVal)
		c.Header("Allow", "HEAD,GET,POST,PUT,PATCH,DELETE,OPTIONS")
		c.Header(constant.ContentType, constant.ApplicationJson)
		c.AbortWithStatus(200)
	}
}

func Secure(c *gin.Context) {
	c.Header(constant.AllowOrigin, constant.AllOrigins)
	c.Header("X-Content-Type-Options", "nosniff")
	c.Header("X-XSS-Protection", "1; mode=block")
	if c.Request.TLS != nil {
		c.Header("Strict-Transport-Security", "max-age=31536000")
	}
}
