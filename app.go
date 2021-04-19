package main

import (
	"Taggy/internal/pkg/route"
	"github.com/gin-gonic/gin"
)

func main(){

	r := gin.Default()
	route.SetUp(r)
	r.Run(":8080")
}
