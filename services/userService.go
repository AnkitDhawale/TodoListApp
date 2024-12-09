package services

import (
	"errors"
	"github.com/AnkitDhawale/TodoListApp/domains"
	"github.com/AnkitDhawale/TodoListApp/dto"
	"github.com/AnkitDhawale/TodoListApp/helpers"
	"github.com/google/uuid"
	"log"
	"time"
)

type UserService interface {
	GetAllUsers() ([]dto.User, error)
	CreatNewUser(dto.User) (string, error)
	UpdateUser(string, dto.User) error
}

type DefaultUserService struct {
	userRepo domains.UserRepo
}

func NewDefaultUserService(repo domains.UserRepo) *DefaultUserService {
	return &DefaultUserService{userRepo: repo}
}

func (service DefaultUserService) GetAllUsers() ([]dto.User, error) {
	time.Sleep(1000 * time.Millisecond)
	data, err := service.userRepo.FindAll()
	if err != nil {
		return nil, err
	}
	users := make([]dto.User, 0, len(data))

	for _, k := range data {
		users = append(users, k.ToDto())
	}

	return users, nil
}

func (service DefaultUserService) CreatNewUser(userInput dto.User) (string, error) {
	// check email in correct format

	encryptedPassword, err := helpers.EncryptPassword(userInput.Password)
	if err != nil {
		log.Println(err)
		return "", errors.New("error while generating bcrypt password")
	}

	user := &domains.User{
		Id:           uuid.New().String(),
		Email:        userInput.Email,
		PasswordHash: encryptedPassword,
		CreatedAt:    time.Now(),
	}

	newId, err := service.userRepo.AddUser(user)
	if err != nil {
		return "", err
	}

	return newId, nil
}

func (service DefaultUserService) UpdateUser(userIdFromToken string, user dto.User) error {
	encryptedPassword, err := helpers.EncryptPassword(user.Password)
	if err != nil {
		log.Println(err)
		return errors.New("error while generating bcrypt password")
	}

	updatedValues := domains.User{
		Email:        user.Email,
		PasswordHash: encryptedPassword,
	}

	err = service.userRepo.UpdateUser(userIdFromToken, &updatedValues)
	if err != nil {
		return err
	}

	return nil
}
