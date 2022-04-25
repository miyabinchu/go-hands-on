package main

import (
	"fmt"
)

// Mydata is structure.
type Mydata struct {
	Name string
	Data []int
}

func main() {
	taro := Mydata{
		"Taro",
		[]int{10, 20, 30},
	}
	fmt.Println(taro)
	taro = rev(taro)
	fmt.Println(taro)
}
