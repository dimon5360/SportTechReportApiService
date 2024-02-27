package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Report struct {
	UserId     uint64
	Document   string
	Created_at time.Time
	Updated_at time.Time
}

type GetReportMongoModel struct {
	ID        primitive.ObjectID `bson:"_id"`
	UserId    uint64             `bson:"user_id"`
	Document  string             `bson:"report"`
	CreatedAt time.Time          `bson:"created_at"`
	UpdatedAt time.Time          `bson:"updated_at"`
}

type PutReportMongoModel struct {
	UserId    uint64    `bson:"user_id"`
	Document  string    `bson:"report,omitempty"`
	CreatedAt time.Time `bson:"created_at"`
	UpdatedAt time.Time `bson:"updated_at"`
}
