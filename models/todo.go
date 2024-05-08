package models

type Todo struct {
	ID          string
	UserID      string
	Title       string
	Description string
	Done        bool
}
