package handler

import (
	"github.com/gin-gonic/gin"
	"github.com/web3ssc/ss-go-common/errno"
	"net/http"
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"message"`
	Data interface{} `json:"data,omitempty"`
}

func SendResponse(c *gin.Context, err error, data interface{}) {
	code, msg := errno.DecodeErr(err)
	c.JSON(http.StatusOK, Response{
		Code: code,
		Msg:  msg,
		Data: data,
	})
}
