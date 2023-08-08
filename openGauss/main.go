package main

import (
	"context"
	"database/sql"
	"flag"
	"fmt"
	"time"

	_ "gitee.com/opengauss/openGauss-connector-go-pq"
)

var (
	h = flag.String("h", "127.0.0.1", "openGauss host")
)

func main() {
	flag.Parse()

	connStr := "host=%s port=5432 user=gaussdb password=Password*123 dbname=gaussdb sslmode=disable"
	db, err := sql.Open("opengauss", fmt.Sprintf(connStr, *h))
	if err != nil {
		panic(err)
	}

	createDBQuery := "CREATE DATABASE IF NOT EXISTS gaussdb"
	_, err = db.Exec(createDBQuery)
	if err != nil {
		panic(err)
	}
	fmt.Println("数据库已创建")

	errCount := 0
	for {
		ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
		var count int64
		err = db.QueryRowContext(ctx, "SELECT count(*) FROM test.xanole_events").Scan(&count)
		cancel()
		if err != nil {
			errCount++
			fmt.Printf("查询遇到错误: %v, 错误次数=%d\n", err, errCount)
			if errCount > 3 {
				fmt.Printf("查询总错误次数大于3次，程序退出")
				return
			}
			time.Sleep(1500 * time.Millisecond)
			continue
		}
		fmt.Printf("查询到表的行数=%d\n", count)
		time.Sleep(1500 * time.Millisecond)
	}
}
