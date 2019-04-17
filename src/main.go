package main

import (
	"github.com/gin-gonic/gin"
	"routers"
)

func main() {
	g := gin.Default()
	routers.Load(g)
}
