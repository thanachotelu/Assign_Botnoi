package repository

import (
	"context"
	"fmt"
	"log"

	"crud/internal/model"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
)

// UserRepository จัดการการทำงานกับ MongoDB
type UserRepository struct {
	collection *mongo.Collection
}

// NewUserRepository สร้างอินสแตนซ์ของ UserRepository
func NewUserRepository(db *mongo.Database) *UserRepository {
	return &UserRepository{
		collection: db.Collection("users"),
	}
}

// AddUser เพิ่มผู้ใช้ใหม่
func (r *UserRepository) AddUser(ctx context.Context, user model.User) (primitive.ObjectID, error) {
	log.Println("Saving to MongoDB - Hashed Password:", user.Passwordhash)
	result, err := r.collection.InsertOne(ctx, user)
	if err != nil {
		return primitive.NilObjectID, fmt.Errorf("failed to insert user: %w", err)
	}
	insertedID, ok := result.InsertedID.(primitive.ObjectID)
	if !ok {
		return primitive.NilObjectID, fmt.Errorf("failed to get inserted ID")
	}
	log.Println("User inserted with hashed password:", user.Passwordhash)
	return insertedID, nil
}

// GetAllUsers ดึงข้อมูลผู้ใช้ทั้งหมด
func (r *UserRepository) GetAllUsers(ctx context.Context) ([]model.User, error) {
	var users []model.User
	cursor, err := r.collection.Find(ctx, bson.M{})
	if err != nil {
		return nil, fmt.Errorf("failed to get users: %w", err)
	}
	defer cursor.Close(ctx)

	for cursor.Next(ctx) {
		var user model.User
		if err := cursor.Decode(&user); err != nil {
			return nil, fmt.Errorf("failed to decode user: %w", err)
		}
		users = append(users, user)
	}

	if err := cursor.Err(); err != nil {
		return nil, fmt.Errorf("cursor error: %w", err)
	}

	return users, nil
}

// GetUserById ดึงข้อมูลผู้ใช้ตาม ID
func (r *UserRepository) GetUserById(ctx context.Context, userID primitive.ObjectID) (model.User, error) {
	var user model.User
	err := r.collection.FindOne(ctx, bson.M{"_id": userID}).Decode(&user)
	if err == mongo.ErrNoDocuments {
		return model.User{}, fmt.Errorf("user not found")
	}
	if err != nil {
		return model.User{}, fmt.Errorf("failed to get user: %w", err)
	}
	return user, nil
}

// UpdateUser อัปเดตข้อมูลผู้ใช้
func (r *UserRepository) UpdateUser(ctx context.Context, userID primitive.ObjectID, updatedUser model.User) error {
	filter := bson.M{"_id": userID}
	update := bson.M{"$set": updatedUser}

	result, err := r.collection.UpdateOne(ctx, filter, update)
	if err != nil {
		return fmt.Errorf("failed to update user: %w", err)
	}
	if result.MatchedCount == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}

// DeleteUser ลบผู้ใช้
func (r *UserRepository) DeleteUser(ctx context.Context, userID primitive.ObjectID) error {
	result, err := r.collection.DeleteOne(ctx, bson.M{"_id": userID})
	if err != nil {
		return fmt.Errorf("failed to delete user: %w", err)
	}
	if result.DeletedCount == 0 {
		return fmt.Errorf("user not found")
	}
	return nil
}
