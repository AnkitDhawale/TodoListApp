package services

import (
	"errors"
	"fmt"
	"github.com/AnkitDhawale/TodoListApp/domains"
	"github.com/AnkitDhawale/TodoListApp/dto"
	"github.com/google/uuid"
	"log"
	"time"
)

type TaskService interface {
	GetAllTasksOfUser(userIdFromToken string) ([]dto.Task, error)
	CreateNewTask(userIdFromToken string, task *dto.Task) (string, error)
	UpdateTask(task *dto.Task) error
	DeleteTask(id string) error
	ViewTaskWithFilter(userIdFromToken string, dueDate string, priority string, category string, status string) ([]dto.Task, error)
}

type DefaultTaskService struct {
	repo domains.TaskRepo
}

func NewDefaultTaskService(repo domains.TaskRepo) *DefaultTaskService {
	return &DefaultTaskService{repo: repo}
}

func (service DefaultTaskService) GetAllTasksOfUser(userIdFromToken string) ([]dto.Task, error) {
	res, err := service.repo.FindAllTaskByUserId(userIdFromToken)
	if err != nil {
		return nil, err
	}
	tasks := make([]dto.Task, 0, len(res))
	for _, v := range res {
		tasks = append(tasks, v.ToDto())
	}

	return tasks, err
}

func (service DefaultTaskService) CreateNewTask(userIdFromToken string, task *dto.Task) (string, error) {
	tt := domains.Task{
		Id:          uuid.New().String(),
		UserId:      userIdFromToken,
		Title:       task.Title,
		Description: task.Description,
		DueDate:     task.DueDate,
		Priority:    task.Priority,
		Category:    task.Category,
		Status:      task.Status,
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
	}

	newTaskId, err := service.repo.CreateNewTask(&tt)
	if err != nil {
		return "", err
	}

	return newTaskId, nil
}

func (service DefaultTaskService) UpdateTask(task *dto.Task) error {
	existingTask, err := service.repo.FindTaskById(task.TaskId)
	if err != nil {
		return err
	}

	// update values
	if task.Title != "" {
		existingTask.Title = task.Title
	}
	if task.Description != "" {
		existingTask.Description = task.Description
	}
	if task.DueDate.String() != "" {
		existingTask.DueDate = task.DueDate
	}
	if task.Priority != "" {
		existingTask.Priority = task.Priority
	}
	if task.Category != "" {
		existingTask.Category = task.Category
	}
	if task.Status != "" {
		existingTask.Status = task.Status
	}

	existingTask.UpdatedAt = time.Now()

	err = service.repo.UpdateTask(existingTask)
	if err != nil {
		return err
	}

	return nil
}

func (service DefaultTaskService) DeleteTask(id string) error {
	existingTask, err := service.repo.FindTaskById(id)
	if err != nil {
		log.Println("returned from 1st if")
		return err
	}

	if existingTask == nil {
		log.Println("returned from 2nd if")
		return errors.New("task does not exists")
	}

	err = service.repo.DeleteTask(id)
	if err != nil {
		return err
	}

	return nil
}

func (service DefaultTaskService) ViewTaskWithFilter(userIdFromToken string, dueDate string, priority string, category string, status string) ([]dto.Task, error) {
	query := `SELECT * FROM tasks WHERE user_id = $1`
	values := []any{userIdFromToken}

	// Add filters dynamically and update placeholders correctly
	if dueDate != "" {
		query += " AND DATE(due_date) = $" + fmt.Sprint(len(values)+1)
		values = append(values, dueDate)
	}
	if priority != "" {
		query += " AND priority = $" + fmt.Sprint(len(values)+1)
		values = append(values, priority)
	}
	if category != "" {
		query += " AND category = $" + fmt.Sprint(len(values)+1)
		values = append(values, category)
	}
	if status != "" {
		query += " AND status = $" + fmt.Sprint(len(values)+1)
		values = append(values, status)
	}

	res, err := service.repo.FindTaskByFilter(query, values...)
	if err != nil {
		return nil, err
	}

	tasks := make([]dto.Task, 0, len(res))
	for _, v := range res {
		tasks = append(tasks, v.ToDto())
	}

	return tasks, err
}
