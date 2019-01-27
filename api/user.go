package api

import (
	"github.com/google/uuid"
	"github.com/gin-gonic/gin"

	
)

func ListUsers(c *gin.Context) {
}

func CreateUUID(c *gin.Context) {
	id := uuid.New().String()
	RespSucess(c, id)
}