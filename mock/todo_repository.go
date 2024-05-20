package mock

import (
	"errors"
	"my-todo-app/models"
	"time"
)

type MockTodoRepository struct {
	todoLists []models.TodoList
}

func NewMockTodoRepository() *MockTodoRepository {
	return &MockTodoRepository{todoLists: []models.TodoList{}}
}

func (m *MockTodoRepository) CreateTodoList(todoList models.TodoList) error {
	m.todoLists = append(m.todoLists, todoList)
	return nil
}

func (m *MockTodoRepository) GetTodoListByID(id uint) (models.TodoList, error) {
	for _, todoList := range m.todoLists {
		if todoList.ID == id {
			return todoList, nil
		}
	}
	return models.TodoList{}, errors.New("todo list not found")
}

func (m *MockTodoRepository) GetTodoListByUser(userID uint, userType string) ([]models.TodoList, error) {
	var temp []models.TodoList
	for _, todoList := range m.todoLists {
		if todoList.UserID == userID || userType == "admin" {
			temp = append(temp, todoList)
		}
	}

	if len(temp) == 0 {
		return nil, errors.New("no todo lists found for the user")
	}

	return temp, nil
}
func (m *MockTodoRepository) UpdateTodoList(todoList models.TodoList) error {
	for i, t := range m.todoLists {
		if t.ID == todoList.ID {
			m.todoLists[i] = todoList
			return nil
		}
	}
	return errors.New("todo list not found")
}

func (m *MockTodoRepository) DeleteTodoList(id uint) error {
	for i, t := range m.todoLists {
		if t.ID == id {
			deletedTime := time.Now()
			m.todoLists[i].DeletedAt = &deletedTime
			return nil
		}
	}
	return errors.New("todo list not found")
}
