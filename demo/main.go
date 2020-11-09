package main

import (
	"fmt"
	"time"
)

func main() {
	after := time.After(time.Millisecond * 150)
	for ; ; {
		select {
		case <-after:
			fmt.Println("after")
		}
	}
}
