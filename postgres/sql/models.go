package sql

import (
	"fmt"
	"time"
)

type TaskModel struct {
	ID          int
	Title       string
	Description string
	Completed   bool
	CreatedAt   time.Time
	CompletedAt *time.Time
}

type InsertTask struct {
	Title       string
	Description string
}

func NewTask(title, description string) InsertTask {
	return InsertTask{
		Title:       title,
		Description: description,
	}
}

func PrintTask(t TaskModel) {
	fmt.Println("---------------------------")
	fmt.Println("id:", t.ID)
	fmt.Println("title:", t.Title)
	fmt.Println("description:", t.Description)
	fmt.Println("completed:", t.Completed)
	fmt.Println("created at:", t.CreatedAt)
	fmt.Println("completed at:", t.CompletedAt)
}
