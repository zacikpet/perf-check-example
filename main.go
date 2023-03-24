package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"os"
	"time"

	docs "test-app/docs"

	"github.com/gin-gonic/gin"
	swaggerfiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"
)

// @Summary		Hello
// @Router			/ [get]
// @x-perfcheck	{ "latency": ["avg < 50", "min < 50", "avg_stat < 50"], "errorRate": ["avg_stat < 0.1"] }
func Helloworld(g *gin.Context) {

	sleep := rand.Intn(100)

	time.Sleep(time.Millisecond * time.Duration(sleep))
	g.JSON(http.StatusOK, "helloworld")
}

// @title			Example API
// @x-perfcheck	{ "stages": [{ "duration": "1s", "target": 5 }] }
func main() {
	r := gin.Default()

	port := os.Getenv("PORT")
	if port == "" {
		port = "8080"
	}

	host := os.Getenv("HOST")
	if host == "" {
		host = "localhost"
	}

	scheme := os.Getenv("SCHEME")
	if scheme == "" {
		scheme = "http"
	}

	docs.SwaggerInfo.Host = fmt.Sprintf("%s:%s", host, port)
	docs.SwaggerInfo.Schemes = []string{scheme}

	r.GET("/", Helloworld)

	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerfiles.Handler))

	r.Run(fmt.Sprintf(":%s", port))
}
