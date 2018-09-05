package main

import (
	"fmt"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Printf("%d", os.Getpid())
	sig := make(chan os.Signal, 1)
	//done := make(chan bool, 1)

	signal.Notify(sig, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
	sigs := <-sig
	log.Printf("Closing down (signal: %v)", sigs)
	fmt.Printf("%d", os.Getpid())
}
