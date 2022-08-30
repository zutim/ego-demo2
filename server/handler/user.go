package handler

import (
	"ego-demo2/internal/service"
	"ego-demo2/server/response"
	"fmt"
	"github.com/gin-gonic/gin"
)

type userHandler struct {
}

func GetUserHandler() *userHandler {
	return &userHandler{}
}

func (u *userHandler) Find(ctx *gin.Context) {
	res, err := service.GetUser().GetAll()
	if err != nil {
		fmt.Println(err)
	}
	response.SuccessRsp(ctx, res)
	return
}
