package main

import (
	"fmt"
	"os"
	route "todobackend2/routes"
	errorHandler "todobackend2/utils"

	jwtware "github.com/gofiber/jwt/v3"
	// "github.com/golang-jwt/jwt/v4"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/cors"

	"github.com/joho/godotenv"
	// "github.com/gofiber/fiber/v2/middleware/basicauth"
)


var println = fmt.Println


func main(){

    prod := os.Getenv("PROD");
	if prod != "true" {
        
		err := godotenv.Load()
		errorHandler.HandleError(err)


	}
    
	
	app := fiber.New()
	app.Use(cors.New())
	
	

	app.Get("/", func(c *fiber.Ctx) error{


		return c.SendString("Hello, World!")
	})


	
	api := app.Group("/api")

    
	v1 := api.Group("/v1")

    v1.Route("user", route.UserRoute)
	

    
	v1.Use(jwtware.New(jwtware.Config{
		SigningKey: []byte("secret"),
	}))


	v1.Route("todo", route.TodoRoutes)
	
	
	// v1.Route("user", userRoute.UserRoutes)


	



    
	

     
	println("Now Listening")


	errorHandler.HandleError(app.Listen(":8080"))


}