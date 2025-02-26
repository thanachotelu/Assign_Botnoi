package model

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID           primitive.ObjectID `bson:"_id,omitempty" json:"user_id"`
	Username     string             `bson:"username" json:"username"`
	Passwordhash string             `bson:"password_hash" json:"password_hash"`
	Firstname    string             `bson:"firstname" json:"firstname"`
	Lastname     string             `bson:"lastname" json:"lastname"`
	Phonenumber  string             `bson:"phonenumber" json:"phonenumber"`
	Email        string             `bson:"email" json:"email"`
	Role         string             `bson:"role" json:"role"`
	CreatedAt    time.Time          `bson:"created_at" json:"created_at"`
	UpdatedAt    time.Time          `bson:"updated_at" json:"updated_at"`
	Status       string             `bson:"status" json:"status"`
}

type NewUser struct {
	Username    string `bson:"username" json:"username"`
	Password    string `bson:"password" json:"password"`
	Firstname   string `bson:"firstname" json:"firstname"`
	Lastname    string `bson:"lastname" json:"lastname"`
	Phonenumber string `bson:"phonenumber" json:"phonenumber"`
	Email       string `bson:"email" json:"email"`
	Role        string `bson:"role" json:"role"`
	Status      string `bson:"status" json:"status"`
}
