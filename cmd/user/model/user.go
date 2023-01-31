package model

type User struct {
	Id   int64  `gorm:"primarykey"`
	Name string `gorm:"size:10"`
}
