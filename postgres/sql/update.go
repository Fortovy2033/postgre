package sql

import (
	"context"
	"time"

	"github.com/jackc/pgx/v5"
)

func UpdateRow(ctx context.Context, conn *pgx.Conn, id int) error {
	sqlQuery := `
	UPDATE tasks
	SET completed = TRUE
	SET completed_at = $1
	WHERE id = $2
	`

	_, err := conn.Exec(ctx, sqlQuery, time.Now(), id)

	return err
}