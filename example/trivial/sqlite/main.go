package main

import (
	"context"
	"database/sql"
	"fmt"

	"github.com/sword0307/bun/"
	"github.com/sword0307/bun/dialect/sqlitedialect"
	"github.com/sword0307/bun/driver/sqliteshim"
	"github.com/sword0307/bun/extra/bundebug"
)

func main() {
	ctx := context.Background()

	// Open an in-memory SQLite database.
	sqlite, err := sql.Open(sqliteshim.ShimName, "file::memory:?cache=shared")
	if err != nil {
		panic(err)
	}

	// Create a Bun db on top of it.
	db := bun.NewDB(sqlite, sqlitedialect.New())

	// Print all queries to stdout.
	db.AddQueryHook(bundebug.NewQueryHook(bundebug.WithVerbose(true)))

	var rnd int64

	// Select a random number.
	if err := db.NewSelect().ColumnExpr("random()").Scan(ctx, &rnd); err != nil {
		panic(err)
	}

	fmt.Println(rnd)
}
