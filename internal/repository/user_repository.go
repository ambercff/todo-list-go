package repository

import (
	"todo-go/configs"
	"todo-go/internal/models"
	"golang.org/x/crypto/bcrypt"
)

type UserRepository struct {
}

// Nova instância de UserRepository
func NewUserRepository() *UserRepository {
	return &UserRepository{}
}

// Cria um novo usuário no banco de dados
func (r * UserRepository) CreateUser(user *models.User) error {
	return database.DB.Create(user).Error
}

func (r * UserRepository) GetUserByUsername(username string) (*models.User, error) {
	var user models.User
	result := database.DB.Where("username = ?", username).First(&user)
	return &user, result.Error
}

func (r * UserRepository) DeleteUser(username string) (string, error) {
	var user models.User

	result := database.DB.Where("username = ?", username).Delete(&user)
	if result.Error != nil {
		return "", result.Error
	}

	if result.RowsAffected == 0 {
		return "No user found with that username", nil
	}

	return "User deleted successfully", nil
}

func (r * UserRepository) CheckPassword(storedPassword, providedPassword string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(storedPassword), []byte(providedPassword))
	return err == nil
}