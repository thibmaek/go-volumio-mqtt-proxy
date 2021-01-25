package main

import (
	"errors"
	"fmt"
	"time"

	mqtt "github.com/eclipse/paho.mqtt.golang"
)

type Client struct {
	client  mqtt.Client
	updates chan string
}

// SetupClient creates a new MQTT connection and returns a client
func SetupClient() *Client {
	brokerURL := fmt.Sprintf("tcp://%s:%d", broker, port)

	clientOpts := mqtt.NewClientOptions()
	clientOpts.AddBroker(brokerURL)
	clientOpts.SetMaxReconnectInterval(time.Second * 1)
	clientOpts.SetClientID(clientId)

	if len(user) > 0 {
		clientOpts.SetUsername(user)
	}
	if len(password) > 0 {
		clientOpts.SetPassword(password)
	}

	c := mqtt.NewClient(clientOpts)
	u := make(chan string)

	return &Client{
		client:  c,
		updates: u,
	}
}

func (c *Client) Connect() error {
	if token := c.client.Connect(); token.Wait() && token.Error() != nil {
		return token.Error()
	}

	fmt.Println("+------------------+")
	fmt.Println("|     Connected    |")
	fmt.Println("+------------------+")

	return nil
}

func (c *Client) Disconnect() {
	if c.client != nil {
		c.client.Disconnect(uint(time.Second * 1))
		fmt.Println("Disconnected")
	}
}

func (c *Client) Subscribe(topic string) (chan string, error) {
	if !c.client.IsConnected() {
		return nil, errors.New("Not connected")
	}

	ch := make(chan string)

	token := c.client.Subscribe(topic, 1, func(_ mqtt.Client, m mqtt.Message) {
		ch <- string(m.Payload())
	})
	token.Wait()
	return ch, token.Error()
}

type Message struct {
	payload string
	retain  bool
}

func (c *Client) Publish(t string, m Message) error {
	if !c.client.IsConnected() {
		return errors.New("Not connecte")
	}

	token := c.client.Publish(t, 1, m.retain, []byte(m.payload))
	token.Wait()
	return token.Error()
}
