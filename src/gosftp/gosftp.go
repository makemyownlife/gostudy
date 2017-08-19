package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("hello ,world")
	duration := time.Duration(10)*time.Second
	time.Sleep(duration)
}


