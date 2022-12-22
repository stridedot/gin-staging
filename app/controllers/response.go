package controllers

import (
	"github.com/gin-gonic/gin"
	"go_code/gintest/app/services"
	"net/http"
	"time"
)

type Response struct {
	Code      int    `json:"code"`
	Msg       string `json:"msg"`
	Data      *Data  `json:"data"`
	Timestamp string `json:"timestamp"`
}

type Data struct {
	Meta *services.Meta `json:"meta"`
	Data interface{}    `json:"data"`
}

type empty struct {
}

type Meta struct {
	Total       int `json:"total"`
	PerPage     int `json:"per_page"`
	CurrentPage int `json:"current_page"`
	LastPage    int `json:"last_page"`
}

func ErrorRes(c *gin.Context, code int, msg interface{}) {
	if msg == nil {
		msg = MsgTender(code)
	}
	c.JSON(http.StatusOK, &Response{
		Code:      code,
		Msg:       msg.(string),
		Data:      &Data{},
		Timestamp: time.Now().Format(time.RFC3339),
	})
}

func SuccessRes(c *gin.Context, data *Data) {
	if data == nil {
		data = &Data{}
	} else if data.Data == nil {
		data = &Data{Data: []*empty{}}
	}

	c.JSON(http.StatusOK, &Response{
		Code:      CodeOk,
		Msg:       MsgTender(CodeOk),
		Data:      data,
		Timestamp: time.Now().Format(time.RFC3339),
	})
}
