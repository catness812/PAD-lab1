package rpctransport_test

import (
	"context"
	"testing"

	rpctransport "github.com/catness812/PAD-labs/user_svc/internal/controller/rpc-transport"
	"github.com/catness812/PAD-labs/user_svc/internal/models"
	"github.com/catness812/PAD-labs/user_svc/internal/pb"
)

type MockUserService struct {
	users map[string]models.User
}

func (m *MockUserService) RegisterNewUser(user models.User) error {
	m.users[user.Username] = user
	return nil
}

func (m *MockUserService) FindUser(username string) (*models.User, error) {
	if user, exists := m.users[username]; exists {
		return &user, nil
	}
	return &models.User{}, nil
}

func (m *MockUserService) DeleteUser(username string) error {
	delete(m.users, username)
	return nil
}

func TestRegisterUser(t *testing.T) {
	mockUserService := &MockUserService{
		users: make(map[string]models.User),
	}
	server := &rpctransport.Server{
		UserService: mockUserService,
	}

	username := "123"
	password := "123"

	req := &pb.RegisterUserRequest{
		User: &pb.User{
			Username: username,
			Password: password,
		},
	}

	resp, err := server.RegisterUser(context.TODO(), req)
	if err != nil {
		t.Errorf("RegisterUser returned an error: %v", err)
	}

	expectedMessage := "User '123' successfully signed up"
	if resp.Message != expectedMessage {
		t.Errorf("RegisterUser returned unexpected message: got %v want %v",
			resp.Message, expectedMessage)
	}
}
