package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main()  {
	// waitForTask()
	// waitForResult()
	waitForFinished()
}

// waitForTask: Think about being a manager and hiring a new employee. In this scenario, you want your new employee
// to perform a task but they need to wait until you are ready. This is because you need to hand them a piece of
// paper before they start.
func waitForTask() {
	ch := make(chan string)

	go func(){
		p := <-ch
		fmt.Println("employee: received signal: ", p)
	}()

	time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
	ch <- "paper"
	fmt.Println("manager: sent signal")

	time.Sleep(time.Second)
	fmt.Println("------------------------------------------------------------------")
}

func waitForResult() {
	ch := make(chan string)

	go func(){
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		ch <- "paper"
		fmt.Println("employee: sent signal")
	}()

	p := <-ch
	fmt.Println("manager: received signal: ", p)

	time.Sleep(time.Second)
	fmt.Println("------------------------------------------------------------------")
}

// waitForFinished: Think about being a manager and hiring a new employee. In this scenario, you want your employee to
// perform the task immediately when they are hired, and you need to wait for the result of their work. You need to
// wait because you can't move on until you know they are done but you don't need anything from them.
func waitForFinished() {
	ch := make(chan struct{})

	go func(){
		time.Sleep(time.Duration(rand.Intn(500)) * time.Millisecond)
		close(ch)
		fmt.Println("employee: sent signal")
	}()

	_, wd := <-ch
	fmt.Println("manager: received signal: ", wd)

	time.Sleep(time.Second)
	fmt.Println("------------------------------------------------------------------")
}