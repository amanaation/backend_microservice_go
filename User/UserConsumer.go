package main

import (
	//"github.com/confluentinc/confluent-kafka-go/kafka"
	"context"
	"strings"

	"github.com/segmentio/kafka-go"
)

func consumingUser() {

	conf := kafka.ReaderConfig{
		Brokers:  []string{"localhost:9092"},
		Topic:    "User",
		MaxBytes: 10,
	}

	reader := kafka.NewReader(conf)

	for {
		m, err := reader.ReadMessage(context.Background())
		if err != nil {
			panic(err)
			// continue
		}
		user(formatValues(string(m.Value)))
	}

}

func formatValues(v string) map[string]string {

	values := make(map[string]string)
	//v := "/user?uid=2&amount=30000"
	val := strings.Split(strings.Split(v, "?")[1], "&")
	//fmt.Println(val)
	var sub []string
	for i := 0; i < len(val); i++ {
		sub = strings.Split(val[i], "=")
		values[sub[0]] = string(sub[1])
	}
	return (values)
}
