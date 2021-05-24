package main

import (
	"fmt"
	"sync"
)

func main() {
	number := make(chan int)
	char := make(chan string)
	waitGroup := &sync.WaitGroup{}
	waitGroup.Add(1)
	go nums(number, char, waitGroup)
	go chars(char, number, waitGroup)

	number <- 1

	waitGroup.Wait()
}

func chars(charChan <-chan string, numChan chan<- int, wait *sync.WaitGroup) {

	i := rune('A')
	for true {
		select {
		case <-charChan:
			if i >= 'Z' {
				wait.Done()
				return
			}
			fmt.Println(string(i))
			i = i + 1

			fmt.Println(string(i))
			i = i + 1
			numChan <- 1
			//default:
			//	time.Sleep(time.Millisecond * 30)

		}
	}
}

func nums(numChan <-chan int, charChan chan<- string, wait *sync.WaitGroup) {
	i := 1
	for true {
		select {
		case <-numChan:
			if i >= 26 {
				charChan <- "A"
				return
			}
			fmt.Println(i)
			i = i + 1
			fmt.Println(i)
			i = i + 1
			charChan <- "A"
			//default:
			//	time.Sleep(time.Millisecond * 30)
		}
	}
}
