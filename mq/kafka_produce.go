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

	//producer, err := sarama.NewSyncProducer(brokers, conf)
	//if err != nil {
	//	log.Println("NewSyncProducer Error: ", err.Error())
	//	return
	//}

	cli, err := sarama.NewClient(brokers, conf)
	if err != nil {
		log.Println("NewClient Error: ", err.Error())
		return
	}

	producer, err := sarama.NewSyncProducerFromClient(cli)
	if err != nil {
		log.Println("NewSyncProducer Error: ", err.Error())
		return
	}

	defer producer.Close()

	partition, offset, err := producer.SendMessage(&sarama.ProducerMessage{
		Topic:     "fish-topic",
		Value:     sarama.StringEncoder("Hello, 2"),
		Partition: 1,
	})
	if err != nil {
		log.Println("SendMessage Error: ", err.Error())
		return
	}

	log.Println("Partition: ", partition)
	log.Println("Offset: ", offset)
}
