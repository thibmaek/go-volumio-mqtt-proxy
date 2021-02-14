package main

import (
	"fmt"
	volumio "go-volumio-mqtt/volumio"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func listen(c *Client) {
	res, err := c.Subscribe("test/topic")

	for update := range res {
		fmt.Printf("Received update: %s\n", update)
	}

	if err != nil {
		fmt.Println(err)
	}
}

// Keeps the connection upon until signal is interrupted or terminated
func keepOpen() chan bool {
	sigs := make(chan os.Signal, 1)
	done := make(chan bool, 1)

	signal.Notify(sigs, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		sig := <-sigs
		fmt.Printf("\n%s\n", sig)
		done <- true
	}()
	return done
}

func updateTrack(client *Client) *time.Ticker {
	ticker := time.NewTicker(time.Second * 1)
	done := make(chan bool)

	go func() {
		for {
			select {
			case <-done:
				return
			case <-ticker.C:
				track, err := volumio.GetCurrentTrack()
				if err != nil {
					fmt.Printf("Failed reading current track: %v", err)
				}

				err = client.Publish("track", Message{
					payload: track,
					retain:  false,
				})
				if err != nil {
					fmt.Printf("Error publishing current track: %v", err)
				}
			}
		}
	}()

	return ticker
}

func main() {
	client := SetupClient()
	defer client.Disconnect()

	err := client.Connect()
	if err != nil {
		panic(err)
	}

	go listen(client)

	ticker := updateTrack(client)
	defer ticker.Stop()

	<-keepOpen()
}
