package routers

import (
	"controllers/ReadExcel"
	"controllers/goEcharts"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Load(g *gin.Engine) *gin.Engine {
	g.HEAD("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{})
	})
	//Excel文件导入(批量)
	g.GET("/pf/practice", ReadExcel.ReadExcel)
	//通过数据生成可视化html页面
	g.GET("/pf/echarts", goEcharts.GoEcharts)

	g.NoRoute(func(c *gin.Context) {
		c.String(http.StatusNotFound, "The incorrect API route.")
	})
	return g
}
