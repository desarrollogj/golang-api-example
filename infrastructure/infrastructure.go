package infrastructure

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

// Repository entities

type MongoUser struct {
	ID          primitive.ObjectID `bson:"_id"`
	Reference   string             `bson:"reference"`
	FirstName   string             `bson:"first_name"`
	LastName    string             `bson:"last_name"`
	Email       string             `bson:"email"`
	IsActive    bool               `bson:"is_active"`
	CreatedDate time.Time          `bson:"created_date"`
	UpdatedDate time.Time          `bson:"updated_date"`
}
