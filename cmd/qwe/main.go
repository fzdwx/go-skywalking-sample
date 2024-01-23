package main

import (
	_ "github.com/apache/skywalking-go"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"helloworld/cmd/qwe/service"
)

func main() {
	engine := gin.New()

	logger, err := zap.NewProduction()
	if err != nil {
		panic(err)
	}

	engine.GET("/helloworld/:name", func(context *gin.Context) {
		logger.Info("hello" + context.Param("name"))
		service.HelloService(context, logger)
	})

	engine.Run(":8888")
}
