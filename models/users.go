package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `json:"_id,omitempty" bson:"_id,omitempty"`
	Name     string             `json:"name" bson:"name,omitempty"`
	Age      int                `json:"age"`
	Password string             `json:"password,omitempty" `
	Email    string             `json:"email" bson:"email"`
}
