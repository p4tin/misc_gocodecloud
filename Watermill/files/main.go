package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"strings"

	"github.com/ThreeDotsLabs/watermill"
	"github.com/ThreeDotsLabs/watermill-io/pkg/io"
)

func main() {
	if len(os.Args) < 2 {
		panic(
			fmt.Errorf("usage: %s /path/to/file.log", os.Args[0]),
		)
	}
	logFile, err := os.OpenFile(os.Args[1], os.O_RDONLY, 0444)
	if err != nil {
		panic(err)
	}

	sub, err := io.NewSubscriber(logFile, io.SubscriberConfig{
		UnmarshalFunc: io.PayloadUnmarshalFunc,
	}, watermill.NewStdLogger(true, false))
	if err != nil {
		panic(err)
	}

	// for io.Subscriber, topic does not matter
	lines, err := sub.Subscribe(context.Background(), "")
	if err != nil {
		panic(err)
	}

	for line := range lines {
		if criterion(string(line.Payload)) {
			//_ = alertPublisher.Publish("alerts", line)
			log.Println(string(line.Payload))
		}
		line.Ack()
	}
}

func criterion(line string) bool {
	if strings.Contains(strings.ToUpper(line), "ERROR") {
		return true
	}
	return false
}
