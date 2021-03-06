package main

import (
	"fmt"
	"time"
)

func main() {
	drop()
}

func drop() {
	const cap = 5
	ch := make(chan string, cap)

	go func() {
		for p := range ch {
			fmt.Printf("employee: received signal %s \n", p)
		}
	}()

	const work = 20
	for w := 0; w < work; w++ {
		select {
		case ch <- "paper":
			fmt.Println("manager sent signal: ", w)
		default:
			fmt.Println("manager dropped data: ", w)
		}
	}

	close(ch)
	fmt.Println("manager sent shutdown signal")

	time.Sleep(time.Second)
	fmt.Println("------------------------------------------------------------------")
}