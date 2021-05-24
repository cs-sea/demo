package main

import (
	"fmt"
	"net"
	"time"
)

func main() {
	listener, err := net.Listen("tcp", ":8881")
	if err != nil {
		panic("sdfsdf")
	}

	var conn net.Conn
	conn, err = listener.Accept()
	for true {
		buf := make([]byte, 20)
		_, err = conn.Read(buf)
		if err != nil {
			conn, err = listener.Accept()
			fmt.Println(err)
			continue
		}
		fmt.Println(string(buf))
		time.Sleep(time.Second * 5)
	}
}
