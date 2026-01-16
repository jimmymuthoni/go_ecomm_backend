package store

import (
	"context"

	"github.com/jimmymuthoni/go_ecomm_backend/models"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)


type MongoProductStore struct {
	db 		*MongoProductStore
	coll    string
}


func NewMongoProductStore(db *mongo.Database) *MongoProductStore {
	return &MongoProductStore{
		db:    db,
		coll: "products",
	}
}


func (s *MongoProductStore) Insert(ctx context.Context, p *models.Product) error {
	res, err := s.db.Collection(s.coll).InsertOne(ctx, p)
	if err != nil {
		return err
	}
	p.ID = res.InsertOne.(primitive.ObjectID)
	return err
}