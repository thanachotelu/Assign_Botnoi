package services

import (
	"context"
	"errors"
	"log"
	"time"

	"crud/internal/model"
	"crud/internal/repository"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"golang.org/x/crypto/bcrypt"
)

type UserService struct {
	userRepo *repository.UserRepository
}

func NewUserService(userRepo *repository.UserRepository) *UserService {
	return &UserService{userRepo: userRepo}
}

// hashPassword แฮชรหัสผ่านก่อนบันทึก
func hashPassword(password string) (string, error) {
	hashedPassword, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
	if err != nil {
		return "", err
	}
	return string(hashedPassword), nil
}

// AddUser เพิ่มผู้ใช้ใหม่
func (s *UserService) AddUser(ctx context.Context, newUser model.NewUser) (model.User, error) {
	log.Println("Raw Password:", newUser.Password)

	hashedPassword, err := hashPassword(newUser.Password)
	if err != nil {
		return model.User{}, err
	}
	log.Println("Hashed Password Before Insert:", hashedPassword)

	user := model.User{
		Username:     newUser.Username,
		Passwordhash: hashedPassword,
		Firstname:    newUser.Firstname,
		Lastname:     newUser.Lastname,
		Phonenumber:  newUser.Phonenumber,
		Email:        newUser.Email,
		Role:         newUser.Role,
		Status:       newUser.Status,
		CreatedAt:    time.Now(),
		UpdatedAt:    time.Now(),
	}

	// บันทึกลงฐานข้อมูล
	insertedID, err := s.userRepo.AddUser(ctx, user)
	if err != nil {
		return model.User{}, err
	}
	user.ID = insertedID

	return user, nil
}

// GetAllUsers ดึงข้อมูลผู้ใช้ทั้งหมด
func (s *UserService) GetAllUsers(ctx context.Context) ([]model.User, error) {
	return s.userRepo.GetAllUsers(ctx)
}

// GetUserById ค้นหาผู้ใช้ตาม ID
func (s *UserService) GetUserById(ctx context.Context, userID string) (model.User, error) {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return model.User{}, errors.New("invalid user ID format")
	}
	return s.userRepo.GetUserById(ctx, objID)
}

// UpdateUser อัปเดตข้อมูลผู้ใช้
func (s *UserService) UpdateUser(ctx context.Context, userID string, updatedUser model.User) error {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("invalid user ID format")
	}
	updateFields := bson.M{}
	if updatedUser.Username != "" {
		updateFields["username"] = updatedUser.Username
	}
	if updatedUser.Passwordhash != "" {
		updateFields["password_hash"] = updatedUser.Passwordhash
	}
	if updatedUser.Firstname != "" {
		updateFields["firstname"] = updatedUser.Firstname
	}
	if updatedUser.Lastname != "" {
		updateFields["lastname"] = updatedUser.Lastname
	}
	if updatedUser.Phonenumber != "" {
		updateFields["phonenumber"] = updatedUser.Phonenumber
	}
	if updatedUser.Email != "" {
		updateFields["email"] = updatedUser.Email
	}
	if updatedUser.Role != "" {
		updateFields["role"] = updatedUser.Role
	}
	if updatedUser.Status != "" {
		updateFields["status"] = updatedUser.Status
	}

	updateFields["updated_at"] = time.Now()

	if len(updateFields) == 0 {
		return errors.New("no valid fields to update")
	}

	return s.userRepo.UpdateUser(ctx, objID, updateFields)
}

// DeleteUser ลบผู้ใช้
func (s *UserService) DeleteUser(ctx context.Context, userID string) error {
	objID, err := primitive.ObjectIDFromHex(userID)
	if err != nil {
		return errors.New("invalid user ID format")
	}
	return s.userRepo.DeleteUser(ctx, objID)
}
