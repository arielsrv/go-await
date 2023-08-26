# go-await
![Coverage](https://img.shields.io/badge/Coverage-100.0%25-brightgreen)

```go
package main

import (
	"context"
	"log"
	"time"

	"github.com/arielsrv/go-await/async"
)

func main() {
	f := func() string {
		log.Println("Warming up ...")
		time.Sleep(3 * time.Second)
		log.Println("Done ...")
		return "Hello world!"
	}

	// Without context
	log.Println("Let's start ...")
	future := async.Run(func() any {
		return f()
	})
	log.Println("Running ...")
	value := future.Await()
	log.Println(value)

	// With context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	log.Println("Let's start ...")
	future = async.RunWithContext(ctx, func() any {
		return f()
	})
	log.Println("Running ...")
	value = future.AwaitWithContext(ctx)
	log.Println(value)
}
```
