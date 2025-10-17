package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		timestamp := time.Now().Format(time.RFC3339)
		fmt.Printf("Hello, World! %s\n", timestamp)
		time.Sleep(5 * time.Second)
	}
}
