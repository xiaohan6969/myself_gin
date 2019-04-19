package goEcharts

import (
	"fmt"
	"github.com/chenjiandongx/go-echarts/charts"
	"github.com/gin-gonic/gin"
	"gopkg.in/mgo.v2/bson"
	"log"
	"os"
)

func GoEcharts(c *gin.Context) {
	items := []string{"111", "232", "是的撒"}
	a := charts.NewBar()
	a.SetGlobalOptions(charts.TitleOpts{Title: "示例图"})
	a.AddXAxis(items).
		AddYAxis("A", []int{12, 231, 5, 45, 6546, 8, 5}).
		AddYAxis("B", []int{123, 43543, 546, 5688, 678})
	f, err := os.Create("a.html")
	if err != nil {
		log.Println(err)
	}
	a.Render(f)
}

func AAAAA(c *gin.Context)  {
	type A struct {
		AA bson.ObjectId `json:"ObjectId"`
 	}
	a := A{}
	c.Bind(&a)
	fmt.Println("===",a.AA)
}
