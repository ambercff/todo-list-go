package main

import (
	"todo-go/configs"
	"todo-go/internal/controllers"
	"todo-go/internal/middlewares"
	"todo-go/internal/models"
	"github.com/gin-gonic/gin"
)

func main() {
	// Conexão com o banco de dados
	database.Connect()

	// Migrando automaticamente a estrutura do modelo
	database.DB.AutoMigrate(&models.User{})
	database.DB.AutoMigrate(&models.Task{})


    router := gin.Default()

	userController := controllers.NewUserController()
	taskController := controllers.NewTaskController()

	public := router.Group("/api")
   {
	public.POST("/register", userController.Register)
	public.POST("/login", userController.Login)
   }

   protected := router.Group("/api")
   protected.Use(middlewares.AuthMiddleware())
   {
	protected.GET("/tasks", taskController.GetAllTasks)
    protected.POST("/tasks", taskController.RegisterTask)
    protected.PATCH("/tasks", taskController.UpdateTask)
   }
   router.Run()
}
