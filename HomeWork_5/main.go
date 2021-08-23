package main

import (
	"fmt"
	"sync"
	"time"
)
const n = 2000

	func main() {
		var (
			counter int
			mutex sync.Mutex

			ch = make(chan struct{}, n)
		)
		for i := 0; i < n; i += 1 {
			go func () {
				mutex.Lock()
				counter += 1
				defer mutex.Unlock()

				ch <- struct{}{}
			}()
		}
		time.Sleep(2*time.Second)
		close(ch)

		i := 0
		for range ch {
			i += 1
		}

		fmt.Println(n)
		fmt.Println(i)

}
