package model

type Test struct {
	Id   int    `gorm:"type:int;primary_key"`
	Name string `gorm:"type:varchar(255)"`
}

func (Test) TableName() string {
	return "test"
}