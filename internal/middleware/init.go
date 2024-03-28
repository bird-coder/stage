package middleware

import "github.com/gin-gonic/gin"

func Init(g *gin.Engine) {
	g.Use(NoCache)
	g.Use(Options)
	g.Use(Secure)

	g.Use(ErrorHandler())
}
