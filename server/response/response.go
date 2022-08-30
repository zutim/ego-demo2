package response

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type SuccessResponse struct {
	Code int         `json:"code"`
	Msg  string      `json:"msg"`
	Data interface{} `json:"data"`
}

func SuccessRsp(ctx *gin.Context, data interface{}) {
	res := SuccessResponse{
		Code: 0,
		Msg:  "success",
		Data: data,
	}
	ctx.JSON(http.StatusOK, res)
}
