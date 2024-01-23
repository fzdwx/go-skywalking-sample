package service

import (
	"github.com/gin-gonic/gin"
	"github.com/redis/go-redis/v9"
	"go.uber.org/zap"
	"net/http"
)

var rdb *redis.Client

func init() {
	rdb = redis.NewClient(&redis.Options{
		Addr:     "localhost:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})
}

func HelloService(c *gin.Context, logger *zap.Logger) {
	logger.Info("hello from service" + c.Param("name"))
	rdb.Set(c, "hello", c.Param("name"), 0)
	http.Get("http://localhost:8000/helloworld/" + c.Param("name"))
}
