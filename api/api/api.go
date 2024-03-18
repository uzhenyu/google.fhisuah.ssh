package api

import (
	"github.com/gin-gonic/gin"
	"zg5/lmonth/api/api/goods"
	"zg5/lmonth/api/api/order"
	"zg5/lmonth/api/api/user"
)

func Regs(c *gin.Engine) {
	user.Reg(c)
	goods.Reg(c)
	order.Reg(c)
}
