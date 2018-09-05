package main

import (
	"fmt"
	"sort"
)

// SortByte .
type SortByte []int

func (a SortByte) Len() int             { return len(a) }
func (a SortByte) Swap(i, j int)        { a[i], a[j] = a[j], a[i] }
func (a SortByte) Less(i, j int) bool { // 从大到小排序
	cardValue1 := a[i] & 0xF
	cardValue2 := a[j] & 0xF
	if cardValue1 == cardValue2 {
		return a[i] == a[j]
	}
	if cardValue1 == 1 {
		cardValue1 += 13
	}
	if cardValue2 == 1 {
		cardValue2 += 13
	}
	return cardValue1 > cardValue2

}

func main() {
	a := []int{
		0x5, 0x1, 0x3, 0x23, 0x13, 0x4, 0x31,
	}
	sort.Sort(SortByte(a))
	fmt.Println(a)
}
