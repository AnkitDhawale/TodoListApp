package dto

const (
	DEFAULT_PRIORITY = "Medium"
	DEFAULT_STATUS   = "Pending"
)

type TaskInputRequest struct {
	Title       string `json:"title" validate:"required"`
	Description string `json:"description"`
	DueDate     string `json:"due_date"`
	Priority    string `json:"priority"`
	Category    string `json:"category"`
	Status      string `json:"status"`
}

// SetDefaults handles default values for priority and status fields
func (input *TaskInputRequest) SetDefaults() {
	if input.Priority == "" {
		input.Priority = DEFAULT_PRIORITY
	}

	if input.Status == "" {
		input.Status = DEFAULT_STATUS
	}
}
