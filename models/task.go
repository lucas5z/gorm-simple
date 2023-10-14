package models

type Task struct {
	Id          int `gorm:"primaryKey"`
	UserId      int
	Descripcion string
	Completado  bool `gorm:"default:false"`
}
