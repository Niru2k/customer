package main

import (
	d "customer_module/driver"
	r "customer_module/router"
	"log"

	"github.com/gofiber/fiber/v2"
)

func main() {
	db := d.DbConnection()
	app := fiber.New()
	r.GetRoutes(db, app)
	log.Println("Server starts in port 8080...")
	//starting the server
	app.Listen(":8080")
}
