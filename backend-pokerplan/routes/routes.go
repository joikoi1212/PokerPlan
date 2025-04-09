package routes

import (
	"backend/handlers"
	"fmt"

	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()
	r.GET("/login", handlers.LoginHandler)
	r.GET("/callback", handlers.CallbackHandler)
	//r.GET("/dashboard", handlers.AuthMiddleware)
	return r
}

func Print() {
	fmt.Println("works")
}
