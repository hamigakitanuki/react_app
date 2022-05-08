package graph

// This file will be automatically regenerated based on the schema, any resolver implementations
// will be copied through when generating and any unknown code will be moved to the end.

import (
	"context"
	"fmt"
	"time"

	"github.com/yuuu/with-apollo-api/graph/generated"
	"github.com/yuuu/with-apollo-api/graph/model"
)

func (r *mutationResolver) CreateTask(ctx context.Context, input model.NewTask) (*model.Task, error) {
	panic(fmt.Errorf("not implemented"))
	// ここから追記
	timestamp := time.Now().Format("2006-01-02 15:04:05")

	task := model.Task{
		Title:     input.Title,
		Note:      input.Note,
		Completed: 0,
		CreatedAt: timestamp,
		UpdatedAt: timestamp,
	}
	r.DB.Create(&task)

	return &task, nil
	// ここまで追記
}

func (r *queryResolver) Tasks(ctx context.Context) ([]*model.Task, error) {
	// ここから追記
	tasks := []*model.Task{}

	r.DB.Find(&tasks)

	return tasks, nil
	// ここまで追記
}

// Mutation returns generated.MutationResolver implementation.
func (r *Resolver) Mutation() generated.MutationResolver { return &mutationResolver{r} }

// Query returns generated.QueryResolver implementation.
func (r *Resolver) Query() generated.QueryResolver { return &queryResolver{r} }

type mutationResolver struct{ *Resolver }
type queryResolver struct{ *Resolver }
