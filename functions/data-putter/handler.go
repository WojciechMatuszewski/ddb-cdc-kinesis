package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/aws/aws-lambda-go/events"
)

// Handler is the lambda handler for the data-putter lambda
type Handler func(ctx context.Context) (events.APIGatewayV2HTTPResponse, error)

type Putter interface {
	PutItem(ctx context.Context) error
}

func newHandler(putter Putter) Handler {
	return func(ctx context.Context) (events.APIGatewayV2HTTPResponse, error) {
		err := putter.PutItem(ctx)
		if err != nil {
			fmt.Println(err.Error())
			return events.APIGatewayV2HTTPResponse{
				Body:       fmt.Sprintf("Failed to put the item, %s", err.Error()),
				StatusCode: http.StatusInternalServerError,
			}, nil
		}

		return events.APIGatewayV2HTTPResponse{
			Body:       "Hi there from a handler!",
			StatusCode: http.StatusOK,
		}, nil
	}
}
