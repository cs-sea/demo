package main

import (
	"context"

	"github.com/segmentio/kafka-go"
)

func main() {
	// to produce messages

	//conn, err := kafka.DialLeader(context.Background(), "tcp", "localhost:9092", "test1", 1)
	//if err != nil {
	//	log.Fatal("failed to dial leader:", err)
	//}
	//
	//fmt.Println(conn.Offset())
	//
	//bs, err := conn.Brokers()
	//if err != nil {
	//	panic(err)
	//}
	//fmt.Println(len(bs))
	//
	writer := &kafka.Writer{
		Addr:     kafka.TCP("localhost:9092"),
		Topic:    "test1",
		Balancer: kafka.CRC32Balancer{},
	}

	writer.WriteMessages(context.Background(), kafka.Message{Value: []byte("one!")},
		kafka.Message{Value: []byte("two!")},
		kafka.Message{Value: []byte("three!")})
}
