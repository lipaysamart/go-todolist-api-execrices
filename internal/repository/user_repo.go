package repository

import (
	"context"

	"github.com/lipaysamart/go-todolist-api-execrices/internal/model"
	"github.com/lipaysamart/go-todolist-api-execrices/pkg/db"
)

type IUserRepo interface {
	Create(ctx context.Context, user *model.User) error
	Update(ctx context.Context, user *model.User) (*model.User, error)
	Find(ctx context.Context) ([]model.User, error)
	FindUserByEmail(ctx context.Context, email string) (*model.User, error)
	FindUserByID(ctx context.Context, id string) (*model.User, error)
	Delete(ctx context.Context, user *model.User) error
}

type UserRepository struct {
	database db.IDatabase
}

func NewUserRepository(db db.IDatabase) *UserRepository {
	return &UserRepository{
		database: db,
	}
}

func (r *UserRepository) Create(ctx context.Context, user *model.User) error {
	return r.database.Create(ctx, user)
}

func (r *UserRepository) Update(ctx context.Context, user *model.User) (*model.User, error) {
	if err := r.database.Save(ctx, user); err != nil {
		return nil, err
	}

	updatedUser, err := r.FindUserByID(ctx, user.ID)
	if err != nil {
		return nil, err
	}

	return updatedUser, nil
}

func (r *UserRepository) Find(ctx context.Context) ([]model.User, error) {
	var user []model.User

	if err := r.database.Find(ctx, user); err != nil {
		return nil, err
	}

	return user, nil
}

func (r *UserRepository) FindUserByEmail(ctx context.Context, email string) (*model.User, error) {
	var user model.User

	query := db.BuildQuery("email = ?", email)
	if err := r.database.FindOne(ctx, &user, db.WithQuery(query)); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) FindUserByID(ctx context.Context, id string) (*model.User, error) {
	var user model.User

	if err := r.database.FindByID(ctx, id, &user); err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *UserRepository) Delete(ctx context.Context, user *model.User) error {
	return r.database.Delete(ctx, user)
}
