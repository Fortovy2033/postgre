package sql

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

func CreateTable(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	CREATE TABLE IF NOT EXISTS tasks (
	    id SERIAL PRIMARY KEY,
			title VARCHAR(200) NOT NULL,
			description VARCHAR(1000) NOT NULL,
			completed BOOLEAN NOT NULL,
			created_at TIMESTAMP NOT NULL,
			completed_at TIMESTAMP
	);
	`

	_, err := conn.Exec(ctx, sqlQuery)

	return err
}

func InsertRow(ctx context.Context, conn *pgx.Conn, task InsertTask) error {
	sqlQuery := `
	INSERT INTO tasks (title, description, completed, created_at)
	VALUES ($1,$2,$3,$4);
	`

	_, err := conn.Exec(ctx, sqlQuery,
		task.Title, task.Description, false, time.Now())

	return err
}

func GetRow(ctx context.Context, conn *pgx.Conn, id int) (TaskModel, error) {
	sqlQuery := `
	SELECT id, title, description, completed, created_at, completed_at
	FROM tasks
	WHERE id = $1;
	`

	row := conn.QueryRow(ctx, sqlQuery, id)

	var task TaskModel

	err := row.Scan(
		&task.ID,
		&task.Title,
		&task.Description,
		&task.Completed,
		&task.CreatedAt,
		&task.CompletedAt,
	)

	return task, err
}

func SelectRows(ctx context.Context, conn *pgx.Conn) error {
	sqlQuery := `
	SELECT id, title, description, completed, created_at, completed_at
	FROM tasks
	ORDER BY id ASC;
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

		PrintTask(task)
	}

	return nil
}

func CompleteTask(ctx context.Context, conn *pgx.Conn, id int) error {
	sqlQuery := `
	UPDATE tasks
	SET completed = TRUE, completed_at = $1
	WHERE id = $2;
	`

	_, err := conn.Exec(ctx, sqlQuery, time.Now(), id)

	return err
}

func UncompleteTask(ctx context.Context, conn *pgx.Conn, id int) error {
	sqlQuery := `
	UPDATE tasks
	SET completed = FALSE, completed_at = $1
	WHERE id = $2;
	`

	_, err := conn.Exec(ctx, sqlQuery, nil, id)

	return err
}

func UpdateTask(ctx context.Context, conn *pgx.Conn, t TaskModel) error {
	sqlQuery := `
	UPDATE tasks
	SET title=$1, description=$2, completed=$3, created_at=$4, comlpeted_at=$5
	WHERE id=$6;
	`

	_, err := conn.Exec(
		ctx,
		sqlQuery,
		t.Title,
		t.Description,
		t.Completed,
		t.CreatedAt,
		t.CompletedAt,
		t.ID,
	)

	return err
}

func DeleteRow(ctx context.Context, conn *pgx.Conn, ids []int) error {
	sqlQuery := `
	DELETE FROM tasks
	WHERE id = ANY($1);
	`

	_, err := conn.Exec(ctx, sqlQuery, ids)
	return err
}
