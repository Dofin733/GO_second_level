package main

import (
	"fmt"
	"os"
	"runtime/trace"
	"sync"
)

var counter int = 0
func main() {

	trace.Start(os.Stderr)
	defer trace.Stop()

	ch := make(chan bool)
	var mutex sync.Mutex

	for i := 1; i < 5; i++{
		go work(i, ch, &mutex)
	}

	for i := 1; i < 5; i++{
		<-ch
	}

	fmt.Println("The End")
}
func work (number int, ch chan bool, mutex *sync.Mutex){
	mutex.Lock()
	counter = 0
	for k := 1; k <= 5; k++{
		counter++
		fmt.Println("Goroutine", number, "-", counter)
	}
	mutex.Unlock()
	ch <- true
}