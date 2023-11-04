package middleware

import (
	h "customer_module/helper"
	"customer_module/model"
	"fmt"
	"os"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
)

// Creating the JWT token
func CreateToken(user model.Customer, c *fiber.Ctx) (string, error) {
	//Loading the .env file
	if err := h.ConfigEnv(".env"); err != nil {
		fmt.Println(h.Err, err)
	}
	exp := time.Now().Add(time.Hour * h.ExpTime).Unix()
	claims := jwt.StandardClaims{
		ExpiresAt: exp,
		Id:        user.Email,
		IssuedAt:  time.Now().Unix(),
	}
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET_KEY")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

// Get a claims from the token
func GetTokenClaims(c *fiber.Ctx) jwt.StandardClaims {
	//Loading the .env file
	if err := h.ConfigEnv(".env"); err != nil {
		fmt.Println(h.Err, err)
	}
	tokenString := c.Get(h.Authorization)
	for index, char := range tokenString {
		if char == ' ' {
			tokenString = tokenString[index+1:]
		}
	}
	claims := jwt.StandardClaims{}
	jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
		return []byte(os.Getenv("SECRET_KEY")), nil
	})
	return claims
}
