package main

import (
	"fmt"
	"log"

	"github.com/gofiber/fiber/v2"
)

type Todo struct {
	ID        int     `json: "id"`
	Completed bool    `json: "completed"`
	Body      string  `json: "body"`
}

func main(){
	fmt.Println("Hello BTS")
	app := fiber.New()

	todos := []Todo{}


	// To get a TODO
	app.Get("/", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(fiber.Map{"msg": "Hello Word"})
	})

	// To post a TODO
	app.Post("/api/todos",  func(c *fiber.Ctx) error {
		todo := &Todo{}  

	   if err := c.BodyParser(todo); err != nil {
		return err
	   }

	   if todo.Body == "" {
		return c.Status(400).JSON(fiber.Map{"error" : "Todo body is required"})
	   }

	   todo.ID = len(todos) + 1
	   todos = append(todos, *todo)

	   return c.Status(201).JSON(todo)
	})

    log.Fatal(app.Listen(":4000"))
}