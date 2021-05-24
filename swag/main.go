package main

import (
	"swag/controller"

	"github.com/gin-gonic/gin"
)

func main() {
	r := gin.Default()

	c := controller.NewController()

	v1 := r.Group("/api/v1")
	{
		accounts := v1.Group("/accounts")
		{
			accounts.GET(":id", c.ShowAccount)
			accounts.GET("", c.ShowAccount)
			accounts.POST("", c.ShowAccount)
			accounts.DELETE(":id", c.ShowAccount)
			accounts.PATCH(":id", c.ShowAccount)
			accounts.POST(":id/images", c.ShowAccount)
		}
	}
	r.Run(":10808")
}
