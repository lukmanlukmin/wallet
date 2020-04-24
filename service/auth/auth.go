package auth

import (
	"fmt"

	dbEntity "github.com/lukmanlukmin/wallet/entity/database"
	httpResponseEntity "github.com/lukmanlukmin/wallet/entity/http/response"
	repository "github.com/lukmanlukmin/wallet/repository/database"
	helper "github.com/lukmanlukmin/wallet/util/helper"
)

type AuthService struct {
	userRepository repository.UserRepositoryInterface
}

func AuthServiceHandler() *AuthService {
	return &AuthService{
		userRepository: repository.UserRepositoryHandler(),
	}
}

type AuthServiceInterface interface {
	GetUserByID(id int) (*httpResponseEntity.UserResponse, error)
	Login(email string, password string) (*httpResponseEntity.LoginResponse, error)
	Logout(ID int) error
}

func (service *AuthService) GetUserByID(id int) (*httpResponseEntity.UserResponse, error) {
	userData := &dbEntity.User{}
	err := service.userRepository.GetUserByID(id, userData)
	result := &httpResponseEntity.UserResponse{}
	result.ID = userData.ID
	result.Email = userData.Email
	result.Username = userData.Username
	result.Password = "hidden content"
	return result, err
}

func (service *AuthService) Login(email string, password string) (*httpResponseEntity.LoginResponse, error) {
	userData := &dbEntity.User{}
	err := service.userRepository.GetUserByEmail(email, userData)
	statusPassword := helper.CompareHash(userData.Password, password)
	if statusPassword == true {
		fmt.Println("match password")
	} else {
		fmt.Println("not match password")
	}
	token, err := helper.CreateToken(int(userData.ID), userData.UserType)
	tokenResult := &httpResponseEntity.LoginResponse{
		Token: token,
	}
	userData.Available = true
	go service.userRepository.UpdateUser(int(userData.ID), userData)
	return tokenResult, err
}

func (service *AuthService) Logout(ID int) error {
	userData := &dbEntity.User{}
	err := service.userRepository.GetUserByID(ID, userData)
	userData.Available = false
	go service.userRepository.UpdateUser(int(userData.ID), userData)
	return err
}
