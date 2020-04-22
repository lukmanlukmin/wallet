package auth

import (
	"fmt"

	dbEntity "github.com/lukmanlukmin/wallet/entity/database"
	httpResponseEntity "github.com/lukmanlukmin/wallet/entity/http/response"
	repository "github.com/lukmanlukmin/wallet/repository/database"
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
	GetUserByID(id int) (*httpResponseEntity.UserResponse, error)
	Login(email string, password string) (*httpResponseEntity.LoginResponse, error)
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

func (service *UserService) Login(email string, password string) (*httpResponseEntity.LoginResponse, error) {
	userData := &dbEntity.User{}
	err := service.userRepository.GetUserByEmailPassword(email, password, userData)
	result := &httpResponseEntity.LoginResponse{}
	result.Token = fmt.Sprintf("%s-%s", userData.Username, userData.Email)
	return result, err
}
