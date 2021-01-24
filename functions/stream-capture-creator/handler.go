package main

import (
	"context"

	"github.com/aws/aws-lambda-go/cfn"
	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func NewHandler(client *dynamodb.Client) cfn.CustomResourceFunction {
	return func(ctx context.Context, event cfn.Event) (physicalResourceID string, data map[string]interface{}, err error) {
		if event.RequestType != cfn.RequestCreate {
			return
		}

		streamArn, found := event.ResourceProperties["StreamArn"].(string)
		if !found {
			return event.PhysicalResourceID, nil, err
		}
		tableName, found := event.ResourceProperties["TableName"].(string)
		if !found {
			return event.PhysicalResourceID, nil, err
		}

		_, enableErr := client.EnableKinesisStreamingDestination(ctx, &dynamodb.EnableKinesisStreamingDestinationInput{
			StreamArn: aws.String(streamArn),
			TableName: aws.String(tableName),
		})

		if enableErr != nil {
			return event.PhysicalResourceID, nil, err
		}
		return
	}
}
