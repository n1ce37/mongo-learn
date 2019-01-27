package api

import (
	"github.com/gin-gonic/gin"

	"github.com/n1ce37/aws/model"
)

func CreateHistory(c *gin.Context) {
	h := model.History{}
	err := c.BindJSON(&h)
	if err != nil {
		RespFail(c, err)
		return
	}

	id := c.Keys["uuid"]
	h.UUID = id.(string)
	data, err := h.Insert()
	if err != nil {
		RespFail(c, err)
		return
	}

	RespSucess(c, data)
}

func ListHistories(c *gin.Context) {
	id := c.Keys["uuid"]
	h := model.History{}
	h.UUID = id.(string)
	data, cnt, err := h.FindMany(0, 10)
	if err != nil {
		RespFail(c, err)
		return
	}

	RespSucess(c, data, cnt)
}

func ReadHistory(c *gin.Context) {}

func UpdateHistory(c *gin.Context) {}

func DeleteHistory(c *gin.Context) {}
