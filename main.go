package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID int
	Completed bool
	Body string
}

// CREANDO LA PRIMERA RUTA CON GO, REACT, TALWIND CSS Y TYPESCRIPT
func main(){
	fmt.Println("Hello BTS")
	app := fiber.New()

	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Hello Word"})
	})

    log.Fatal(app.Listen(":4000"))
}