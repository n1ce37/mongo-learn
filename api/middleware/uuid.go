package middleware

import (
	"errors"

	"github.com/gin-gonic/gin"
	"github.com/n1ce37/mongo-learn/api"
)

func GetUUID(c *gin.Context) {
	id := c.Request.Header.Get("uuid")
	if len(id) == 0 {
		api.RespFail(c, errors.New("header must have uuid"))
		c.Abort()
		return
	}
	c.Keys = make(map[string]interface{})
	c.Keys["uuid"] = id
}
