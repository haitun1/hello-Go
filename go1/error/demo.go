package main

import (
	"errors"
	"fmt"
)

// ErrCompareequality 比牌相等
var ErrCompareequality = errors.New("CompareCard equality")

func main() {
	er := err()
	if er == ErrCompareequality {
		fmt.Println(er)
	}
	var a []int
	fmt.Println(a)
}

func err() error {
	return ErrCompareequality
}
