package database

import (
	"context"
	"fmt"
	"log"
	"time"

	"crud/internal/config"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

type MongoDB struct {
	client     *mongo.Client
	database   *mongo.Database
	collection *mongo.Collection
	uri        string
}

func NewMongoDB(cfg *config.Config) (*MongoDB, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	client, err := mongo.Connect(ctx, options.Client().ApplyURI(cfg.MongoURI))
	if err != nil {
		return nil, fmt.Errorf("failed to connect to MongoDB: %w", err)
	}

	// ตรวจสอบการเชื่อมต่อ
	if err := client.Ping(ctx, nil); err != nil {
		_ = client.Disconnect(ctx)
		return nil, fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	db := client.Database(cfg.MongoDBName)

	collectionName := "users"
	if err := CollectionExists(db, collectionName); err != nil {
		return nil, fmt.Errorf("failed to create collection: %w", err)
	}

	log.Println("MongoDB connected successfully!")

	return &MongoDB{
		client:     client,
		database:   db,
		collection: db.Collection(collectionName),
	}, nil
}

func (db *MongoDB) Reconnect() error {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	// ปิดการเชื่อมต่อเดิม
	if err := db.client.Disconnect(ctx); err != nil {
		return fmt.Errorf("failed to disconnect from MongoDB: %w", err)
	}

	// เชื่อมต่อใหม่
	client, err := mongo.Connect(ctx, options.Client().ApplyURI(db.uri))
	if err != nil {
		return fmt.Errorf("failed to reconnect to MongoDB: %w", err)
	}

	// ตรวจสอบการเชื่อมต่อ
	if err = client.Ping(ctx, nil); err != nil {
		_ = client.Disconnect(ctx)
		return fmt.Errorf("failed to ping MongoDB: %w", err)
	}

	db.client = client
	db.database = client.Database(db.database.Name())

	return nil
}

func (db *MongoDB) Close() error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	return db.client.Disconnect(ctx)
}

func (db *MongoDB) GetCollection() *mongo.Collection {
	return db.collection
}

func CollectionExists(db *mongo.Database, collectionName string) error {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	collections, err := db.ListCollectionNames(ctx, bson.M{})
	if err != nil {
		return err
	}

	for _, name := range collections {
		if name == collectionName {
			return nil
		}
	}

	return db.CreateCollection(ctx, collectionName)
}
