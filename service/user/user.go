package user

import (
	dbEntity "github.com/lukmanlukmin/wallet/entity/database"
	httpRequestEntity "github.com/lukmanlukmin/wallet/entity/http/request"
	httpResponseEntity "github.com/lukmanlukmin/wallet/entity/http/response"
	repository "github.com/lukmanlukmin/wallet/repository/database"
	helper "github.com/lukmanlukmin/wallet/util/helper"
)

type UserService struct {
	userRepository repository.UserRepositoryInterface
}

func UserServiceHandler() *UserService {
	return &UserService{
		userRepository: repository.UserRepositoryHandler(),
	}
}

type UserServiceInterface interface {
	CreateUser(userData httpRequestEntity.UserRequest) error
	GetUserByID(id int) (*httpResponseEntity.UserResponse, error)
}

func (service *UserService) GetUserByID(id int) (*httpResponseEntity.UserResponse, error) {
	userData := &dbEntity.User{}
	err := service.userRepository.GetUserByID(id, userData)

	result := &httpResponseEntity.UserResponse{}

	result.ID = userData.ID
	result.Email = userData.Email
	result.Username = userData.Username
	result.Password = "hidden content"
	return result, err
}

func (service *UserService) CreateUser(userData httpRequestEntity.UserRequest) error {
	insertData := &dbEntity.User{
		Username: userData.Username,
		Email:    userData.Email,
		Password: helper.HashString(userData.Password),
		UserType: userData.UserType,
	}
	error := service.userRepository.InsertUser(insertData)
	return error
}
