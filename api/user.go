package api

import (
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

func ListUsers(c *gin.Context) {
}

func CreateUUID(c *gin.Context) {
	id := uuid.New().String()
	RespSucess(c, id)
}
