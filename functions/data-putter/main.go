package main

import (
	"context"
	"ddb-kinesis/foundation/database"
	"os"

	"github.com/aws/aws-lambda-go/lambda"
	"github.com/aws/aws-sdk-go-v2/config"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
)

func main() {
	cfg, err := config.LoadDefaultConfig(context.TODO())
	if err != nil {
		panic(err)
	}

	dynamo := dynamodb.NewFromConfig(cfg)
	dbService := database.NewService(dynamo, os.Getenv("TABLE_NAME"))

	h := newHandler(dbService)
	lambda.Start(h)
}
