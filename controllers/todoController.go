package controller

import (
	"context"
	"os"

	model "todobackend2/models"
	errorHandler "todobackend2/utils"

	db "todobackend2/config"

	"github.com/gofiber/fiber/v2"
	"github.com/golang-jwt/jwt/v4"
	// "github.com/golang-jwt/jwt/v4"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo/options"
)



func GetAll(c *fiber.Ctx) error{


	user := c.Locals("user").(*jwt.Token)
	claims := user.Claims.(jwt.MapClaims)
	username := claims["username"].(string)
	

    collection := db.HandleTodoDB(os.Getenv("DATABASE_URL"))
	cur, err := collection.Find(context.Background(), bson.D{})
	if err != nil{


		return err
	}
	defer cur.Close(context.Background())
	var result []model.Todo
	if err := cur.All(context.Background(), &result); err != nil{

       return err
	}
	
    
	return c.JSON(fiber.Map{"result": result, "user": username })
	   
 }
 
 
 
 func GetOne(c *fiber.Ctx) error{
 


	collection := db.HandleTodoDB(os.Getenv("DATABASE_URL"))

	var result model.Todo
    id, err := primitive.ObjectIDFromHex(c.Params("id"))

	if err != nil{

		return err
	}
	
	err = collection.FindOne(context.Background(), bson.D{primitive.E{Key:"_id", Value:id}}).Decode(&result)
	if err != nil{


		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Document was not found"})
	}
	
	
	if result.Id.IsZero(){

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "No document with such id"})
	}

    return c.JSON(fiber.Map{"result":result})
	 
 }
 
 
 func PostTodo(c *fiber.Ctx) error{
 
	 collection := db.HandleTodoDB(os.Getenv("DATABASE_URL"))


	 data := new(model.Todo)

	 err := c.BodyParser(data);
	 errorHandler.HandleError(err)
    
	 _, err = collection.InsertOne(context.Background(), bson.M{"title": data.Title, "content": data.Content})

	 if err != nil{

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "It was a failure"})
	 }
	 
	 return c.JSON(fiber.Map{"message": "It was a success"})


 }
 
 func DeleteTodo(c *fiber.Ctx) error{
 
     
	

	collection := db.HandleTodoDB(os.Getenv("DATABASE_URL"))
	id, err := primitive.ObjectIDFromHex(c.Params("id"))
	errorHandler.HandleError(err)
	var deletedDocument model.Todo
    err = collection.FindOneAndDelete(context.Background(), bson.D{primitive.E{Key:"_id", Value:id }}).Decode(&deletedDocument)

	if err != nil{


		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error deleting document"})
	}


	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": "document has been deleted", "document": deletedDocument})
 }



 func UpdateTodo(c *fiber.Ctx) error {
  

	collection := db.HandleTodoDB(os.Getenv("DATABASE_URL"))
	id, err := primitive.ObjectIDFromHex(c.Params("id"))

    errorHandler.HandleError(err)
	data := new(model.Todo)
	err = c.BodyParser(data)
	opts := options.FindOneAndUpdate().SetUpsert(true)

	filter := bson.D{primitive.E{Key:"_id", Value:id}}
    

	var update primitive.D
	if data.Title == "" && len(data.Content) > 0 {


         update = bson.D{{ Key: "$set", Value: bson.D{primitive.E{ Key:"content", Value:data.Content}} }}
	} else if len(data.Title) > 0 && data.Content == ""{

		update = bson.D{{ Key:"$set", Value: bson.D{primitive.E{Key: "title",Value:data.Title}} }}
	} else {
        
		update = bson.D{{ Key:"$set", Value: bson.D{primitive.E{Key: "title", Value:data.Title}, primitive.E{Key:"content", Value:data.Content}} }}

	}
	
	
	errorHandler.HandleError(err)

	var updatedDocument model.Todo
	err = collection.FindOneAndUpdate(context.Background(), filter, update, opts).Decode(&updatedDocument)
    
	if err != nil{

		return c.Status(fiber.StatusInternalServerError).JSON(fiber.Map{"message": "Error Updated document"})
	}

	return c.Status(fiber.StatusAccepted).JSON(fiber.Map{"message": "document has been updated", "document": updatedDocument})


 }