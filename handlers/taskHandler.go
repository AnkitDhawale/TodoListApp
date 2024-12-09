package handlers

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/AnkitDhawale/TodoListApp/dto"
	"github.com/AnkitDhawale/TodoListApp/helpers"
	"github.com/AnkitDhawale/TodoListApp/middlewares"
	"github.com/AnkitDhawale/TodoListApp/services"
	"github.com/go-chi/chi/v5"
	"github.com/go-playground/validator/v10"
	"log"
	"net/http"
	"time"
)

type TaskHandler struct {
	Service    services.TaskService
	Validatorr *validator.Validate
}

// GetAllTasksOfAUser godoc
// @Summary Get all tasks of a user
// @Description Fetches all tasks associated with the authenticated user
// @Tags tasks
// @Produce json
// @Success 200 {array} dto.Task "List of tasks"
// @Failure 500 {object} helpers.Response "UserId not found in context or server error"
// @Router /todoapp/tasks [get]
func (th TaskHandler) GetAllTasksOfAUser(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(middlewares.UserIDKey)
	if id == nil {
		helpers.WriteResponse(w, http.StatusInternalServerError, nil, errors.New("userId not found in context"))
		return
	}
	userIdFromToken := id.(string)
	log.Println("GetAllTasksOfAUser: ", userIdFromToken)

	tasks, err := th.Service.GetAllTasksOfUser(userIdFromToken)
	if err != nil {
		helpers.WriteResponse(w, http.StatusInternalServerError, nil, err)
		return
	}
	if len(tasks) == 0 {
		helpers.WriteResponse(w, http.StatusOK, nil, nil)
		return
	} else {
		helpers.WriteResponse(w, http.StatusOK, tasks, nil)
		return
	}
}

// CreateNewTask godoc
// @Summary Creates a new task
// @Description Creates a new task for the authenticated user
// @Tags tasks
// @Accept json
// @Produce json
// @Param task body dto.TaskInputRequest true "Task input data"
// @Success 200 {string} string "New task created successfully"
// @Failure 400 {object} helpers.Response "Invalid request payload or error while creating task"
// @Router /todoapp/tasks [post]
func (th TaskHandler) CreateNewTask(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(middlewares.UserIDKey)
	userId := id.(string)

	var input dto.TaskInputRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err)
		helpers.WriteResponse(w, http.StatusBadRequest, nil, errors.New("invalid request payload"))
		return
	}

	// Parse due_date string into time.Time
	dueDate, err := time.Parse("2006-01-02 15:04:05", input.DueDate)
	if err != nil {
		log.Println(err)
		helpers.WriteResponse(w, http.StatusBadRequest, nil, errors.New("invalid due_date format, expected 'YYYY-MM-DD HH:mm:ss'"))
		return
	}

	input.SetDefaults()
	err = ValidateIncomingRequest(th, &input)
	if err != nil {
		helpers.WriteResponse(w, http.StatusBadRequest, nil, fmt.Errorf("input validation failed, err: %v", err))
		return
	}

	task := dto.Task{
		Title:       input.Title,
		Description: input.Description,
		DueDate:     dueDate,
		Priority:    input.Priority,
		Category:    input.Category,
		Status:      input.Status,
	}

	newTaskId, err := th.Service.CreateNewTask(userId, &task)
	if err != nil {
		helpers.WriteResponse(w, http.StatusBadRequest, nil, err)
		return
	} else {
		helpers.WriteResponse(w, http.StatusOK, fmt.Sprintf("new task created successfully with id: %s", newTaskId), err)
		return
	}
}

// UpdateTask godoc
// @Summary Updates an existing task
// @Description Updates only the provided fields of a specific task
// @Tags tasks
// @Accept json
// @Produce json
// @Param id path string true "Task ID"
// @Param task body dto.TaskInputRequest true "Task input data"
// @Success 200 {string} string "Task updated successfully"
// @Failure 400 {object} helpers.Response "Invalid request payload or task ID missing"
// @Router /todoapp/tasks/{id} [put]
func (th TaskHandler) UpdateTask(w http.ResponseWriter, r *http.Request) {
	taskIdFromParam := chi.URLParam(r, "id")
	if taskIdFromParam == "" {
		helpers.WriteResponse(w, http.StatusBadRequest, nil, errors.New("task id missing from request"))
		return
	}

	var input dto.TaskInputRequest
	if err := json.NewDecoder(r.Body).Decode(&input); err != nil {
		log.Println(err)
		helpers.WriteResponse(w, http.StatusBadRequest, nil, errors.New("invalid request payload"))
		return
	}

	// Parse due_date string into time.Time
	dueDate, err := time.Parse("2006-01-02 15:04:05", input.DueDate)
	if err != nil {
		log.Println(err)
		helpers.WriteResponse(w, http.StatusBadRequest, nil, errors.New("invalid due_date format, expected 'YYYY-MM-DD HH:mm:ss'"))
		return
	}

	task := dto.Task{
		Title:       input.Title,
		Description: input.Description,
		DueDate:     dueDate,
		Priority:    input.Priority,
		Category:    input.Category,
		Status:      input.Status,
	}

	task.TaskId = taskIdFromParam

	err = th.Service.UpdateTask(&task)
	if err != nil {
		helpers.WriteResponse(w, http.StatusBadRequest, nil, err)
		return
	} else {
		helpers.WriteResponse(w, http.StatusOK, "task updated successfully...", err)
		return
	}
}

// DeleteTask godoc
// @Summary Deletes a task
// @Description Deletes a specific task by its ID
// @Tags tasks
// @Produce json
// @Param id path string true "Task ID"
// @Success 200 {string} string "Task deleted successfully"
// @Failure 400 {object} helpers.Response "Task ID missing or deletion failed"
// @Router /todoapp/tasks/{id} [delete]
func (th TaskHandler) DeleteTask(w http.ResponseWriter, r *http.Request) {
	taskIdFromParam := chi.URLParam(r, "id")
	if taskIdFromParam == "" {
		helpers.WriteResponse(w, http.StatusBadRequest, nil, errors.New("task id missing from request"))
		return
	}

	err := th.Service.DeleteTask(taskIdFromParam)
	if err != nil {
		helpers.WriteResponse(w, http.StatusBadRequest, nil, err)
		return
	} else {
		helpers.WriteResponse(w, http.StatusOK, "task deleted successfully...", err)
		return
	}
}

// ViewTasksWithFilter godoc
// @Summary View tasks with filter options
// @Description Fetches tasks with optional filtering by due date, priority, category, and status
// @Tags tasks
// @Produce json
// @Param due_date query string false "Filter tasks by due date (format: YYYY-MM-DD)"
// @Param priority query string false "Filter tasks by priority (e.g., Low, Medium, High)"
// @Param category query string false "Filter tasks by category"
// @Param status query string false "Filter tasks by status (e.g., Pending, Completed)"
// @Success 200 {array} dto.Task "Filtered list of tasks"
// @Failure 500 {object} helpers.Response "Internal server error"
// @Router /todoapp/tasks/view [get]
func (th TaskHandler) ViewTasksWithFilter(w http.ResponseWriter, r *http.Request) {
	id := r.Context().Value(middlewares.UserIDKey)
	if id == nil {
		helpers.WriteResponse(w, http.StatusInternalServerError, nil, errors.New("userId not found in context"))
		return
	}
	userIdFromToken := id.(string)

	// get option filters from query parameters
	dueDate := r.URL.Query().Get("due_date")
	priority := r.URL.Query().Get("priority")
	category := r.URL.Query().Get("category")
	status := r.URL.Query().Get("status")

	log.Println("filters: ", dueDate, priority, category, status)

	// Fetch tasks from service with filters
	tasks, err := th.Service.ViewTaskWithFilter(userIdFromToken, dueDate, priority, category, status)
	if err != nil {
		helpers.WriteResponse(w, http.StatusInternalServerError, nil, errors.New("error while fetching tasks"))
		return
	} else {
		helpers.WriteResponse(w, http.StatusOK, tasks, nil)
		return
	}
}

func ValidateIncomingRequest(th TaskHandler, input *dto.TaskInputRequest) error {
	err := th.Validatorr.Struct(input)
	if err != nil {
		log.Println("input validation err: ", err)
		return err
	}
	return nil
}
