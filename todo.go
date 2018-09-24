package main

import "time"

type Todo struct {
	ID        int       `json:"id"`
	Name      string    `json:"name" gorm:"not null"`
	Completed bool      `json:"completed" gorm:"not null;default:false"`
	Due       time.Time `json:"due" gorm:"not null"`
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Todos []Todo
