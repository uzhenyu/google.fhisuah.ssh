package middleware

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"zg5/lmonth/api/conts"
	"zg5/lmonth/api/utils"
)

func Auth(c *gin.Context) {
	token := c.Request.Header.Get("token")
	err := utils.SetJwtToken(conts.Servername, token)
	if err != nil {
		c.JSON(http.StatusBadGateway, "token为空")
		c.Abort()
	}
	return
}
