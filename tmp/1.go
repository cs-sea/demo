package main

import (
	"fmt"
	"google.golang.org/grpc"
	"time"
)

func main() {
	var opt []grpc.DialOption

	opt = append(opt, grpc.WithInsecure())
	opt = append(opt, grpc.WithBlock())
	conn, err := grpc.Dial(":8888", opt...)
	fmt.Println(err)
	fmt.Println(conn)
	time.Sleep(time.Hour)
	conn.Close()
}
