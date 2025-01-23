package main

import (
	"fmt"

	"github.com/gofiber/fiber/v2"
)

func main(){
	fmt.Println("Hello Word")
	app := fiber.New()

    log.Fatal(app.Listen(":4000"))
}