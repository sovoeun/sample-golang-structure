package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

const (
	Success = 0
	Fail    = 100
)

type Response struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func RespOk(c *gin.Context, msg string, data interface{}) {
	c.JSON(http.StatusOK, Response{
		Code: Success,
		Msg:  msg,
		Data: data,
	})
}

type ArticleListParam struct {
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}
