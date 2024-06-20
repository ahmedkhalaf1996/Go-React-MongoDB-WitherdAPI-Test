package models

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
)

type User struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username" validate:"required"`
	Email    string             `bson:"email" validate:"required,email"`
	Password string             `bson:"password" validate:"required,min=6"`
	Lat      float64            `bson:"lat,omitempty" validate:"required"`
	Lon      float64            `bson:"lon,omitempty" validate:"required"`
}

type PublicUser struct {
	ID       primitive.ObjectID `bson:"_id,omitempty"`
	Username string             `bson:"username"`
	Email    string             `bson:"email"`
	Lat      float64            `bson:"lat,omitempty"`
	Lon      float64            `bson:"lon,omitempty"`
}

func (u *User) SanitizeUser() PublicUser {
	return PublicUser{
		ID:       u.ID,
		Username: u.Username,
		Email:    u.Email,
		Lat:      u.Lat,
		Lon:      u.Lon,
	}
}
