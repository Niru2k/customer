package router

import (
	h "customer_module/handler"
	"customer_module/middleware"
	"database/sql"

	"github.com/gofiber/fiber/v2"
)

func GetRoutes(db *sql.DB, app *fiber.App) {
	controller := h.Database{Db: db}
	app.Post("/v1/signUp", controller.Signup)
	app.Post("/v1/login", controller.Login)
	AuthRoutes := app.Use(middleware.IsAuthenticate(db))
	AuthRoutes.Get("/v1/getUser", controller.GetCustomer)
}
