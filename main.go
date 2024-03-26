package main

import (
	"fmt"

	"github.com/gin-gonic/gin"
	"prolixless/api/handlers"
)

func main() {
	fmt.Println("Hello world!")

	gin.SetMode(gin.DebugMode)

	r := gin.New()
	handler := handlers.GinHandler{}
	handler.SetupEndpoints(r)

	if err := r.Run(":3000"); err != nil {
		panic("Failed to run gin server")
	}
}
