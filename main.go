package main

import (
	"github.com/Bhupesh-V/roman-to-no-api/controllers"
	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()
	r.POST("/roman", controllers.ConvertRoman)
	r.Run()
}
