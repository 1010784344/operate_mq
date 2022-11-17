package main

import (
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	uri := "amqp://guest:guest@127.0.0.1:5672/"

	err := Use_mq(uri)
	fmt.Println(err)
}

// 消费者
func Use_mq(uri string) error {
	// 建立连接
	conn, err := amqp.Dial(uri)
	if err != nil {
		log.Println("Failed to connect to RabbitMQ:", err.Error())
		return err
	}
	defer conn.Close()
	// 启动一个通道
	ch, err := conn.Channel()
	if err != nil {
		log.Println("Failed to open a channel:", err.Error())
		return err
	}


	// 注册消费者
	msgs, err := ch.Consume(
		"TEST01",    // queue
		"hello", // 标签（无实际意义）
		true,      // auto-ack
		false,     // exclusive
		false,     // no-local
		false,     // no-wait
		nil,       // args
	)
	if err != nil {
		log.Println("Failed to register a consumer:", err.Error())
		return err
	}
	//阻塞获取消费者信息
	forever := make(chan bool)
	go func() {
		for d := range msgs {
			log.Println(d.Type)
			log.Println(d.MessageId)
			log.Printf("Received a message: %s", d.Body)
		}
	}()
	log.Printf("Waiting for messages. To exit press CTRL+C")
	<-forever
	return nil
}
