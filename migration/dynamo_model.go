package migration

import (
	"astragpt/http_server/database"
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/awserr"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type ModelInfo struct {
	Model     interface{}
	TableName string
}

func Migrate(modelsInfo ...ModelInfo) {
	for _, modelInfo := range modelsInfo {
		err := createTable(modelInfo.Model, modelInfo.TableName)
		if err != nil {
			log.Fatalf("Failed to create table %s: %v", modelInfo.TableName, err)
		}
	}
}

func createTable(model interface{}, tableName string) error {
	// Check if the table exists
	_, err := database.DB.DescribeTable(&dynamodb.DescribeTableInput{
		TableName: aws.String(tableName),
	})

	// If the table doesn't exist, create it
	if err != nil {
		if aerr, ok := err.(awserr.Error); ok && aerr.Code() == dynamodb.ErrCodeResourceNotFoundException {
			_, err := database.DB.CreateTable(createTableInput(model, tableName))
			if err != nil {
				return err
			}
		} else {
			return err
		}
	}

	return nil
}

func createTableInput(model interface{}, tableName string) *dynamodb.CreateTableInput {
	tags := ParseDynamoTags(model)
	attributeDefinitions := []*dynamodb.AttributeDefinition{}
	keySchema := []*dynamodb.KeySchemaElement{}

	for _, tag := range tags {
		attributeDefinitions = append(attributeDefinitions, &dynamodb.AttributeDefinition{
			AttributeName: aws.String(tag.AttributeName),
			AttributeType: aws.String(tag.AttributeType),
		})

		if tag.HashKey {
			keySchema = append(keySchema, &dynamodb.KeySchemaElement{
				AttributeName: aws.String(tag.AttributeName),
				KeyType:       aws.String("HASH"),
			})
		}
	}

	return &dynamodb.CreateTableInput{
		AttributeDefinitions: attributeDefinitions,
		KeySchema:            keySchema,
		BillingMode:          aws.String("PAY_PER_REQUEST"),
		TableName:            aws.String(tableName),
	}
}
