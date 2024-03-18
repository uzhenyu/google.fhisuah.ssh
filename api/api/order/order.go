package order

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/shopspring/decimal"
	"zg5/lmonth/api/conts"
	"zg5/lmonth/api/service"
	"zg5/lmonth/api/utils"
	"zg5/lmonth/framework/http"
	"zg5/lmonth/message/order"
)

func Create(c *gin.Context) {
	var orders *order.OrderInfo
	if err := c.ShouldBind(&orders); err != nil {
		http.Re(c, conts.PRM_ERROR, err.Error(), nil)
	}
	goods, err := service.SelectGoods(c, orders.GoodsId)
	if err != nil {
		http.Re(c, conts.PRM_ERROR, err.Error(), nil)
		return
	}
	p, _ := decimal.NewFromString(goods.Price)
	orders.Total = fmt.Sprintf("%v", p.Mul(decimal.NewFromInt(orders.Num)))
	orders.OrderIdCard = utils.Order()
	createOrder, err := service.CreateOrder(c, orders)
	if err != nil {
		http.Re(c, conts.PRM_ERROR, err.Error(), nil)
		return
	}
	http.Re(c, conts.SUCCESS, "成功", createOrder)
	return
}
