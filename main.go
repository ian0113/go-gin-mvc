package main

import (
	"fmt"
	"log"

	"github.com/ian0113/go-gin-mvc/infra"
	"github.com/ian0113/go-gin-mvc/routes"

	"github.com/gin-gonic/gin"
)

func main() {
	infra.InitConfig(".", "config")
	cfg := infra.GetConfig()
	log.Printf("config --> %+v", cfg)

	infra.InitLogger(cfg)
	infra.InitDB(cfg)
	infra.InitRedis(cfg)

	r := gin.Default()
	routes.SetRouters(r)

	host := fmt.Sprintf("%s:%d", cfg.App.HostName, cfg.App.HostPort)
	r.Run(host)
}
