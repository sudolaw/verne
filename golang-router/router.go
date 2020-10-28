package main

import (
	"golang-router/authmod"

	"github.com/gin-gonic/gin"
)

func setAuthRoute(r *gin.Engine) {
	authController := new(authmod.Controller)
	r.GET("/", authController.Auth)

}

func InitRoute() *gin.Engine {
	route := gin.New()
	route.Use(gin.Logger())
	route.Use(gin.Recovery())
	setAuthRoute(route)
	return route

}

func main() {

	router := InitRoute()
	router.Run("127.0.0.1:4000")
}
