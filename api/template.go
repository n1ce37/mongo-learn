package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mongodb/mongo-go-driver/bson/primitive"

	"github.com/n1ce37/mongo-learn/model"
)

func CreateTemplate(c *gin.Context) {
	t := model.Template{}
	err := c.BindJSON(&t)
	if err != nil {
		RespFail(c, err)
		return
	}
	data, err := t.Insert()
	if err != nil {
		RespFail(c, err)
		return
	}

	RespSucess(c, data)
}

func ListTemplates(c *gin.Context) {
	t := model.Template{}
	page, size := getPageAndSize(c)
	data, cnt, err := t.FindMany(page, size)
	if err != nil {
		RespFail(c, err)
		return
	}

	RespSucess(c, data, cnt)
}

func ReadTemplate(c *gin.Context) {
	id := c.Param("id")
	t := model.Template{}
	objectID, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		RespFail(c, err)
		return
	}

	t.ID = objectID
	err = t.FindOne()
	if err != nil {
		RespFail(c, err)
		return
	}

	RespSucess(c, t)
}
