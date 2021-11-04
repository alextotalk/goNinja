package main

import (
	"context"
	"time"
)

func main() {
	ctx := context.Background()
	ctx, _ = context.WithTimeout(ctx, time.Second)
}
