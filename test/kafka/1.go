package main

import (
	"context"
	"fmt"

	"github.com/segmentio/kafka-go"
)

func main() {

	cus := kafka.NewReader(kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		GroupID:  "test1",
		Topic:    "test1",
		MinBytes: 10e3,
		MaxBytes: 10e6,
	})
	fmt.Println(cus.Offset())

	for true {
		msgs, err := cus.ReadMessage(context.Background())
		if err != nil {
			panic(err)
		}
		fmt.Println(msgs.Offset)
		fmt.Println(string(msgs.Value))
		cus.CommitMessages(context.Background(), msgs)
	}

}
