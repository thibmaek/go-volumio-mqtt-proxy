package main

import "fmt"

const clientId = "go-volumio-mqtt"

var (
	broker   string
	user     string
	password string
	port     = 1883
)

func listen(c *Client) {
	res, err := c.Subscribe("test/topic")

	for update := range res {
		fmt.Printf("Received update: %s", update)
	}

	if err != nil {
		fmt.Println(err)
	}
}

func main() {
	client := SetupClient()
	defer client.Disconnect()

	err := client.Connect()
	if err != nil {
		panic(err)
	}

	go listen(client)

	for {
	}
}
