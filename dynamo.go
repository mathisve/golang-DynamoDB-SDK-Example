package main

import (
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
	"strconv"
)

var dynamo *dynamodb.DynamoDB

// connectDynamo returns a dynamoDB client
func connectDynamo() (db *dynamodb.DynamoDB) {
	return dynamodb.New(session.Must(session.NewSession(&aws.Config{
		Region: &RegionName,
	})))
}

// CreateTable creates a table
func CreateTable() error {
	_, err := dynamo.CreateTable(&dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("Id"),
				AttributeType: aws.String("N"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("Id"),
				KeyType: aws.String("HASH"),
			},
		},
		BillingMode: aws.String(dynamodb.BillingModePayPerRequest),
		TableName: &TableName,
	})

	return err
}

// PutItem inserts the struct Person
func PutItem(person Person) error {
	_, err := dynamo.PutItem(&dynamodb.PutItemInput{
		Item: map[string]*dynamodb.AttributeValue {
			"Id": {
				N: aws.String(strconv.Itoa(person.Id)),
			},
			"Name": {
				S: aws.String(person.Name),
			},
			"Website": {
				S: aws.String(person.Website),
			},
		},
		TableName: &TableName,
	})

	return err
}

// UpdateItem updates the Person based on the Person.Id
func UpdateItem(person Person) error {
	_, err := dynamo.UpdateItem(&dynamodb.UpdateItemInput{
		ExpressionAttributeNames: map[string]*string {
			"#N": aws.String("Name"),
			"#W": aws.String("Website"),
		},
		ExpressionAttributeValues: map[string]*dynamodb.AttributeValue{
			":Name": {
				S: aws.String(person.Name),
			},
			":Website": {
				S: aws.String(person.Website),
			},
		},
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				N: aws.String(strconv.Itoa(person.Id)),
			},
		},
		TableName: &TableName,
		UpdateExpression: aws.String("SET #N = :Name, #W = :Website"),
	})

	return err
}

// DeleteItem deletes a Person based on the Person.Id
func DeleteItem(person Person) error {
	_, err := dynamo.DeleteItem(&dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				N: aws.String(strconv.Itoa(person.Id)),
			},
		},
		TableName: &TableName,
	})

	return err
}

// GetItem gets a Person based on the Id, returns a person
func GetItem(id int) (person Person, err error) {
	result, err := dynamo.GetItem(&dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"Id": {
				N: aws.String(strconv.Itoa(id)),
			},
		},
		TableName: &TableName,
	})

	if err != nil {
		return person, err
	}

	err = dynamodbattribute.UnmarshalMap(result.Item, &person)

	return person, err

}
