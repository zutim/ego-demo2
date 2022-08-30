// route 使用单独的路由模块来管理中间件与路由
package route

import (
	"ego-demo2/server/handler"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/zutim/ego/component"
	"strings"
)

// Load
func Loader(router *gin.Engine) {
	userHandler := handler.GetUserHandler()
	v1 := router.Group("v1")
	//v1.Use(JWT())
	//{
	//
	//}
	v1.GET("index", handler.Index)
	v1.GET("index2", handler.Index2)
	v1.GET("users", userHandler.Find)
}

func JWT() gin.HandlerFunc {
	validator := func(ctx *gin.Context) error {
		token := strings.TrimPrefix(ctx.GetHeader("Authorization"), "Bearer ")
		if token == "" {
			return fmt.Errorf("invalid token")
		}

		claims, err := component.Provider().Jwt().ParseToken(token)
		if err != nil {
			return fmt.Errorf("parse token: %v", err)
		}

		// 令牌信息存入context
		ctx.Set("claimsKey", claims)

		return nil
	}

	return func(ctx *gin.Context) {
		// 解析token
		if err := validator(ctx); err != nil {
			//response.WrapContext(ctx).Error(401, err.Error())

			ctx.Abort()
			return
		}

		ctx.Next()
	}
}
