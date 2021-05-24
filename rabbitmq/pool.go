package main

import (
	"fmt"
	"net/http"
)

func main() {

	for i := 0; i < 1000000; i++ {
		_, err := http.Get("http://localhost:9999")
		fmt.Println(err)
		fmt.Println(i)
	}
}
