package routers

import (
	"controllers/ReadExcel"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(g *gin.Engine) *gin.Engine {
	g.HEAD("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})
	//文件导入(批量)
	g.POST("/pf/operation/excel", ReadExcel.ReadExcel)
	g.GET("/pf/practice", ReadExcel.ReadExcel)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})
	return g
}
