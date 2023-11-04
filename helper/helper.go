package helper

import (
	"customer_module/model"

	"github.com/go-playground/validator"
	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

func ConfigEnv(file string) error {
	if err := godotenv.Load(file); err != nil {
		return err
	}
	return nil
}

func StructValidator(data model.Customer, c *fiber.Ctx) error {
	validate := validator.New()
	if err := validate.Struct(data); err != nil {
		return err
	}
	return nil
}
