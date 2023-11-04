package router

import (
	h "customer_module/handler"

	"github.com/gofiber/fiber/v2"
	"github.com/jinzhu/gorm"
)

func GetRoutes(db *gorm.DB, app *fiber.App) {
	controller := h.Database{Db: db}
	app.Post("/v1/signUp", controller.Signup)
	app.Post("/v1/login", controller.Login)
	app.Get("/v1/getUser", controller.GetCustomer)
}
