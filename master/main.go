package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Hello World!!")
	n := 0
	for {
		fmt.Println(n)
		time.Sleep(1 * time.Second)
		n++
	}
}
