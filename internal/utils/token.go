package utils

import (
	"os"
	"time"
	"github.com/dgrijalva/jwt-go"
)

type Claims struct {
	UserID int `json: "user_id"`
	jwt.StandardClaims
}

var secretKey = []byte(os.Getenv("JWT_SECRET_KEY"))

func GenerateToken(userID int) (string, error){
	claims := Claims{
		UserID: userID,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(time.Hour * 24).Unix(), // Configurando a expiração do token em 24 horas"
			Issuer: "todo-go",
		},
	}

	// Criando o token
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenStr, err := token.SignedString(secretKey)

	if err != nil {
		return "", err
	}

	return tokenStr, nil
}

func ValidateToken(tokenString string) (*Claims, error) {
	token, err := jwt.ParseWithClaims(tokenString, &Claims{}, func(token * jwt.Token) (interface{}, error){
	
	return secretKey, nil
})

	if err != nil || !token.Valid {
		return nil, err
	}

	if claims, ok := token.Claims.(*Claims); ok {
		return claims, nil
	}

	return nil, err
}