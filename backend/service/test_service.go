package service

import (
	"backend/data/request"
	"backend/data/response"
)

type TestService interface {
	Create(test request.CreateTestRequest)
	Update(test request.UpdateTestRequest)
	Delete(testId int)
	FindById(testId int) response.TestResponse
	FindAll() []response.TestResponse
}