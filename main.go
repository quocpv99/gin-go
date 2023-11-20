package main

import (
	"./handlers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	handlers.InitRouter(r)
	r.Run()
}
