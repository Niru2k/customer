package handler

import (
	h "customer_module/helper"
	"customer_module/middleware"
	"customer_module/model"
	r "customer_module/repository"

	"github.com/gofiber/fiber/v2"
	"golang.org/x/crypto/bcrypt"

	"database/sql"

	_ "github.com/lib/pq"
)

type Database struct {
	Db *sql.DB
}

// To sign-up
func (db Database) Signup(c *fiber.Ctx) error {
	var UserData model.Customer
	if err := c.BodyParser(&UserData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			h.Status: fiber.StatusBadRequest,
			h.Err:    h.ReqBdyErr + err.Error(),
		})
	}
	if err := h.StructValidator(UserData, c); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			h.Status: fiber.StatusBadRequest,
			h.Err:    h.ValidatationErr + err.Error(),
		})
	}
	//Hashing the Password
	password, err := bcrypt.GenerateFromPassword([]byte(UserData.Password), bcrypt.DefaultCost)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			h.Status: fiber.StatusBadRequest,
			h.Err:    h.HashErr + err.Error(),
		})
	}
	UserData.Password = string(password)
	//Creating the new user
	if err = r.CreateUser(db.Db, UserData); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			h.Status: fiber.StatusBadRequest,
			h.Err:    err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		h.Status: fiber.StatusOK,
		h.Msg:    h.SignUpSuccess,
	})
}

// To Login
func (db Database) Login(c *fiber.Ctx) error {
	var data model.Customer
	var auth model.TokenAuthentication
	var err error
	if err = c.BodyParser(&data); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			h.Status: fiber.StatusBadRequest,
			h.Err:    h.ReqBdyErr + err.Error(),
		})
	}
	user, err := r.FindUser(db.Db, data.Email)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			h.Status: fiber.StatusBadRequest,
			h.Err:    h.UserExistErr,
		})
	}
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(data.Password)); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			h.Status: fiber.StatusBadRequest,
			h.Err:    h.PasswordErr,
		})
	}
	//creating the token
	token, err := middleware.CreateToken(user, c)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			h.Status: fiber.StatusBadRequest,
			h.Err:    err.Error(),
		})
	}
	auth.Email, auth.Token, auth.UserId = user.Email, token, user.CustomerId
	// Adding the token
	if err = r.AddToken(db.Db, auth); err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			h.Status: fiber.StatusBadRequest,
			h.Err:    err.Error(),
		})
	}
	return c.JSON(fiber.Map{
		h.Status: fiber.StatusOK,
		h.Token:  auth.Token,
	})
}

// To Get customer
func (db Database) GetCustomer(c *fiber.Ctx) error {
	tokenClaims := middleware.GetTokenClaims(c)
	customerData, err := r.FindUser(db.Db, tokenClaims.Subject)
	if err != nil {
		return c.Status(fiber.StatusBadRequest).JSON(fiber.Map{
			h.Status: fiber.StatusBadRequest,
			h.Err:    h.TokenErr,
		})
	}
	return c.JSON(fiber.Map{
		h.Res: map[string]interface{}{
			h.Fname:    customerData.FirstName,
			h.Lname:    customerData.LastName,
			h.MobileNo: customerData.MobileNo,
			h.Email:    customerData.Email,
		},
	})
}
