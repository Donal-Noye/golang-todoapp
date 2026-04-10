package adapters_in_transport_http

import (
	"time"

	"github.com/Donal-Noye/golang-todoapp/internal/core/domain"
	"github.com/google/uuid"
)

type TaskDTOResponse struct {
	ID           uuid.UUID  `json:"id"             example:"060551fb-5b8e-43dd-a7c2-374df4135e4c"`
	Version      int        `json:"version" example:"1"`
	Title        string     `json:"title" example:"Домашка"`
	Description  *string    `json:"description" example:"Сделать до четверга домашнее задание по математике"`
	Completed    bool       `json:"completed" example:"false"`
	CreatedAt    time.Time  `json:"created_at" example:"2020-09-20T14:00:00+09:00"`
	CompletedAt  *time.Time `json:"completed_at" example:"null"`
	AuthorUserId uuid.UUID  `json:"author_user_id" example:"550e8400-e29b-41d4-a716-446655440000"`
}

func taskDTOFromDomain(task domain.Task) TaskDTOResponse {
	return TaskDTOResponse{
		ID:           task.ID,
		Version:      task.Version,
		Title:        task.Title,
		Description:  task.Description,
		Completed:    task.Completed,
		CreatedAt:    task.CreatedAt,
		CompletedAt:  task.CompletedAt,
		AuthorUserId: task.AuthorUserId,
	}
}

func taskDTOsFromDomains(tasks []domain.Task) []TaskDTOResponse {
	dtos := make([]TaskDTOResponse, len(tasks))

	for i, task := range tasks {
		dtos[i] = taskDTOFromDomain(task)
	}

	return dtos
}
