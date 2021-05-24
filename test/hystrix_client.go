package main

import (
	"fmt"
	"net/http"
	"time"
)

func main() {
	for i := 0; i < 1000; i++ {
		go func() {
			fmt.Println(http.Get("http://localhost:13333"))
		}()
	}

	time.Sleep(time.Second * 5)
}
