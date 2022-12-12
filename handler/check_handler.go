package handler

import (
	"github.com/gin-gonic/gin"
)

const (
	B  = 1
	KB = 1024 * B
	MB = 1024 * KB
	GB = 1024 * MB
)

func Health(c *gin.Context) {
	SendResponse(c, nil, nil)
}
