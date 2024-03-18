package user

import (
	"github.com/gin-gonic/gin"
)

func Reg(c *gin.Engine) {
	c.POST("/login", Login)
	c.POST("/register", Register)
	c.POST("/message", Message)
}
