package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"go-jwt-auth/controllers"
)

func main() {

	r:= gin.Default()

	r.GET("/hello",func(ctx *gin.Context) {
		ctx.JSON(http.StatusOK,gin.H{ "message": "Hello World!"})
})
r.POST("/signup", controllers.SignUp)
r.POST("/login", controllers.Login)
r.Run(":8080")
}
