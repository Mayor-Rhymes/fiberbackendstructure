package controller

import (
	"context"
	"time"

	"os"
	db "todobackend2/config"
	errorHandler "todobackend2/utils"

	model "todobackend2/models"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)






func Signup(c *fiber.Ctx) error {

    collection := db.HandleUserDB(os.Getenv("DATABASE_URL"))
	data := new(model.User)
	err := c.BodyParser(data)

	errorHandler.HandleError(err)
    hashedPassword, err := bcrypt.GenerateFromPassword([]byte(data.Password), 10)

    var result model.User
	errorHandler.HandleError(err)
	_ = collection.FindOne(context.Background(), bson.D{primitive.E{Key:"email", Value:data.Email}}).Decode(&result)

	if result.Email != ""{

         return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{"message": "Email already in use"});
	}
    _, err = collection.InsertOne(context.Background(), bson.M{"email": data.Email, "username": data.Username, "password": hashedPassword})

	
	if err != nil{

		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{"message": "User wasn't created"})
	}




	claims := jwt.MapClaims{
		
		"email": data.Email,
		"username":  data.Username,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "User created successfully", "user": data, "token": t})


	

	
}



func Login(c *fiber.Ctx) error {

    collection := db.HandleUserDB(os.Getenv("DATABASE_URL"))
	data := new(model.User)
	err := c.BodyParser(data)

	errorHandler.HandleError(err)
	var result model.User
	err = collection.FindOne(context.Background(), bson.D{primitive.E{ Key:"email", Value:data.Email}}).Decode(&result)
	errorHandler.HandleError(err)
	err = bcrypt.CompareHashAndPassword([]byte(result.Password), []byte(data.Password))

	

	


	
	if err != nil{

		return c.Status(fiber.StatusNotAcceptable).JSON(fiber.Map{"message": "Password is incorrect"})
	}


	claims := jwt.MapClaims{
		
		"email": data.Email,
		"username":  data.Username,
		"exp":   time.Now().Add(time.Hour * 72).Unix(),
	}

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	// Generate encoded token and send it as response.
	t, err := token.SignedString([]byte("secret"))

	if err != nil {
		return c.SendStatus(fiber.StatusInternalServerError)
	}
	
	return c.Status(fiber.StatusOK).JSON(fiber.Map{"message": "Welcome back, " + data.Username, "user": data, "token": t})


	

	
}



