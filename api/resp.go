package api

import (
	"github.com/gin-gonic/gin"
)

type respMeta struct {
	Status bool        `json:"status"`
	Data   interface{} `json:"data"`
	Total  int         `json:"total,omitempty"`
}

func RespSucess(c *gin.Context, data interface{}, total ...int) {
	respData := respMeta{
		Status: true,
		Data:   data,
	}
	if len(total) != 0 {
		respData.Total = total[0]
	}
	c.JSON(200, &respData)
}

func RespFail(c *gin.Context, err error) {
	respData := respMeta{
		Status: false,
		Data:   err.Error(),
	}
	c.JSON(200, &respData)
}
