package models

import "time"

type Todo struct {
	ID           string
	UserID       string
	Title        string
	Description  string
	Done         bool
	CreatedTime  time.Time
	ModifiedTime time.Time
}
