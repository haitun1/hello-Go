package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"os/signal"
	"strconv"
	"syscall"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {

	str := strconv.Itoa(os.Getpid()) + "\n"
	txt := []byte(str)
	err := ioutil.WriteFile("../osRead/pid.txt", txt, 0644)

	check(err)

	sigs := make(chan os.Signal, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)

	for sig := range sigs {
		fmt.Printf("received ctrl+c(%v)\n", sig)
		str = strconv.Itoa(0)
		txt = []byte(str)
		ioutil.WriteFile("../osRead/pid.txt", txt, 0644)
		os.Exit(0)
	}

}
