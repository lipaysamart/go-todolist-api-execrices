package service

import (
	"context"

	"github.com/lipaysamart/go-todolist-api-execrices/internal/model"
	"github.com/lipaysamart/go-todolist-api-execrices/internal/repository"
	"github.com/lipaysamart/go-todolist-api-execrices/pkg/utils"
)

type IUserService interface {
	Register(ctx context.Context, req *model.UserRegisterReq) error
	Login(ctx context.Context, req *model.UserLoginReq) (*model.User, error)
	UpdateProfile(ctx context.Context, id string, req *model.UserRegisterReq) (*model.User, error)
}

type UserService struct {
	userRepo repository.IUserRepo
}

func NewUserService(repo repository.IUserRepo) *UserService {
	return &UserService{
		userRepo: repo,
	}
}

func (s *UserService) Register(ctx context.Context, req *model.UserRegisterReq) error {
	var user model.User

	utils.Copy(&user, req)

	if err := s.userRepo.Create(ctx, &user); err != nil {
		return err
	}

	return nil
}

func (s *UserService) Login(ctx context.Context, req *model.UserLoginReq) (*model.User, error) {
	var user model.User

	utils.Copy(&user, req)

	resp, err := s.userRepo.FindUserByEmail(ctx, req.Email)
	if err != nil {
		return nil, err
	}

	return resp, nil
}

func (s *UserService) UpdateProfile(ctx context.Context, id string, req *model.UserRegisterReq) (*model.User, error) {
	user, err := s.userRepo.FindUserByID(ctx, id)
	if err != nil {
		return nil, err
	}

	utils.Copy(&user, req)

	resp, err := s.userRepo.Update(ctx, user)
	if err != nil {
		return nil, err
	}

	return resp, nil
}
