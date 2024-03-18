package order

import "github.com/gin-gonic/gin"

func Reg(c *gin.Engine) {
	c.POST("/createOrder", Create)
}
