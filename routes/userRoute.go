
package route


import (
	"github.com/gofiber/fiber/v2"
	controller "todobackend2/controllers"

)


func UserRoute (user fiber.Router){


    user.Post("/signup", controller.Signup)
	user.Post("/login", controller.Login)

}



