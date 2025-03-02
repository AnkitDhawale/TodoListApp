package services

import (
	"errors"
	"github.com/AnkitDhawale/TodoListApp/domains"
	"github.com/AnkitDhawale/TodoListApp/dto"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
	"testing"
	"time"
)

type MockTaskRepo struct {
	mock.Mock
}

func (m *MockTaskRepo) CreateNewTask(task *domains.Task) (string, error) {
	args := m.Called(task)

	return args.String(0), args.Error(1)
}

func (m *MockTaskRepo) FindAllTaskByUserId(userIdFromToken string) ([]domains.Task, error) {
	args := m.Called(userIdFromToken)

	return args.Get(0).([]domains.Task), args.Error(1)
}

func (m *MockTaskRepo) UpdateTask(task *domains.Task) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockTaskRepo) DeleteTask(id string) error {
	//TODO implement me
	panic("implement me")
}

func (m *MockTaskRepo) FindTaskById(taskId string) (*domains.Task, error) {
	//TODO implement me
	panic("implement me")
}

func (m *MockTaskRepo) FindTaskByFilter(query string, values ...any) ([]domains.Task, error) {
	//TODO implement me
	panic("implement me")
}

/*
	func TestDefaultTaskService_CreateNewTask(t *testing.T) {
		tests := []struct {
			name                 string
			userIDFromTokenInput string
			taskInput            *dto.Task
			expectedResult       string
			expectedError        error
		}{
			{
				name:                 "Success",
				userIDFromTokenInput: "0451ff67-df6f-463b-9e35-b6dabe0e8a55",
				taskInput: &dto.Task{
					Title:       "Fitness",
					Description: "Stay healthy and fit with gym",
					DueDate:     time.Time{},
					Priority:    "High",
					Category:    "GYM",
					Status:      "Pending",
					CreatedAt:   time.Time{},
					UpdatedAt:   time.Time{},
				},
				expectedResult: "1211ff67-df6f-463b-9e35-b6dabe0e8a99",
				expectedError:  nil,
			},
		}

		for _, tt := range tests {
			t.Run(tt.name, func(t *testing.T) {
				mockTaskRepo := new(MockTaskRepo)
				taskService := NewDefaultTaskService(mockTaskRepo)

				newTask := domains.Task{
					Id:          "1211ff67-df6f-463b-9e35-b6dabe0e8a99",
					UserId:      tt.userIDFromTokenInput,
					Title:       tt.taskInput.Title,
					Description: tt.taskInput.Description,
					DueDate:     tt.taskInput.DueDate,
					Priority:    tt.taskInput.Priority,
					Category:    tt.taskInput.Category,
					Status:      tt.taskInput.Status,
					CreatedAt:   tt.taskInput.CreatedAt,
					UpdatedAt:   tt.taskInput.UpdatedAt,
				}

				mockTaskRepo.On("CreateNewTask", &newTask).Return(tt.expectedResult, nil)
				res, err := taskService.CreateNewTask(tt.userIDFromTokenInput, tt.taskInput)

				assert.Equal(t, tt.expectedResult, res)
				assert.NoError(t, err)
				mockTaskRepo.AssertExpectations(t)
			})
		}
	}
*/
func TestDefaultTaskService_GetAllTasksOfUser(t *testing.T) {
	tests := []struct {
		name                 string
		userIDFromTokenInput string
		expectedTasksOutput  []dto.Task
		mock                 func(*MockTaskRepo)
		expectedError        error
	}{
		{
			name:                 "Success",
			userIDFromTokenInput: "0451ff67-df6f-463b-9e35-b6dabe0e8a55",
			expectedTasksOutput: []dto.Task{
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
			mock: func(mm *MockTaskRepo) {
				tasks := []domains.Task{
					{
						Id:          "332ae328-a2da-4560-8934-c4f63f376933",
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
						Id:          "333ae328-a2da-4560-8934-c4f63f376934",
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
				mm.On("FindAllTaskByUserId", "0451ff67-df6f-463b-9e35-b6dabe0e8a55").Return(tasks, nil)
			},
			expectedError: nil,
		},
		{
			name:                 "failure",
			userIDFromTokenInput: "0451ff67-df6f-463b-9e35-b6dabe0e8a55",
			expectedTasksOutput:  nil,
			mock: func(mm *MockTaskRepo) {
				mm.On("FindAllTaskByUserId", "0451ff67-df6f-463b-9e35-b6dabe0e8a55").Return([]domains.Task{}, errors.New("error while getting rows"))
			},
			//expectedError: fmt.Errorf("error while getting rows"),
			expectedError: errors.New("error while getting rows"),
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			var mockTaskRepo = new(MockTaskRepo)
			taskService := NewDefaultTaskService(mockTaskRepo)

			tt.mock(mockTaskRepo)
			res, err := taskService.GetAllTasksOfUser(tt.userIDFromTokenInput)

			// Assertions
			if tt.expectedError != nil {
				assert.NotNil(t, err)
				assert.EqualError(t, err, tt.expectedError.Error())
			} else {
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedTasksOutput, res)
			}

			mockTaskRepo.AssertExpectations(t)
		})
	}
}
