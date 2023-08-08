package main

import (
	"context"
	"fmt"
)

func main() {
	ctx := context.WithValue(
		context.Background(),
		"testKey", "testValue",
	)

	print(ctx)
}

func print(ctx context.Context) {
	fmt.Println(ctx.Value("testKey"))

	ctx2 := context.WithValue(
		ctx,
		"testKey", "testValue2",
	)
	fmt.Println(ctx2.Value("testKey"))
}
