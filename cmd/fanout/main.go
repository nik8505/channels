package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	// fanout()
	fanoutSem()
}

func fanout() {
	emps := 20
	ch := make(chan string, emps)

	for e := 0; e < emps; e++ {
		go func(emp int){
			time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
			ch <- "paper"
			fmt.Println("employee: sent signal: ", emp)
		}(e)
	}

	for emps > 0 {
		p := <-ch
		fmt.Println(p)
		fmt.Println("manager: received signal: ", emps)
		emps--
	}

	time.Sleep(time.Second)
	fmt.Println("------------------------------------------------------------------")
}

func fanoutSem() {
	emps := 20
	ch := make(chan string, emps)

	const cap = 5
	sem := make(chan bool, cap)

	for e := 0; e < emps; e++ {
		go func(emp int){
			sem <- true
			{
				time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
				ch <- "paper"
				fmt.Println("employee: sent signal: ", emp)
			}
			<-sem
		}(e)
	}

	for emps > 0 {
		p := <-ch
		fmt.Println(p)
		fmt.Println("manager: received signal: ", emps)
		emps--
	}

	time.Sleep(time.Second)
	fmt.Println("------------------------------------------------------------------")
}
