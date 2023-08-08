package main

import (
	"context"
	"fmt"
	"time"
)

func main() {
	date := time.Date(1, 1, 1, 1, 1, 1, 1, time.UTC)
	fmt.Println(date.Sub(time.Now()))
	context.WithValue()
}
