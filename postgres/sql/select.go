package sql

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5"
)

func SelectRows(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	SELECT id, title, description, completed, created_at, completed_at
	FROM tasks
	ORDER BY id ASC
	`

	rows, err := conn.Query(ctx, sqlQuery)
	if err != nil {
		return nil
	}
	defer rows.Close()

	for rows.Next() {
		var task TaskModel

		err := rows.Scan(
			&task.ID,
			&task.Title,
			&task.Description,
			&task.Completed,
			&task.CreatedAt,
			&task.CompletedAt,
		)
		if err != nil {
			return err
		}

		printTask(task)
	}

	return nil
}


func printTask(t TaskModel) {
	fmt.Println("---------------------------")
	fmt.Println("id:", t.ID)
	fmt.Println("title:", t.Title)
	fmt.Println("description:", t.Description)
	fmt.Println("completed:", t.Completed)
	fmt.Println("created at:", t.CreatedAt)
	fmt.Println("completed at:", t.CompletedAt)
}