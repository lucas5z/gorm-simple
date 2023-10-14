package models

import "time"

type User struct {
	Id     int    `gorm:"primaryKey"`
	Nombre string `gorm:"type:varchar(30)"`
	Fecha  time.Time
	Tasks  []Task
}
