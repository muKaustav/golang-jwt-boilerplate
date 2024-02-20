package models

import "time"

type Author struct {
	ID        uint `gorm:"primarykey"`
	Name      string
	Email     string
	Password  string
	Verified  int
	CreatedAt time.Time
	UpdatedAt time.Time
}
