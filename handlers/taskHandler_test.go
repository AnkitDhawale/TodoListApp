package handlers

import (
	"errors"
	"github.com/AnkitDhawale/TodoListApp/dto"
	"github.com/AnkitDhawale/TodoListApp/middlewares"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"golang.org/x/net/context"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"
)

var testUserId = "0451ff67-df6f-463b-9e35-b6dabe0e8a55"

type MockTaskService struct {
	mock.Mock
}

func (m MockTaskService) GetAllTasksOfUser(userIdFromToken string) ([]dto.Task, error) {
	res := m.Called(userIdFromToken)

	return res.Get(0).([]dto.Task), res.Error(1)
}

func (m MockTaskService) CreateNewTask(userIdFromToken string, task *dto.Task) (string, error) {
	//TODO implement me
	panic("implement me")
}

func (m MockTaskService) UpdateTask(task *dto.Task) error {
	//TODO implement me
	panic("implement me")
}

func (m MockTaskService) DeleteTask(id string) error {
	//TODO implement me
	panic("implement me")
}

func (m MockTaskService) ViewTaskWithFilter(userIdFromToken string, dueDate string, priority string, category string, status string) ([]dto.Task, error) {
	//TODO implement me
	panic("implement me")
}

func TestTaskHandler_GetAllTasksOfAUser(t *testing.T) {
	tests := []struct {
		name                 string
		contextUserID        string
		mockServiceOutput    []dto.Task
		mock                 func(ms *MockTaskService, contextUserId string)
		expectedResponseBody string
		expectedStatusCode   int
	}{
		{
			name:          "Success #1",
			contextUserID: testUserId,
			mockServiceOutput: []dto.Task{
				{
					TaskId:      "332ae328-a2da-4560-8934-c4f63f376933",
					Title:       "Fitness",
					Description: "Stay healthy and fit with gym",
					DueDate:     time.Time{},
					Priority:    "High",
					Category:    "GYM",
					Status:      "Pending",
					CreatedAt:   time.Time{},
					UpdatedAt:   time.Time{},
				},
				{
					TaskId:      "333ae328-a2da-4560-8934-c4f63f376934",
					Title:       "Fitness2",
					Description: "Stay healthy and fit with gym",
					DueDate:     time.Time{},
					Priority:    "Medium",
					Category:    "GYM",
					Status:      "Complete",
					CreatedAt:   time.Time{},
					UpdatedAt:   time.Time{},
				},
			},
			mock: func(ms *MockTaskService, contextUserId string) {
				tasks := []dto.Task{
					{
						TaskId:      "332ae328-a2da-4560-8934-c4f63f376933",
						Title:       "Fitness",
						Description: "Stay healthy and fit with gym",
						DueDate:     time.Time{},
						Priority:    "High",
						Category:    "GYM",
						Status:      "Pending",
						CreatedAt:   time.Time{},
						UpdatedAt:   time.Time{},
					},
					{
						TaskId:      "333ae328-a2da-4560-8934-c4f63f376934",
						Title:       "Fitness2",
						Description: "Stay healthy and fit with gym",
						DueDate:     time.Time{},
						Priority:    "Medium",
						Category:    "GYM",
						Status:      "Complete",
						CreatedAt:   time.Time{},
						UpdatedAt:   time.Time{},
					},
				}
				ms.On("GetAllTasksOfUser", contextUserId).Return(tasks, nil)
			},
			expectedStatusCode: http.StatusOK,
			expectedResponseBody: `{
				"data": [
					{
						"task_id": "332ae328-a2da-4560-8934-c4f63f376933",
						"title": "Fitness",
						"description": "Stay healthy and fit with gym",
						"due_date": "0001-01-01T00:00:00Z",
						"priority": "High",
						"category": "GYM",
						"status": "Pending",
						"created_at": "0001-01-01T00:00:00Z",
						"updated_at": "0001-01-01T00:00:00Z"
					},
					{
						"task_id": "333ae328-a2da-4560-8934-c4f63f376934",
						"title": "Fitness2",
						"description": "Stay healthy and fit with gym",
						"due_date": "0001-01-01T00:00:00Z",
						"priority": "Medium",
						"category": "GYM",
						"status": "Complete",
						"created_at": "0001-01-01T00:00:00Z",
						"updated_at": "0001-01-01T00:00:00Z"
					}
				]
			}`,
		},
		{
			name:              "Success #2: No task data available",
			contextUserID:     testUserId,
			mockServiceOutput: nil,
			mock: func(ms *MockTaskService, contextUserId string) {
				var tasks []dto.Task
				ms.On("GetAllTasksOfUser", contextUserId).Return(tasks, nil)
			},
			expectedResponseBody: `{"data":[]}`,
			expectedStatusCode:   http.StatusNoContent,
		},
		{
			name:               "Failure: No use id in context",
			mock:               nil,
			expectedStatusCode: http.StatusInternalServerError,
			expectedResponseBody: `{
				"data":null,
				"error_message":"userId not found in context"
			}`,
		},
		{
			name:              "Failure: err form task service",
			contextUserID:     testUserId,
			mockServiceOutput: nil,
			mock: func(ms *MockTaskService, contextUserId string) {
				var tasks []dto.Task
				ms.On("GetAllTasksOfUser", contextUserId).Return(tasks, errors.New("something went wrong"))
			},
			expectedResponseBody: `{
				"data":null,
				"error_message":"something went wrong"
			}`,
			expectedStatusCode: http.StatusInternalServerError,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			mockTaskService := new(MockTaskService)
			taskHandler := TaskHandler{Service: mockTaskService}

			// Create a mock HTTP request and response
			req := httptest.NewRequest(http.MethodGet, "/tasks", nil)
			w := httptest.NewRecorder()

			// Add the user ID to the request context
			ctx := req.Context()
			if tt.contextUserID != "" {
				ctx = context.WithValue(ctx, middlewares.UserIDKey, tt.contextUserID)
			}

			req = req.WithContext(ctx)

			if tt.mock != nil {
				// Set up mock behavior
				tt.mock(mockTaskService, tt.contextUserID)
			}

			// Call the handler
			taskHandler.GetAllTasksOfAUser(w, req)

			// Assert & verify the response
			assert.Equal(t, tt.expectedStatusCode, w.Code)

			assert.JSONEq(t, tt.expectedResponseBody, strings.TrimSpace(w.Body.String()))
			mockTaskService.AssertExpectations(t)
		})
	}
}
