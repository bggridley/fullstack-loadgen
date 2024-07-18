package service

import (
	"backend/data/request"
	"backend/data/response"
	"backend/helper"
	"backend/model"
	"backend/repository"

	"github.com/go-playground/validator/v10"
)

type TestServiceImpl struct {
	TestRepository repository.TestRepository
	Validate       *validator.Validate
}

func NewTestServiceImpl(testRepository repository.TestRepository, validate *validator.Validate) TestService {
	return &TestServiceImpl{
		TestRepository: testRepository,
		Validate:       validate,
	}
}

// Create implements TestService
func (t *TestServiceImpl) Create(test request.CreateTestRequest) {
	err := t.Validate.Struct(test)
	helper.ErrorPanic(err)
	testModel := model.Test{
		Name: test.Name,
	}
	t.TestRepository.Save(testModel)
}

// Delete implements TestService
func (t *TestServiceImpl) Delete(testId int) {
	t.TestRepository.Delete(testId)
}

// FindAll implements TestService
func (t *TestServiceImpl) FindAll() []response.TestResponse {
	result := t.TestRepository.FindAll()

	var tests []response.TestResponse
	for _, value := range result {
		test := response.TestResponse{
			Id:   value.Id,
			Name: value.Name,
		}
		tests = append(tests, test)
	}

	return tests
}

// FindById implements TestService
func (t *TestServiceImpl) FindById(testId int) response.TestResponse {
	testData, err := t.TestRepository.FindById(testId)
	helper.ErrorPanic(err)

	testResponse := response.TestResponse{
		Id:   testData.Id,
		Name: testData.Name,
	}
	return testResponse
}

// Update implements TestService
func (t *TestServiceImpl) Update(test request.UpdateTestRequest) {
	testData, err := t.TestRepository.FindById(test.Id)
	helper.ErrorPanic(err)
	testData.Name = test.Name
	t.TestRepository.Update(testData)
}