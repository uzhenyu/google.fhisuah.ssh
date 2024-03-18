package goods

import (
	"github.com/gin-gonic/gin"
)

func Reg(c *gin.Engine) {
	//c.Use(middleware.Auth)
	c.POST("/create", Create)
	c.POST("/delete", Delete)
	c.POST("/update", Update)
	c.POST("/select", Select)
	c.POST("/getGoods", GetGoods)
}
