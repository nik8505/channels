package main

import (
	"fmt"
	"time"
)

func main() {
	pooling()
}

func pooling() {
	ch := make(chan string)

	const emps = 2
	for e := 0; e < emps; e++ {
		go func(emp int) {
			for p := range ch {
				fmt.Printf("employee %d: received signal %s \n", emp, p)
			}
			fmt.Printf("employee %d: received shutdown signal \n", emp)
		}(e)
	}

	const work = 10
	for w := 0; w < work; w++ {
		ch <- "paper"
		fmt.Println("manager sent signal: ", w)
	}

	close(ch)
	fmt.Println("manager sent shutdown signal")

	time.Sleep(time.Second)
	fmt.Println("------------------------------------------------------------------")
}