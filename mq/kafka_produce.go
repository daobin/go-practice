package main

import (
	"github.com/IBM/sarama"
	"log"
)

func main() {
	brokers := []string{"127.0.0.1:9092"}

	conf := sarama.NewConfig()
	conf.Producer.RequiredAcks = sarama.WaitForAll
	conf.Producer.Return.Successes = true

	cli, err := sarama.NewSyncProducer(brokers, conf)
	if err != nil {
		log.Println("NewSyncProducer Error: ", err.Error())
		return
	}
	defer cli.Close()

	pid, offset, err := cli.SendMessage(&sarama.ProducerMessage{
		Topic: "fish-topic",
		Value: sarama.StringEncoder("Hello, Kafka ..."),
	})
	if err != nil {
		log.Println("SendMessage Error: ", err.Error())
		return
	}

	log.Println("Partition: ", pid)
	log.Println("Offset: ", offset)
}
