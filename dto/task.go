package dto

import (
	"time"
)

type Task struct {
	TaskId      string    `json:"task_id,omitempty"`
	Title       string    `json:"title,required"`
	Description string    `json:"description"`
	DueDate     time.Time `json:"due_date"`
	Priority    string    `json:"priority"`
	Category    string    `json:"category"`
	Status      string    `json:"status"`
	CreatedAt   time.Time `json:"created_at,omitempty"`
	UpdatedAt   time.Time `json:"updated_at,omitempty"`
}
