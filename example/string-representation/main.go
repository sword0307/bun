package main

import (
	"database/sql"
	"fmt"

	"github.com/sword0307/bun/"
	"github.com/sword0307/bun/dialect/sqlitedialect"
	"github.com/sword0307/bun/driver/sqliteshim"
)

type Item struct {
	ID int64 `bun:",pk,autoincrement"`
}

func main() {
	sqldb, err := sql.Open(sqliteshim.ShimName, "file::memory:?cache=shared")
	if err != nil {
		panic(err)
	}

	db := bun.NewDB(sqldb, sqlitedialect.New())
	defer db.Close()

	q := db.NewSelect().Model((*Item)(nil)).Where("id > ?", 0)

	fmt.Println(q.String())
}
