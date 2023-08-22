package main

import (
	"context"
	"fmt"
	"github.com/IBM/sarama"
	"log"
	"time"
)

type cGroupHandler struct{}

func (c *cGroupHandler) Setup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c *cGroupHandler) Cleanup(_ sarama.ConsumerGroupSession) error {
	return nil
}

func (c *cGroupHandler) ConsumeClaim(sess sarama.ConsumerGroupSession, claim sarama.ConsumerGroupClaim) error {
	for msg := range claim.Messages() {
		// todo 可做一些在消费标记之前的业务逻辑处理
		fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
		sess.MarkMessage(msg, "")
	}

	return nil
}

func main() {
	brokers := []string{"127.0.0.1:9092"}
	topPic := "fish-topic"

	conf := sarama.NewConfig()
	conf.Consumer.Offsets.AutoCommit.Enable = true
	conf.Consumer.Offsets.AutoCommit.Interval = 1 * time.Second
	conf.Consumer.Offsets.Initial = sarama.OffsetNewest
	conf.Consumer.Offsets.Retry.Max = 3

	// 使用消费组
	for {
		cGroup, err := sarama.NewConsumerGroup(brokers, "user", conf)
		if err != nil {
			log.Println("NewConsumerGroup Error: ", err.Error())
			return
		}

		err = cGroup.Consume(context.Background(), []string{topPic}, &cGroupHandler{})
		if err != nil {
			log.Println("Consume Error: ", err.Error())
			return
		}

		time.Sleep(time.Second)
	}

	// 不使用消费组
	//consumer, err := sarama.NewConsumer(brokers, conf)
	//if err != nil {
	//	log.Println("NewConsumer Error: ", err.Error())
	//	return
	//}
	//
	//pc, err := consumer.ConsumePartition(topPic, 0, sarama.OffsetNewest)
	//for {
	//	msg := <-pc.Messages()
	//	fmt.Printf("Partition:%d Offset:%d Key:%v Value:%v\n", msg.Partition, msg.Offset, msg.Key, string(msg.Value))
	//}

}
