package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	Id       primitive.ObjectID `json:"_id" bson:"_id,omit"`
	Name     string             `json:"name"`
	Age      int                `json:"age omitempty"`
	Password string             `json:"password omitempty" `
	Email    string             `json:"email omitempty" `
}
