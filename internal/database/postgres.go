package database

import (
	"context"
	"fmt"
	"os"
	"time"

	"github.com/jackc/pgx/v5"
)

func CarregaBanco() {
	dsn := os.Getenv("DATABASE_URL")
	if dsn == "" {
		panic("URL do banco de dados não definido!")
	}

	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	conn, err := pgx.Connect(ctx, dsn)
	if err != nil {
		panic(err)
	}
	defer func() {
		context.WithTimeout(context.Background(), 5*time.Second)
		defer cancel()
		_ = conn.Close(context.Background())
	}()

	var version string
	if err := conn.QueryRow(ctx, "select version()").Scan(&version); err != nil {
		panic(err)
	}
	fmt.Println(version)
}
