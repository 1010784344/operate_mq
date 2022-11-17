package main

import (
	"encoding/json"
	"fmt"
	"github.com/streadway/amqp"
	"log"
)

func main() {
	uri := "amqp://guest:guest@127.0.0.1:5672/"

	routing_key := "TEST01"
	content := map[string]interface{}{
		"name": "zelda",
	}

	err := Pub_mq(uri,routing_key,content)
	fmt.Println(err)
}

// 生产者
func Pub_mq(uri,routing_key string, content map[string]interface{}) error {
	// 建立连接
	connection, err := amqp.Dial(uri)
	if err != nil {
		log.Println("Failed to connect to RabbitMQ:", err.Error())
		return err
	}
	defer connection.Close()
	// 创建一个Channel
	channel, err := connection.Channel()
	if err != nil {
		log.Println("Failed to open a channel:", err.Error())
		return err
	}
	defer channel.Close()


	// 发送
	messageBody,err := json.Marshal(content)
	if err = channel.Publish(
		"",    // exchange
		routing_key, // routing key
		false,       // mandatory
		false,       // immediate
		amqp.Publishing{
			Headers:         amqp.Table{},
			ContentType:     "text/plain",
			ContentEncoding: "",
			Body:            []byte(messageBody),
			//Expiration:      "60000", // 消息过期时间
		},
	); err != nil {
		log.Println("Failed to publish a message:", err.Error())
		return err
	}
	return nil
}
