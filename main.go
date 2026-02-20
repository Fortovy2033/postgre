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

	// if err := sql.CreateTable(ctx, conn); err != nil {
	// 	panic(err)
	// }

	// if err := sql.InsertRow(ctx, conn, sql.NewTask("завтрак", "овсяные хлопья")); err != nil {
	// 	panic(err)
	// }

	// if err := sql.UpdateTask(ctx, conn, sql.TaskModel{}); err != nil {
	// 	panic(err)
	// }

	// if err := sql.DeleteRow(ctx, conn, []int{2,4}); err != nil {
	// 	panic(err)
	// }

	// if err := sql.CompleteTask(ctx, conn, 3); err != nil {
	// 	panic(err)
	// }

	// if err := sql.UncompleteTask(ctx, conn, 3); err != nil {
	// 	panic(err)
	// }

	if err := sql.SelectRows(ctx, conn); err != nil {
		panic(err)
	}

	// if err := sql.CompleteTask(ctx, conn, 3); err != nil {
	// 	panic(err)
	// }

	// task, err := sql.GetRow(ctx, conn, 3)
	// if err != nil {
	// 	panic(err)
	// }

	// sql.PrintTask(task)
}
