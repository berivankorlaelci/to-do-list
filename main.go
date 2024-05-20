package main

import (
	"my-todo-app/handlers"
	"my-todo-app/middlewares"

	"github.com/gin-gonic/gin"
)

func main() {
	// Gin router oluştur
	r := gin.Default()

	// Kullanıcı kaydı ve giriş endpointleri
	r.POST("/register", handlers.RegisterUser)
	r.POST("/login", handlers.LoginUser)

	// Auth middleware'ini kullanarak korunan rotalar
	auth := r.Group("/")
	auth.Use(middlewares.AuthMiddleware())
	{
		auth.POST("/todo", handlers.CreateTodoList)
		// Diğer CRUD işlemleri için endpointler
	}

	r.GET("/users", handlers.GetAllUsers)

	// Sunucuyu başlat
	r.Run(":8080")
}
