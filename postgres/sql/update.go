package sql

import (
	"context"

	"github.com/jackc/pgx/v5"
)

func UpdateRow(ctx context.Context, conn *pgx.Conn, id int) error {
	sqlQuery := `
	UPDATE tasks
	SET completed = TRUE
	WHERE id = $1
	`

	_, err := conn.Exec(ctx, sqlQuery, id)

	return err
}
