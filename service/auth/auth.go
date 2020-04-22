package auth

import (
	httpEntity "example_app/entity/http"
	repositoryAPI "example_app/repository/api"
	repository "example_app/repository/db"
	"sync"
)

type UserService struct {
	userRepository    repository.UserRepositoryInterface
	userRepositoryAPI repositoryAPI.FriendAPIRepositoryInterface
}

func UserServiceHandler() *UserService {
	return &UserService{
		userRepository:    repository.UserRepositoryHandler(),
		userRepositoryAPI: repositoryAPI.ThirdPartyAPIHandler(),
	}
}

type UserServiceInterface interface {
	GetUserByID(id int) *httpEntity.UserDetailResponse
	Login(email string, password string) []httpEntity.UserResponse
}

func (service *UserService) GetUserByID(id int, waitGroup *sync.WaitGroup) *httpEntity.UserDetailResponse {
	result := &httpEntity.UserDetailResponse{}
	return result
}

func (service *UserService) Login(id int, waitGroup *sync.WaitGroup) *httpEntity.UserDetailResponse {
	result := &httpEntity.UserDetailResponse{}
	return result
}
