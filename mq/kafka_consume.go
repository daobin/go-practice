package main

import (
	"fmt"
	"github.com/IBM/sarama"
	"log"
)

func main() {
	brokers := []string{"127.0.0.1:9092"}
	consumer, err := sarama.NewConsumer(brokers, nil)
	if err != nil {
		log.Println("NewConsumer Error: ", err.Error())
		return
	}

	topPic := "fish-topic"
	pc, err := consumer.ConsumePartition(topPic, 0, sarama.OffsetNewest)
	for {
		msg := <-pc.Messages()
		fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
	}

}
