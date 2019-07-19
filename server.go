package main

import (
	"fmt"
	"github.com/gin-contrib/gzip"
	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis"
	"net/http"
	"strconv"
)

var client *redis.Client
func hello(c *gin.Context){
	c.String(http.StatusOK, "Hello")
}

func plus(c *gin.Context){
	result := client.Incr("test")
	c.String(http.StatusOK, strconv.FormatInt(result.Val(), 10))
}

func main(){
	client = redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	router := gin.Default()
	router.Use(gzip.Gzip(gzip.DefaultCompression))
	router.GET("/hello", hello)
	router.GET("/plus", plus)
	router.Run(":8080")
}

