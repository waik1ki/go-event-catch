package repository

import (
	"context"
	"go-smart-contract-event-catch/env"

	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type Repository struct {
	env *env.Env

	client *mongo.Client
	db     *mongo.Database

	Tx            *mongo.Collection
	NFT           *mongo.Collection
	NFTCollection *mongo.Collection
}

func NewRepository(env *env.Env) (*Repository, error) {
	r := &Repository{env: env}

	var err error
	ctx := context.Background()

	if r.client, err = mongo.Connect(ctx, options.Client().ApplyURI(env.MongoDB.Uri)); err != nil {
		return nil, err
	} else if err = r.client.Ping(ctx, nil); err != nil {
		return nil, err
	} else {
		r.db = r.client.Database(env.MongoDB.Database)

		r.NFTCollection = r.db.Collection(env.MongoDB.Collection.NFTCollection)
		r.NFT = r.db.Collection(env.MongoDB.Collection.NFT)
		r.Tx = r.db.Collection(env.MongoDB.Collection.Tx)
	}

	return r, nil
}
