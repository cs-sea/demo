package main

import (
	"demo/server/discover"
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	dis := discover.NewConsulService()

	for i := 0; i < 10000; i++ {
		s, err := dis.Get("hello")
		if err != nil {
			log.Fatalln(err)
		}
		time.Sleep(time.Second)
		res, err := http.Get(fmt.Sprintf("http://%s:%d", s[0].Address, s[0].Port))
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println(res)
	}

}
