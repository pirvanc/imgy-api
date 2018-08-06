package dbaccessor

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

var (
	returnConsumedCapacity = "TOTAL"
	consistentRead         = false
)

// DBAccessor is...
type DBAccessor struct {
	dynamoDBClient dynamodb.DynamoDB
	tableName      string
}

// NewDBAccessor is...
func NewDBAccessor(awsRegion string, tableName string) *DBAccessor {

	// setup dynamoDB client
	dynamoDBClient, err := newDynamoDBClient(awsRegion)
	if err != nil {
		return nil
	}

	return &DBAccessor{
		dynamoDBClient: *dynamoDBClient,
		tableName:      tableName,
	}
}

// newDynamoDBClient is ..
func newDynamoDBClient(awsRegion string) (*dynamodb.DynamoDB, error) {
	session, err := session.NewSession(&aws.Config{
		Region: aws.String(awsRegion)},
	)
	if err != nil {
		return nil, err
	}

	// Create DynamoDB client
	dynamoDBClient := dynamodb.New(session)
	return dynamoDBClient, nil
}

// PutItem is ...
func (dba DBAccessor) PutItem(item interface{}) (interface{}, error) {

	dynamoDBItem, err := dynamodbattribute.MarshalMap(item)
	if err != nil {
		return nil, err
	}

	// create item
	input := &dynamodb.PutItemInput{
		Item:      dynamoDBItem,
		TableName: aws.String(dba.tableName),
	}

	putItemOutput, err := dba.dynamoDBClient.PutItem(input)
	if err != nil {
		return nil, err
	}

	return putItemOutput, nil
}

// GetItemByHashKey ...
func (dba DBAccessor) GetItemByHashKey(hashKeyName string, hashKeyValue string) (interface{}, error) {

	// get item
	consistentRead := false
	returnConsumedCapacity := "TOTAL"
	result, err := dba.dynamoDBClient.GetItem(&dynamodb.GetItemInput{
		TableName:              aws.String(dba.tableName),
		ConsistentRead:         &consistentRead,
		ReturnConsumedCapacity: &returnConsumedCapacity,
		Key: map[string]*dynamodb.AttributeValue{
			hashKeyName: {
				S: aws.String(hashKeyValue),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return result.Item, nil
}

// GetItemByHashKeyAndRangeKey ...
func (dba DBAccessor) GetItemByHashKeyAndRangeKey(hashKeyName string, hashKeyValue string, rangeKeyName string, rangeKeyValue string) (interface{}, error) {

	// get item
	consistentRead := false
	returnConsumedCapacity := "TOTAL"
	result, err := dba.dynamoDBClient.GetItem(&dynamodb.GetItemInput{
		TableName:              aws.String(dba.tableName),
		ConsistentRead:         &consistentRead,
		ReturnConsumedCapacity: &returnConsumedCapacity,
		Key: map[string]*dynamodb.AttributeValue{
			hashKeyName: {
				S: aws.String(hashKeyValue),
			},
			rangeKeyName: {
				S: aws.String(rangeKeyValue),
			},
		},
	})
	if err != nil {
		return nil, err
	}

	return result.Item, nil
}

// DeleteItemByHashKey is ...
func (dba DBAccessor) DeleteItemByHashKey(hashKeyName string, hashKeyValue string) (interface{}, error) {

	// delete item
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			hashKeyName: {
				S: aws.String(hashKeyValue),
			},
		},
		TableName: aws.String(dba.tableName),
	}

	deleteItemOutput, err := dba.dynamoDBClient.DeleteItem(input)
	if err != nil {
		return nil, err
	}

	return deleteItemOutput, nil
}

// DeleteItemByHashKeyAndRangeKey is ...
func (dba DBAccessor) DeleteItemByHashKeyAndRangeKey(hashKeyName string, hashKeyValue string, rangeKeyName string, rangeKeyValue string) (interface{}, error) {

	// delete item
	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			hashKeyName: {
				S: aws.String(hashKeyValue),
			},
			rangeKeyName: {
				S: aws.String(rangeKeyValue),
			},
		},
		TableName: aws.String(dba.tableName),
	}

	deleteItemOutput, err := dba.dynamoDBClient.DeleteItem(input)
	if err != nil {
		return nil, err
	}

	return deleteItemOutput, nil
}
