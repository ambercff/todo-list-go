package controllers

import (
	"net/http"
	"todo-go/internal/models"
	"todo-go/internal/repository"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"todo-go/internal/utils"
)

type UserController struct {
	UserRepository * repository.UserRepository
}

type UserDTO struct {
Username string `json: "username" binding:"required"`
Password string `json: "password" binding:"required"`
}

func NewUserController() *UserController {
	userRepository := repository.NewUserRepository()
	return &UserController{UserRepository: userRepository}
}

func (u * UserController) Register (c * gin.Context){

	var userDTO UserDTO

	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "verifique novamente os campos!"})
		return
	}

	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(userDTO.Password), bcrypt.DefaultCost)

	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "erro ao criptografar a senha"})
		return
	}

	user := models.User{
		Username: userDTO.Username,
		Password: string(hashedPassword),
	}

	if err := u.UserRepository.CreateUser(&user); err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "falha ao registrar usuário"})
		return
	}

	c.IndentedJSON(http.StatusOK, gin.H{"message": "usuário registrado com sucesso!"})
}

func (u * UserController) Login (c * gin.Context) {
	var userDTO UserDTO

	if err := c.ShouldBindJSON(&userDTO); err != nil {
		c.IndentedJSON(http.StatusBadRequest, gin.H{"message": "verifique novamente os campos!"})
		return
	}

	user, err := u.UserRepository.GetUserByUsername(userDTO.Username)

	if err != nil {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "usuário e/ou senha inválidos"})
		return
	}

	if !u.UserRepository.CheckPassword(user.Password, userDTO.Password) {
		c.IndentedJSON(http.StatusUnauthorized, gin.H{"message": "usuário e/ou senha inválidos"})
	}

	token, err := utils.GenerateToken(user.ID)
	
	if err != nil {
		c.IndentedJSON(http.StatusInternalServerError, gin.H{"message": "erro ao gerar token"})
		return
	}
	c.IndentedJSON(http.StatusOK, gin.H{"token": token})
}