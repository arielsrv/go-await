package main

import (
	"context"
	"log"
	"time"

	"github.com/arielsrv/go-await/async"
)

func SomeTask() int {
	log.Println("Warming up ...")
	time.Sleep(3 * time.Second)
	log.Println("Done ...")
	return 1
}

func main() {
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Let's start ...")
	future := async.RunWithContext(ctx, func() any {
		return SomeTask()
	})
	log.Println("Running ...")
	value := future.AwaitWithContext(ctx)
	log.Println(value)
}
