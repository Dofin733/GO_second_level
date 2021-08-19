package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func main() {
Chan := make(chan os.Signal, 1)
signal.Notify(Chan, syscall.SIGTERM, syscall.SIGINT)

go func() {
for i := 0; i < 5; i ++{
	fmt.Printf("%d", i)
    log.Println("tik-tak")
    time.Sleep(time.Second)
}
}()
sig := <-Chan
log.Printf("Caught SIGTERM %v", sig)
}