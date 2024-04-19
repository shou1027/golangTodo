package task

import "context"

type TaskRepository interface {
	Save(ctx context.Context, task *Task) error
	FindById(ctx context.Context, id string) (*Task, error)
	FindByUserId(ctx context.Context, userId string) ([]*Task, error)
}
