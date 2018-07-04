package main

import (
	"fmt"
	"net/http"
	"runtime"

	"github.com/gin-gonic/gin"
)

func main() {
	initializeConfiguration()

	runtime.GOMAXPROCS(runtime.NumCPU())
	gin.SetMode(config.MODE)
	router := gin.Default()

	router.GET("/", Index)

	server := &http.Server{
		Addr:    "127.0.0.1:8080",
		Handler: router,
	}
	if err := server.ListenAndServe(); err != nil {
		fmt.Println(err)
	}
}

// Index _
func Index(c *gin.Context) {
	c.String(http.StatusOK, "Hello world")
}
