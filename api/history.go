package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mongodb/mongo-go-driver/bson/primitive"

	"github.com/n1ce37/mongo-learn/model"
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
	page, size := getPageAndSize(c)
	data, cnt, err := h.FindMany(page, size)
	if err != nil {
		RespFail(c, err)
		return
	}

	RespSucess(c, data, cnt)
}

func ReadHistory(c *gin.Context) {
	id := c.Param("id")
	h := model.History{}
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		RespFail(c, err)
		return
	}

	h.ID = objectID
	h.UUID = c.Keys["uuid"].(string)
	err = h.FindOne()
	if err != nil {
		RespFail(c, err)
		return
	}

	RespSucess(c, h)
}

func UpdateHistory(c *gin.Context) {
	id := c.Param("id")
	h := model.History{}
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		RespFail(c, err)
		return
	}
	err = c.BindJSON(&h)
	if err != nil {
		RespFail(c, err)
		return
	}

	h.ID = objectID
	h.UUID = c.Keys["uuid"].(string)

	err = h.UpdateOne()
	if err != nil {
		RespFail(c, err)
		return
	}

	RespSucess(c, h)
}