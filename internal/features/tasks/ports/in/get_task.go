package ports_in

import (
	"github.com/Donal-Noye/golang-todoapp/internal/core/domain"
	"github.com/google/uuid"
)

// GetTaskParams / GetTaskResult — параметры и результат входящего порта для получения задачи.
// Принадлежат ядру (уровню сервиса): описывают, что сервис ожидает от любого входящего адаптера.
type GetTaskParams struct {
	ID uuid.UUID
}

func NewGetTaskParams(
	id uuid.UUID,
) GetTaskParams {
	return GetTaskParams{
		ID: id,
	}
}

type GetTaskResult struct {
	Task domain.Task
}

func NewGetTaskResult(
	task domain.Task,
) GetTaskResult {
	return GetTaskResult{
		Task: task,
	}
}
