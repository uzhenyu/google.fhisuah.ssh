package api

import (
	"google.golang.org/grpc"
	"zg5/lmonth/message/goods"
	"zg5/lmonth/message/order"
	"zg5/lmonth/message/user"
)

func Reg(c grpc.ServiceRegistrar) {
	user.RegisterUserServer(c, ServiceUser{})
	goods.RegisterGoodsServer(c, ServiceGood{})
	order.RegisterOrderServer(c, ServerOrder{})
}
