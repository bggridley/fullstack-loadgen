package repository

import (
	"errors"
	"backend/data/request"
	"backend/helper"
	"backend/model"

	"gorm.io/gorm"
)

type TestRepositoryImpl struct {
	Db *gorm.DB
}

func NewTestRepositoryImpl(Db *gorm.DB) TestRepository {
	return &TestRepositoryImpl{Db: Db}
}

// Delete implements TestRepository
func (t *TestRepositoryImpl) Delete(testId int) {
	var test model.Test
	result := t.Db.Where("id = ?", testId).Delete(&test)
	helper.ErrorPanic(result.Error)
}

// FindAll implements TestRepository
func (t *TestRepositoryImpl) FindAll() []model.Test {
	var test []model.Test
	result := t.Db.Find(&test)
	helper.ErrorPanic(result.Error)
	return test
}

// FindById implements TestRepository
func (t *TestRepositoryImpl) FindById(testId int) (test model.Test, err error) {
	var tag model.Test
	result := t.Db.Find(&tag, testId)
	if result != nil {
		return tag, nil
	} else {
		return tag, errors.New("tag is not found")
	}
}

// Save implements TestRepository
func (t *TestRepositoryImpl) Save(test model.Test) {
	result := t.Db.Create(&test)
	helper.ErrorPanic(result.Error)
}

// Update implements TestRepository
func (t *TestRepositoryImpl) Update(test model.Test) {
	var updateTest = request.UpdateTestRequest{
		Id:   test.Id,
		Name: test.Name,
	}
	result := t.Db.Model(&test).Updates(updateTest)
	helper.ErrorPanic(result.Error)
}