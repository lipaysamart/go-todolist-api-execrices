package repository

import (
	"context"

	"github.com/lipaysamart/go-todolist-api-exerices/internal/model"
	"github.com/lipaysamart/gocommon/dbs"
)

type ITaskRepo interface {
	Create(ctx context.Context, item *model.Item) error
	FindItemByID(ctx context.Context, id string) (*model.Item, error)
	Find(ctx context.Context) ([]model.Item, error)
	Delete(ctx context.Context, id string) error
	Update(ctx context.Context, item *model.Item) (*model.Item, error)
}

type TaskRepository struct {
	database dbs.IDatabase
}

func NewTaskRepository(db dbs.IDatabase) *TaskRepository {
	return &TaskRepository{
		database: db,
	}
}

func (r *TaskRepository) Create(ctx context.Context, item *model.Item) error {
	return r.database.Create(ctx, item)
}

func (r *TaskRepository) Find(ctx context.Context) ([]model.Item, error) {
	var items []model.Item

	if err := r.database.Find(ctx, &items); err != nil {
		return nil, err
	}
	return items, nil
}

func (r *TaskRepository) Update(ctx context.Context, item *model.Item) (*model.Item, error) {
	if err := r.database.Save(ctx, item); err != nil {
		return nil, err
	}

	query := dbs.BuildQuery("title = ?", item.Title)

	err := r.database.FindOne(ctx, item, dbs.WithQuery(query))
	if err != nil {
		return nil, err
	}

	return item, nil
}

func (r *TaskRepository) Delete(ctx context.Context, id string) error {
	item, err := r.FindItemByID(ctx, id)
	if err != nil {
		return err
	}

	return r.database.Delete(ctx, item)
}

func (r *TaskRepository) FindItemByID(ctx context.Context, id string) (*model.Item, error) {
	var item model.Item

	if err := r.database.FindByID(ctx, id, &item); err != nil {
		return nil, err
	}

	return &item, nil
}
