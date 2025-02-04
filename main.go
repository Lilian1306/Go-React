package main

import (    
	"fmt"
	"log"
	"os"
	"context"
	
	"github.com/joho/godotenv"
    "github.com/gofiber/fiber/v2"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
    "go.mongodb.org/mongo-driver/mongo/options"

)

type Todo struct {
	ID         primitive.ObjectID    `json:"_id,omitempty" bson:"_id,omitempty"`
	Completed  bool                  `json:"completed"`
	Body	   string                `json:"body"`
}

var collection  *mongo.Collection

func main(){
	fmt.Println("Hello World")

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file:", err)
	}

	MONGODB_URI := os.Getenv("MONGODB_URI")
	clientOptions := options.Client().ApplyURI(MONGODB_URI)
	client, err := mongo.Connect(context.Background(), clientOptions)

	if err != nil {
		log.Fatal(err)
	}

	defer client.Disconnect(context.Background())

	err = client.Ping(context.Background(), nil)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Println("Connected to MONGODB ATLAS")

	collection = client.Database("golang_db").Collection("todos")

	app := fiber.New()

	app.Get("/api/todos", getTodos)
	app.Post("/api/todos", createTodo)
	app.Patch("/api/todos/:id", updateTodo)
	app.Delete("/api/todos/:id", deleteTodo)

	port := os.Getenv("PORT")
	 if port == ""{
		port = "5000"
	 }

	 log.Fatal(app.Listen("0.0.0.0:" + port))
}

// Codigo para obtener los TODO del proyecto. 
func getTodos(c *fiber.Ctx) error {
	var todos []Todo

	cursor, err := collection.Find(context.Background(), bson.M{})

	if err != nil {
		return err
	}

	defer cursor.Close(context.Background())

	for cursor.Next(context.Background()){
		var todo Todo 
		if err := cursor.Decode(&todo); err != nil {
			return err
		}
		todos = append(todos, todo)
	}

	return c.JSON(todos)
}

// Codigo para crear un nuevo TODO en el proyecto
func createTodo(c *fiber.Ctx) error {
	todo := new(Todo)

	if err := c.BodyParser(todo); err != nil {
		return err
	}

	if todo.Body == ""{
		return c.Status(400).JSON(fiber.Map{"error": "Todo body cannot be empty"})
	}
	insertResult, err := collection.InsertOne(context.Background(),todo)
	if err != nil {
		return err
	}
	todo.ID = insertResult.InsertedID.(primitive.ObjectID)

	return c.Status(201).JSON(todo)
}

// Codigo para actualizar el TODO en el proyecto
func updateTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID, err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid todo ID"})
	}

	filter := bson.M{"_id": objectID}
	update := bson.M{"$set":bson.M{"completed": true}}

	_, err = collection.UpdateOne(context.Background(), filter, update)
	if err != nil {
		return err
	}

	return c.Status(200).JSON(fiber.Map{"success": true})
}

// Codigo para eliminar el TODO en el proyecto
func deleteTodo(c *fiber.Ctx) error {
	id := c.Params("id")
	objectID,err := primitive.ObjectIDFromHex(id)

	if err != nil {
		return c.Status(400).JSON(fiber.Map{"error": "Invalid todo ID"})
	}
    
	filter := bson.M {"_id":objectID}
	_,err = collection.DeleteOne(context.Background(), filter)

	if err != nil {
		return err
	}
	return c.Status(200).JSON(fiber.Map{"success": true})
}









/*package main

import (
	"fmt"
	"log"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/joho/godotenv"
)

type Todo struct {
	ID        int     `json: "id"`
	Completed bool    `json: "completed"`
	Body      string  `json: "body"`
}

func main(){
	fmt.Println("Hello BTS")
	app := fiber.New()

	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	PORT := os.Getenv("PORT")
	

	todos := []Todo{}


	// To get a TODO
	app.Get("/api/todos", func(c *fiber.Ctx) error {
		return c.Status(200).JSON(todos)
	})

	// Create a TODO
	app.Post("/api/todos",  func(c *fiber.Ctx) error {                            // Creamos un nuevo endpoint en la aplicacion. 
		todo := &Todo{}                                                           // Se crea una nueva instancia de la estructura "todo"

	   if err := c.BodyParser(todo); err != nil {                                 // Trata de leer el cuerpo de la solicitud, generalmente en formato JSON y lo convierte en un objeto de tipo "todo". Si occure un error al analizar el cuerpo (si no es un JSON valido), el error se devuelve y se detiene le proceso. 
		return err
	   }

	   if todo.Body == "" {                                                       // Se veerifica si el campo body de la tarea esta vacio
		return c.Status(400).JSON(fiber.Map{"error" : "Todo body is required"})   // Si esta vacio, ser responde con un codigo de estado 400 y un mensaje de error.
	   }

	   todo.ID = len(todos) + 1                                                  // se asigna un ID unico a la tarea, este ID es calculado como el tamaño actual de la lista "todos" + 1. La lista todos contiene todas las tareas existentes. Al calcular len(todos), obtenemos el número de tareas actuales, y al sumar 1, obtenemos el siguiente ID disponible.
	   todos = append(todos, *todo)                                              // Se agrega la tarea recien creada a la lista "todos" usando la funcion "append"

	   return c.Status(201).JSON(todo)                                           // Se responde con un codigo 201 (created), indica que se ha creado la tarea correctamente. La respuesta incluye el objeto "todo" reacien creado en formato JSON.
	})

	// update a TODO
	app.Patch("/api/todos/:id", func(c *fiber.Ctx) error {  // Crear una ruta patch en la aplicacion app
		id := c.Params("id")                                // Extraemos el parametro id de url usando este codigo, convirtiendo el id en un string

		for i, todo := range todos {                        // Se recorre la list "todos" con un for
			if fmt.Sprint(todo.ID) == id {                  // En este codigo comparamos el id ingresado en la lista de "todos" y fmt.Sprint(todo.ID) convierte todo.ID a string para compararlo con id, ya que id es un string.
              todos[i].Completed = true                     // Se accede al indice "i" en "todos" y se actualiza el campo
			  return c.Status(200).JSON(todos[i])           // Se responde con un estado 200 OK y luebgo se devuelve la tarea actualizada en formato JSON.
			}
		}

		return c.Status(400).JSON(fiber.Map{"error": "Todo not found"})   // si el for no encuentra la tarea o el ID ingresado, significa que el ID no existe y se devuelve un estado 400 Bad Request con un mensaje de error. 
	})

	// delete a TODO
    app.Delete("/api/todos/:id", func(c *fiber.Ctx) error {                // Se crea un nuevo endpoint con la ruta DELETE en la aplicacion, esta ruta espera un parametro de URL (id) para indentificar cual tarea eliminar. 
		id := c.Params("id")                                               // obtiene el parametro (id) de la URL, los parametrod de URL siempre son cadenas, "id" sera un string
 
		for i, todo := range todos {                                       // Se recorre la lista "todos" usando un bucle for con range, "i" es el indice del elemento actual, "todo" es el objeto "todo" en la posicion "i" dentro de array "todos"
			if fmt.Sprint(todo.ID) == id {                                 // "todo.ID" es un numero entero (int), pero "id" es un string, if fmt.Sprint(todo.ID) convierte el "ID" a un string para compararlo con "ID", si hay concidencia, significa que hemos encontrado la tarea que queremos eliminar. 
				todos = append(todos[:i], todos[i+1:]...)                  // Eliminar la tarea encontrada, todos[:i] toma los elementos antes del indice, todos[i+1:] tomo los elementos despues del indice, append(...) Une ambas partes, excluyendo el elemento en "i" 
				return c.Status(200).JSON(fiber.Map{"success": true})      // Despues de eliminar el "todo", enviamos una respuesta 200 OK con un JSON confimando que la operacion fue exitosa. Como usamos return, la función termina aquí y no sigue ejecutando el bucle.
			}
		}

		return c.Status(400).JSON(fiber.Map{"error": "Todos not found"})   // Si el bucle termina sin encontrar el "ID", significa que la tarea no existe y se responde con un codigo 400 y un mesaje de error. 
	})

    log.Fatal(app.Listen(":"+PORT))
}*/