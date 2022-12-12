package router

import (
	"github.com/gin-gonic/gin"
	"github.com/web3ssc/ss-go-we-chat-gpt/handler"
)

func Load(g *gin.Engine) {
	common := g.Group("api/common")
	{
		common.GET("/checkHealth", handler.Health)
	}
}
