package mock

import (
	"errors"
	"my-todo-app/models"
)

type MockUserRepository struct {
	users []models.User
}

func NewMockUserRepository() *MockUserRepository {
	return &MockUserRepository{users: []models.User{}}
}

//	func (m *MockUserRepository) CreateUser(user models.User) (models.User, error) {
//		m.users = append(m.users, user)
//		return user, nil
//	}
func (m *MockUserRepository) CreateUser(user models.User) error {
	for _, u := range m.users {
		if u.Username == user.Username {
			return errors.New("user already exists")
		}
	}
	m.users = append(m.users, user)
	return nil
}

func (m *MockUserRepository) GetUserByUsername(username string) (models.User, error) {
	for _, user := range m.users {
		if user.Username == username {
			return user, nil
		}
	}
	return models.User{}, errors.New("user not found")
}

func (m *MockUserRepository) GetAllUsers() []models.User {
	return m.users
}
