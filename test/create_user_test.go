package test

import (
	"bytes"
	"encoding/json"
	"errors"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/enghasib/laundry_service/config"
	"github.com/enghasib/laundry_service/domain"
	middleware "github.com/enghasib/laundry_service/rest/middlewares"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/mock"
)

type UserResponse struct {
	Message  string `json:"message"`
	ID       int    `json:"id"`
	UserName string `json:"username"`
	Email    string `json:"email"`
	Role     string `json:"role"`
}

type createUserRequest struct {
	UserName string `json:"username"`
	Email    string `json:"email"`
	Password string `json:"password"`
	Role     string `json:"role"`
}

type Service interface {
	Create(user domain.User) (domain.User, error)
}

type UserHandler struct {
	middleware middleware.Middlewares
	cnf        *config.Config
	srv        Service
}

func (u *UserHandler) CreateUserHandler(rec *httptest.ResponseRecorder, req *http.Request) {
	panic("unimplemented")
}

// MockUserService implements the UserService interface for testing
type MockUserService struct {
	mock.Mock
}

func (m *MockUserService) Create(user domain.User) (domain.User, error) {
	args := m.Called(user)
	return args.Get(0).(domain.User), args.Error(1)
}

func TestCreateUserHandler(t *testing.T) {
	tests := []struct {
		name           string
		requestBody    interface{}
		mockSetup      func(*MockUserService)
		expectedStatus int
		expectedBody   interface{}
	}{
		{
			name: "successful_user_creation",
			requestBody: createUserRequest{
				UserName: "testuser",
				Email:    "test@example.com",
				Password: "password123",
				Role:     "user",
			},
			mockSetup: func(m *MockUserService) {
				m.On("Create", domain.User{
					UserName: "testuser",
					Email:    "test@example.com",
					Password: "password123",
					Role:     "user",
				}).Return(domain.User{
					Id:       1,
					UserName: "testuser",
					Email:    "test@example.com",
					Role:     "user",
				}, nil)
			},
			expectedStatus: http.StatusCreated,
			expectedBody: UserResponse{
				Message:  "User created Successfully!",
				ID:       1,
				UserName: "testuser",
				Email:    "test@example.com",
				Role:     "user",
			},
		},
		{
			name:           "invalid_json_body",
			requestBody:    "invalid json",
			mockSetup:      func(m *MockUserService) {},
			expectedStatus: http.StatusBadRequest,
		},
		{
			name: "service_error",
			requestBody: createUserRequest{
				UserName: "testuser",
				Email:    "test@example.com",
				Password: "password123",
				Role:     "user",
			},
			mockSetup: func(m *MockUserService) {
				m.On("Create", mock.Anything).Return(domain.User{}, errors.New("service error"))
			},
			expectedStatus: http.StatusBadRequest,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			// Create mock service and set up expectations
			mockService := new(MockUserService)
			tt.mockSetup(mockService)

			// Create handler with mock service
			handler := &UserHandler{srv: mockService}

			// Prepare request body
			var body bytes.Buffer
			if str, ok := tt.requestBody.(string); ok {
				body.WriteString(str)
			} else {
				json.NewEncoder(&body).Encode(tt.requestBody)
			}

			// Create request and response recorder
			req := httptest.NewRequest(http.MethodPost, "/users", &body)
			rec := httptest.NewRecorder()

			// Execute handler
			handler.CreateUserHandler(rec, req)

			// Assert status code
			assert.Equal(t, tt.expectedStatus, rec.Code)

			// Assert response body for successful cases
			if tt.expectedBody != nil {
				var response UserResponse
				err := json.Unmarshal(rec.Body.Bytes(), &response)
				assert.NoError(t, err)
				assert.Equal(t, tt.expectedBody, response)
			}

			// Verify that all expected mock calls were made
			mockService.AssertExpectations(t)
		})
	}
}
