package main

import (
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

var HandleMessages mqtt.MessageHandler = func(c mqtt.Client, m mqtt.Message) {
	fmt.Printf("Received message %s with payload %s", m.Topic(), m.Payload())
}

func SendMessage(c mqtt.Client, topic string, payload interface{}) {
	token := c.Publish(topic, 0, false, payload)
	token.Wait()

	err := token.Error()
	if err != nil {
		panic(token.Error())
	}

	time.Sleep(time.Second)
}

func Subscribe(c mqtt.Client, topic string, cb mqtt.MessageHandler) mqtt.Token {
	token := c.Subscribe(topic, 0, cb)
	return token
	// token.Wait()

	// err := token.Error()
	// if err != nil {
	// 	panic(token.Error())
	// }

	// fmt.Printf("Subscribed to %s", topic)
}
