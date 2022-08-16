package main

import (
	"backend.api/config"
	"backend.api/routers"
	"github.com/gin-gonic/gin"
)

func init() {
	config.ConnectDatabase()
}

func main() {
	r := gin.Default()

	routers.SetupRouter(r)

	r.Run(":8080")
}
