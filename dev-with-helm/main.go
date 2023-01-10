package main

import (
	"fmt"
	"time"
)

func main() {
	for {
		fmt.Println("Hello world v8!")

		time.Sleep(time.Second * 1)
	}
}
