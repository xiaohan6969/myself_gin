package routers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"controllers/ReadExcel"
)

func Load(g *gin.Engine) *gin.Engine {
	g.HEAD("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
		})
	})
	//文件导入发货(批量)
	g.POST("/batch/delivery", ReadExcel.ReadExcel)





	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})
	return g
}
