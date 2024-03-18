package goods

import (
	"github.com/beego/beego/v2/core/logs"
	"github.com/gin-gonic/gin"
	"zg5/lmonth/api/conts"
	"zg5/lmonth/api/service"
	"zg5/lmonth/framework/http"
	"zg5/lmonth/message/goods"
)

func Create(c *gin.Context) {
	var Good *goods.GoodsInfo
	if err := c.ShouldBind(&Good); err != nil {
		http.Re(c, conts.PRM_ERROR, err.Error(), nil)
	}
	good, err := service.CreateGood(c, Good)
	if err != nil {
		http.Re(c, conts.PRM_ERROR, "添加失败", nil)
		return
	}
	http.Re(c, conts.PRM_ERROR, "添加成功", good)
	return
}

func Delete(c *gin.Context) {
	var Good struct {
		Id int64 `json:"Id"`
	}
	if err := c.ShouldBind(&Good); err != nil {
		http.Re(c, conts.PRM_ERROR, err.Error(), nil)
	}
	if Good.Id == 0 {
		http.Re(c, conts.PRM_ERROR, "id不能为0", nil)
	}
	err := service.DeleteGood(c, Good.Id)
	if err != nil {
		http.Re(c, conts.PRM_ERROR, "删除失败", nil)
		return
	}
	http.Re(c, conts.SUCCESS, "删除成功", nil)
	return
}

func Update(c *gin.Context) {
	var Good *goods.GoodsInfo
	if err := c.ShouldBind(&Good); err != nil {
		http.Re(c, conts.PRM_ERROR, err.Error(), nil)
	}
	good, err := service.UpdateGood(c, Good)
	if err != nil {
		http.Re(c, conts.PRM_ERROR, "修改失败", nil)
		return
	}
	http.Re(c, conts.SUCCESS, "修改成功", good)
	return
}

func Select(c *gin.Context) {
	var Good struct {
		Id int64 `json:"Id"`
	}
	if err := c.ShouldBind(&Good); err != nil {
		http.Re(c, conts.PRM_ERROR, err.Error(), nil)
		return
	}
	if Good.Id == 0 {
		http.Re(c, conts.PRM_ERROR, "id不能为0", nil)
		return
	}
	info, err := service.SelectGoods(c, Good.Id)
	if err != nil {
		http.Re(c, conts.PRM_ERROR, "查询失败", nil)
		return
	}
	http.Re(c, conts.SUCCESS, "查询成功", info)
	return
}

func GetGoods(c *gin.Context) {
	var Good struct {
		Offset    int64 `json:"Offset"`
		Limit     int64 `json:"Limit"`
		GoodsType int64 `json:"GoodsType"`
	}
	if err := c.ShouldBind(&Good); err != nil {
		http.Re(c, conts.PRM_ERROR, err.Error(), nil)
		return
	}
	getGoods, err := service.GetGoods(c, Good.Offset, Good.Limit, Good.GoodsType)
	if err != nil {
		return
	}
	http.Re(c, conts.SUCCESS, "查询成功", getGoods)
	logs.Info("11")
	return
}
