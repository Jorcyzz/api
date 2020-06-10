package util

import (
	"api/pkg/e"

	"github.com/gin-gonic/gin"
)

type Response struct {
	Code int
	Msg string
	Data interface{}
}

func ResponseWithJson(code int, data interface{}, c *gin.Context) {
	c.JSON(200,&Response{
		Code:code,
		Msg:e.GetMsg(code),
		Data: data,
	})
}