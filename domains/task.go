package domains

import (
	"github.com/AnkitDhawale/TodoListApp/dto"
	"time"
)

type Task struct {
	Id          string    `db:"id"`
	UserId      string    `db:"user_id"`
	Title       string    `db:"title"`
	Description string    `db:"description"`
	DueDate     time.Time `db:"due_date"`
	Priority    string    `db:"priority_enum"`
	Category    string    `db:"category"`
	Status      string    `db:"status_enum"`
	CreatedAt   time.Time `db:"created_at"`
	UpdatedAt   time.Time `db:"updated_at"`
}

type TaskRepo interface {
	FindAllTaskByUserId(userIdFromToken string) ([]Task, error)
	CreateNewTask(task *Task) (string, error)
	UpdateTask(task *Task) error
	DeleteTask(id string) error
	FindTaskById(taskId string) (*Task, error)
	FindTaskByFilter(query string, values ...any) ([]Task, error)
}

func (t Task) ToDto() dto.Task {
	return dto.Task{
		TaskId:      t.Id,
		Title:       t.Title,
		Description: t.Description,
		DueDate:     t.DueDate,
		Priority:    t.Priority,
		Category:    t.Category,
		Status:      t.Status,
		CreatedAt:   t.CreatedAt,
		UpdatedAt:   t.UpdatedAt,
	}
}
