package main

import (
	"fmt"
	"github.com/go-vgo/robotgo"
	"time"
)

func main() {
	err := robotgo.ActiveName("slack")
	if err != nil {
		fmt.Println(err)
	}

	for {
		robotgo.MoveSmooth(300, 300)
		time.Sleep(5 * time.Second)
		robotgo.MoveSmooth(600, 600)
		time.Sleep(5 * time.Minute)
	}
}
