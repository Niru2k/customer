package middleware

import (
	h "customer_module/helper"
	"customer_module/model"
	r "customer_module/repository"
	"database/sql"
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/dgrijalva/jwt-go"
	"github.com/gofiber/fiber/v2"
	_ "github.com/lib/pq"
)

// Creating the JWT token
func CreateToken(user model.Customer, c *fiber.Ctx) (string, error) {
	//Loading the .env file
	if err := h.ConfigEnv(".env"); err != nil {
		fmt.Println(h.Err, err)
	}
	exp := time.Now().Add(time.Hour * h.ExpTime).Unix()
	userStr := strconv.Itoa(user.CustomerId)
	claims := jwt.StandardClaims{
		ExpiresAt: exp,
		Id:        userStr,
		IssuedAt:  time.Now().Unix(),
		Subject:   user.Email,
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

func IsAuthenticate(db *sql.DB) fiber.Handler {
	//Loading the .env file
	if err := h.ConfigEnv(".env"); err != nil {
		fmt.Println(h.Err, err)
	}
	return func(c *fiber.Ctx) error {
		tokenString := c.Get("Authorization")
		if tokenString == "" {
			return c.Status(fiber.StatusUnauthorized).JSON(fiber.Map{
				"message": "token is empty",
			})
		}
		for index, char := range tokenString {
			if char == ' ' {
				tokenString = tokenString[index+1:]
			}
		}
		claims := jwt.StandardClaims{}
		jwt.ParseWithClaims(tokenString, &claims, func(token *jwt.Token) (interface{}, error) {
			return []byte(os.Getenv("SECRET_KEY")), nil
		})
		_, err := r.FindUsersByToken(db, tokenString)
		if err != nil {
			return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
				h.Status: fiber.StatusBadRequest,
				h.Err:    h.UserExistErr,
			})
		}
		return c.Next()
	}
}
