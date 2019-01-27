package main

import "github.com/n1ce37/aws/router"

func main() {
	s := router.SetupRouter()
	s.Run(":8080")
}
