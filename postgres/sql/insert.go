package sql

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

func InsertRow(ctx context.Context, conn *pgx.Conn, title, description string) error {
	sqlQuery := `
	INSERT INTO tasks (title, description, completed, created_at)
	VALUES ($1,$2,$3,$4);
	`

	_, err := conn.Exec(ctx, sqlQuery,
		title, description, false, time.Now())

	return err
}
