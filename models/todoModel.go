package model

import "go.mongodb.org/mongo-driver/bson/primitive"

type Todo struct {

	Id    primitive.ObjectID `bson:"_id" json:"id"`
	Title string        `bson:"title" json:"title"`

	Content string `bson:"content" json:"content"`

	User_email string `bson:"user_email" json:"user_email"`

}
