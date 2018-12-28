package main

import (
	"fmt"
	"sync"
	"time"
)

func main() { // by default, main is a go routine

	var waitGrp sync.WaitGroup
	waitGrp.Add(2)
	go func() { //this is a go routine, like OS thread, but go routine is being managed by go, so it will be fast
		defer waitGrp.Done()
		time.Sleep(5 * time.Second)
		fmt.Println("Just woke up, going back to sleep")
	}() // ()at the last make this function self-executing, means it will execute automatically wothout being called
	go func() {
		defer waitGrp.Done()
		fmt.Println("Hello world")
	}()
	waitGrp.Wait()
}
