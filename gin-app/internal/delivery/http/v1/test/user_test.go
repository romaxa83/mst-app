package test

import (
	"bytes"
	"context"
	"errors"
	"github.com/gin-gonic/gin"
	"github.com/golang/mock/gomock"
	"github.com/magiconair/properties/assert"
	"github.com/romaxa83/mst-app/gin-app/internal/delivery/http/v1"
	"github.com/romaxa83/mst-app/gin-app/internal/services"
	mock_services "github.com/romaxa83/mst-app/gin-app/internal/services/mocks"
	"net/http/httptest"
	"testing"
)

func TestHandler_userSignUp(t *testing.T) {
	// принимает мок сервиса и структуру пользователя
	type mockBehavior func(s *mock_services.MockUsers, input services.UserSignUpInput)
	// тестовая таблица с вариантами действия
	testTable := []struct {
		name                string                   // имя теста
		inputBody           string                   // тело запроса
		inputUser           services.UserSignUpInput // структура пользователя, которую передаем в сервис
		mockBehavior        mockBehavior             // mockBehavior
		expectedStatusCode  int                      // ожидаемый статус код
		expectedRequestBody string                   // ожидаемый ответ
	}{
		{
			name: "OK",
			inputBody: `{
    			"name": "cubic",
    			"email": "rubic@rubic.com",
    			"phone": "00000000001",
    			"password": "password"
			}`,
			inputUser: services.UserSignUpInput{
				Name:     "cubic",
				Email:    "rubic@rubic.com",
				Phone:    "00000000001",
				Password: "password",
			},
			mockBehavior: func(s *mock_services.MockUsers, input services.UserSignUpInput) {
				s.EXPECT().SignUp(context.Background(), input).Return(1, nil)
			},
			expectedStatusCode:  201,
			expectedRequestBody: `{"id":1}`,
		},
		{
			name: "Service failure",
			inputBody: `{
    			"name": "cubic",
    			"email": "rubic@rubic.com",
    			"phone": "00000000001",
    			"password": "password"
			}`,
			inputUser: services.UserSignUpInput{
				Name:     "cubic",
				Email:    "rubic@rubic.com",
				Phone:    "00000000001",
				Password: "password",
			},
			mockBehavior: func(s *mock_services.MockUsers, input services.UserSignUpInput) {
				s.EXPECT().SignUp(context.Background(), input).Return(0, errors.New("service failure"))
			},
			expectedStatusCode:  500,
			expectedRequestBody: `{"message":"service failure"}`,
		},
		{
			name: "Empty name",
			inputBody: `{
    			"email": "rubic@rubic.com",
    			"phone": "00000000001",
    			"password": "password"
			}`,
			mockBehavior:        func(s *mock_services.MockUsers, input services.UserSignUpInput) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
		{
			name: "Empty email",
			inputBody: `{
    			"name": "cubic",
    			"phone": "00000000001",
    			"password": "password"
			}`,
			mockBehavior:        func(s *mock_services.MockUsers, input services.UserSignUpInput) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
		{
			name: "Empty phone",
			inputBody: `{
    			"name": "cubic",
    			"email": "rubic@rubic.com",
    			"password": "password"
			}`,
			mockBehavior:        func(s *mock_services.MockUsers, input services.UserSignUpInput) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
		{
			name: "Empty phone",
			inputBody: `{
    			"name": "cubic",
    			"email": "rubic@rubic.com",
    			"phone": "00000000001",
			}`,
			mockBehavior:        func(s *mock_services.MockUsers, input services.UserSignUpInput) {},
			expectedStatusCode:  400,
			expectedRequestBody: `{"message":"invalid input body"}`,
		},
	}

	for _, testCase := range testTable {
		t.Run(testCase.name, func(t *testing.T) {
			// Init Dependencies
			c := gomock.NewController(t)
			defer c.Finish()

			userService := mock_services.NewMockUsers(c)
			testCase.mockBehavior(userService, testCase.inputUser)

			services := &services.Services{Users: userService}
			handler := v1.Handler{Services: services}

			// Test Server
			r := gin.New()
			r.POST("/sign-up", handler.UserSignUp)

			// Test Request
			w := httptest.NewRecorder()
			req := httptest.NewRequest("POST", "/sign-up",
				bytes.NewBufferString(testCase.inputBody))

			// Perform Request
			r.ServeHTTP(w, req)

			// Assert
			assert.Equal(t, testCase.expectedStatusCode, w.Code)
			assert.Equal(t, testCase.expectedRequestBody, w.Body.String())
		})
	}
}
