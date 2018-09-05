package main

import (
	"fmt"
	"io/ioutil"
	"os/exec"
	"strconv"
	"time"
)

func main() {

	b, err := ioutil.ReadFile("pid.txt")

	if err != nil {
		fmt.Print(err)
	}
	fmt.Println(b)
	str := string(b)
	in, err := strconv.Atoi(str)
	oldpid := in
	ticker := time.NewTicker(10 * time.Second)
	for _ = range ticker.C {
		fmt.Println(time.Now())
		b, err = ioutil.ReadFile("pid.txt")

		if err != nil {
			fmt.Print(err)
		}
		str = string(b)
		in, err = strconv.Atoi(str)
		fmt.Println(in)
		if in != 0 && in != oldpid {
			cmd := exec.Command("cmd", "/C", "start", "start_center.bat")
			// linux: cmd := exec.Command("/bin/bash", "-c", "start_center.sh")
			//bytes, err := cmd.Output()
			err := cmd.Run()
			if err != nil {
				fmt.Println("cmd.Output: ", err)
				return
			}
			//fmt.Println(string(bytes))
			oldpid = in
			fmt.Println("oldPid", oldpid, in)
		}
	}
}
