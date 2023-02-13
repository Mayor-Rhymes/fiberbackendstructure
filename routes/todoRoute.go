package route

import (
	"github.com/gofiber/fiber/v2"
	controller "todobackend2/controllers"
	
    
)





func TodoRoutes (todo fiber.Router){
   
    todo.Get("/", controller.GetAll)
	todo.Get("/:id", controller.GetOne)
	todo.Post("/", controller.PostTodo)
	todo.Delete("/:id", controller.DeleteTodo)
	todo.Patch("/:id", controller.UpdateTodo)

}