package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	//创建一个服务
	ginServer := gin.Default()

	//访问地址，处理我们的请求 Request Response
	//gin.Context，封装了request和response
	ginServer.GET("/hello", func(context *gin.Context) {
		context.JSON(200, gin.H{
			"msg": "hello world",
		})
	})
	//服务器端口
	//Run("里面不指定端口号默认为8080")
	ginServer.Run(":8081")
}
