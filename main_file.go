package main

import (
	"database/sql"
	"fmt"
	"net/http"
	"os"

	"github.com/confluentinc/confluent-kafka-go/kafka"
	_ "github.com/go-sql-driver/mysql"
)

var db, err = sql.Open("mysql", "newuser:password@/demodb")

func main() {

	defer db.Close()
	http.HandleFunc("/user", produceUser)
	http.HandleFunc("/order", produceOrder)

	http.ListenAndServe(":8080", nil)
}

func produceUser(w http.ResponseWriter, r *http.Request) {
	producing("User", r.URL.RequestURI())
	//pvalues(r.URL.Query())

}
func produceOrder(w http.ResponseWriter, r *http.Request) {
	producing("Order", r.URL.RequestURI())
	//pvalues(r.URL.Query())

}

func producing(topic string, value string) {

	p, err := kafka.NewProducer(&kafka.ConfigMap{
		"bootstrap.servers": "localhost:9092",
		"client.id":         "localhost",
		"acks":              "all"})

	if err != nil {
		fmt.Printf("Failed to create producer: %s\n", err)
		os.Exit(1)
	}
	deliveryChan := make(chan kafka.Event, 10000)

	err = p.Produce(&kafka.Message{
		TopicPartition: kafka.TopicPartition{Topic: &topic},
		Value:          []byte(value)},
		deliveryChan,
	)

	e := <-deliveryChan
	m := e.(*kafka.Message)

	if m.TopicPartition.Error != nil {
		fmt.Printf("Delivery failed: %v\n", m.TopicPartition.Error)
	} else {
		fmt.Printf("Delivered message to topic %s [%d] at offset %v\n",
			*m.TopicPartition.Topic, m.TopicPartition.Partition, m.TopicPartition.Offset)
	}

	close(deliveryChan)

}
