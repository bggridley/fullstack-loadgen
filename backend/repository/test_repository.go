package repository

import (
	"backend/model"
)

type TestRepository interface {
	Save(test model.Test)
	Update(test model.Test)
	Delete(testId int)
	FindById(testId int) (test model.Test, err error)
	FindAll() []model.Test
}