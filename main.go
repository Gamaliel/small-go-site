package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main() {
	app := fiber.New()

	app.Static("/", "./client/dist")

	app.Listen(":3000")
	fmt.Println("Serve at port: 3000")
}
