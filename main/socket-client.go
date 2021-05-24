package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	conn, err := net.Dial("tcp", ":8881")
	if err != nil {
		panic("sdfsdf")
	}
	defer conn.Close()

	i := 0
	for true {
		if conn == nil {
			conn, err = net.Dial("tcp", ":8881")
			continue

		}
		_, err = conn.Write([]byte(fmt.Sprintf("我是谁 %v", i)))
		if err != nil {
			conn, err = net.Dial("tcp", ":8881")
			fmt.Println(err)
			continue
		}

		i++
		time.Sleep(time.Second)
	}
}
