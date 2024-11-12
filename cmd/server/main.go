package main

import (
	"todo-go/configs"
	"todo-go/internal/models"
	"todo-go/internal/controllers"
	"github.com/gin-gonic/gin"
	"todo-go/internal/database"
	// "todo-go/internal/middlewares"
)

func main() {
	// Conexão com o banco de dados
	database.Connect()

	// Migrando automaticamente a estrutura do modelo
	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.Task{})


    router := gin.Default()

	userController := controllers.NewUserController()

	public := router.Group("/api")
   {
	public.POST("/register", userController.Register)
	public.POST("/login", userController.Login)
   }

//    protected := router.Group("/api")
//    protected.Use(middlewares.AuthMiddleware()){
// 	protected.GET("/tasks", taskController.GetTasks)
//     protected.POST("/tasks", taskController.CreateTask)
//    }
   router.Run()
}
