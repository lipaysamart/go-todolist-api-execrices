package service

import (
	"context"

	"github.com/lipaysamart/go-todolist-api-execrices/internal/model"
	"github.com/lipaysamart/go-todolist-api-execrices/internal/repository"
	"github.com/lipaysamart/go-todolist-api-execrices/pkg/utils"
)

type IUserService interface {
	Register(ctx context.Context, req *model.UserRegisterReq) error
	Login(ctx context.Context, req *model.UserLoginReq) error
	UpdateProfile(ctx context.Context, req *model.Item) (*model.Item, error)
}

type UserService struct {
	userRepo repository.UserRepository
}

func NewUserService(repo repository.UserRepository) *UserService {
	return &UserService{
		userRepo: repo,
	}
}

func (s *UserService) Register(ctx context.Context, req *model.UserRegisterReq) error {
	var user model.User

	utils.Copy(user, req)

	if err := s.userRepo.Create(ctx, &user); err != nil {
		return err
	}

	return nil
}
