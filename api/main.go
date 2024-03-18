package main

import (
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
	"zg5/lmonth/api/api"
)

func main() {
	r := gin.Default()
	config := cors.DefaultConfig()
	config.AllowAllOrigins = true
	r.Use(cors.New(config))
	api.Regs(r)
	r.Run(":8080")
}
