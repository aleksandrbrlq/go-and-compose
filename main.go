package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Hello one")
		time.Sleep(time.Second * 3)
	}
}
