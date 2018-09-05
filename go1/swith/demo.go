package main

import (
	"fmt"
)

func main() {
	fmt.Println(3 % 4)
	for i := 2; i != i%3; i++ {
		fmt.Println(i)
		switch 3 {
		case 3:

			fmt.Println("3")
		case 4:
			fmt.Println("4")
		}
	}
}
