package main

import (
	"fmt"
	"time"
)

type V struct {
	val         string
	expiresTime time.Time
}

func main() {

	ticker := time.NewTicker(time.Second * 5)
	timer := time.NewTimer(time.Second)

	for true {
		select {
		case <-timer.C:
			fmt.Println("timer c")
		case <-ticker.C:
			fmt.Println("ticker c")
		}
	}

}
