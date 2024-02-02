package repository

import "go.mongodb.org/mongo-driver/mongo"

const (
	userCollection = "users"
)

type userRepo struct {
	DB         mongo.Database
	Collection mongo.Collection
}

func NewUserRepo(DB mongo.Database) *userRepo {
	return &userRepo{
		DB,
		*DB.Collection(userCollection),
	}
}
