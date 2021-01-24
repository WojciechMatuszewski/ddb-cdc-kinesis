// Package database provides support for access the database.
package database

import (
	"context"

	"github.com/aws/aws-sdk-go-v2/aws"
	"github.com/aws/aws-sdk-go-v2/feature/dynamodb/attributevalue"
	"github.com/aws/aws-sdk-go-v2/service/dynamodb"
	"github.com/google/uuid"
)

// Client represents database client
type Client interface {
	PutItem(ctx context.Context, params *dynamodb.PutItemInput, optFns ...func(*dynamodb.Options)) (*dynamodb.PutItemOutput, error)
}

// DB represents the database
type DB struct {
	tableName string
	client    Client
}

// NewService creates DB service
func NewService(client Client, tableName string) *DB {

	return &DB{tableName: tableName, client: client}
}

type item struct {
	PK    string `dynamodbav:"pk"`
	Value string `dynamodbav:"value"`
}

// PutItem puts a random item to the database
func (d DB) PutItem(ctx context.Context) error {
	item := item{
		PK:    uuid.NewString(),
		Value: uuid.NewString(),
	}
	av, err := attributevalue.MarshalMap(item)
	if err != nil {
		return err
	}

	_, err = d.client.PutItem(ctx, &dynamodb.PutItemInput{
		Item:      av,
		TableName: aws.String(d.tableName),
	})
	if err != nil {
		return err
	}

	return nil
}
