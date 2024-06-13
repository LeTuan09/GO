package service

import "GO-API-SERVER/internal/repo"

type UserService struct {
	UserRepo *repo.UserRepo
}

func NewUserService() *UserService {
	return &UserService{
		UserRepo: repo.NewUserRepo(),
	}
}

func (us *UserService) GetByUserService() string {
	return us.UserRepo.GetByUserRepo()
}