package handler

import (
	"ego-demo2/internal/service"
	"ego-demo2/server/response"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Res struct {
	Id   int
	Name string
}

func Index(ctx *gin.Context) {
	service.GetInstanceIndex()
	res := &Res{
		1,
		"nihao",
	}
	response.SuccessRsp(ctx, res)
	return
}

func Index2(ctx *gin.Context) {
	ctx.String(http.StatusOK, "home")
	return
}
