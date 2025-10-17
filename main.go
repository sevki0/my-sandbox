package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Argo CD whispers: deploy complete 😎")
		time.Sleep(3 * time.Second)
	}
}
