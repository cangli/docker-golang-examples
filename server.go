package main

import (
	"fmt"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"go.uber.org/zap"
	"net/http"
	"strconv"
)

var sugar *zap.SugaredLogger
var client *redis.Client
func hello(c *gin.Context){
	c.String(http.StatusOK, "Hello")
	sugar.Infow("hello", "key", "values")
}

func plus(c *gin.Context){
	result := client.Incr("test")
	c.String(http.StatusOK, strconv.FormatInt(result.Val(), 10))
	sugar.Infow("redis incr", "test", strconv.FormatInt(result.Val(), 10))
}

func main(){
	client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	cfg := zap.NewProductionConfig()
	cfg.OutputPaths = []string{"/logs/log.log"}
	logger, err := cfg.Build()
	fmt.Println(err)
	sugar = logger.Sugar()
	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.GET("/hello", hello)
	router.GET("/plus", plus)
	router.Run(":8080")
}

