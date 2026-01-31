package main

import (
	"context"
	"study/postgres/connection"
	"study/postgres/sql"
)

func main() {
	ctx := context.Background()

	conn, err := connection.CreateConnection(ctx)
	if err != nil {
		panic(err)
	}

	if err := sql.CreateTable(ctx, conn); err != nil {
		panic(err)
	}

	if err := sql.InsertRow(ctx, conn, "homework", "p123, ex4"); err != nil {
		panic(err)
	}
}
