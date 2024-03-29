package main

import (
	"config"
	"github.com/gin-gonic/gin"
	"github.com/spf13/viper"
	"log"
	"net/http"
	"routers"
)

func main() {
	if err := config.Init(""); err != nil {
		panic(err)
	}
	//asda
	gin.SetMode(viper.GetString("runmode"))
	g := gin.Default()
	routers.Load(g)
	err := http.ListenAndServe(viper.GetString("port"), g)
	if err != nil {
		//logging.Error(err.Error())
		log.Println(err)
	}
}
