package main

import (
	"fmt"
	"time"
)

const count = 1000

func main(){
	var(
		counter int64
		ch = make(chan struct{}, count)
	)
	for i := 0 ; i < count; i++ {
		go func() {
			counter++
			ch <- struct{}{}
		}()
	}
		time.Sleep(time.Second)
		close(ch)

		n := 0
		for range ch {
			n++
		}
		fmt.Println(counter)
		fmt.Println(n)
	}
