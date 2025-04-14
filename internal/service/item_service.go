package service

import (
	"context"

	"github.com/lipaysamart/go-todolist-api-execrices/internal/model"
	"github.com/lipaysamart/go-todolist-api-execrices/internal/repository"
	"github.com/lipaysamart/go-todolist-api-execrices/pkg/utils"
)

type ITaskService interface {
	AddItem(ctx context.Context, req *model.ItemReq) error
	DelItem(ctx context.Context, req *model.ItemReq) error
	UpdateItem(ctx context.Context, req *model.ItemReq) (*model.Item, error)
	GetItem(ctx context.Context, id string) (*model.Item, error)
	GetItemList(ctx context.Context) ([]model.Item, error)
}

type TaskService struct {
	taskRepo repository.ITaskRepo
}

func NewTaskService(repo repository.ITaskRepo) *TaskService {
	return &TaskService{
		taskRepo: repo,
	}
}

func (s *TaskService) AddItem(ctx context.Context, req *model.ItemReq) error {
	var item model.Item

	utils.Copy(&item, req)

	if err := s.taskRepo.Create(ctx, &item); err != nil {
		return err
	}

	return nil
}

func (s *TaskService) DelItem(ctx context.Context, req *model.ItemReq) error {
	var item model.Item

	utils.Copy(&item, req)

	if err := s.taskRepo.Delete(ctx, &item); err != nil {
		return err
	}

	return nil
}

func (s *TaskService) UpdateItem(ctx context.Context, req *model.ItemReq) (*model.Item, error) {
	var item model.Item

	utils.Copy(&item, req)

	updatedItem, err := s.taskRepo.Update(ctx, &item)
	if err != nil {
		return nil, err
	}

	return updatedItem, nil
}

func (s *TaskService) GetItem(ctx context.Context, id string) (*model.Item, error) {
	item, err := s.taskRepo.FindItemByID(ctx, id)
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (s *TaskService) GetItemList(ctx context.Context) ([]model.Item, error) {
	items, err := s.taskRepo.Find(ctx)
	if err != nil {
		return nil, err
	}
	return items, nil
}
