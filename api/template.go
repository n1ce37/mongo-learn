package api

import (
	"github.com/gin-gonic/gin"

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

func ListTemplates(c *gin.Context) {}

func ReadTemplate(c *gin.Context) {}

func UpdateTemplate(c *gin.Context) {}

func DeleteTemplate(c *gin.Context) {}
