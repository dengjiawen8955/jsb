package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"jsb/handler"
	"jsb/middleware"
	"jsb/util/my_restful"
)

func main() {
}
//export start
func start(){
	engine := gin.Default()
	//中间件
	engine.Use(middleware.HeaderMiddlerware)
	engine.Use(cors.Default())
	//登录
	engine.POST("/jsb/login",handler.LoginHandler)
	engine.POST("/jsb/register",handler.RegisterHandler)
	//构建Websocket server
	wsServer := handler.NewWsServer()
	engine.GET("/jsb/send-message",wsServer.WsHandler)
	engine.GET("/jsb/test", func(ctx *gin.Context) {
		ctx.Writer.Write(my_restful.Ok("go测试"))
	})
	engine.Run("0.0.0.0:8080")
}

