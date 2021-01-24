package main

import (
	"context"
	"fmt"

	"github.com/aws/aws-lambda-go/events"
)

type Handler func(ctx context.Context, event events.KinesisEvent) error

func NewHandler() Handler {
	return func(ctx context.Context, event events.KinesisEvent) error {
		for _, record := range event.Records {
			dataText := string(record.Kinesis.Data)

			fmt.Printf("%s Data = %s \n", record.EventName, dataText)
		}

		return nil
	}
}
