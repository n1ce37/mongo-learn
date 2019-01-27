package model

func getOffset(page, size int) int {
	return (page - 1) * size
}
