package api

import (
	"strconv"

	"github.com/gin-gonic/gin"
)

func getPageAndSize(c *gin.Context) (int, int) {
	pageStr, _ := c.GetQuery("p")
	sizeStr, _ := c.GetQuery("s")

	page, _ := strconv.Atoi(pageStr)
	size, _ := strconv.Atoi(sizeStr)

	return page, size
}
