package main

import (
	"fmt"
	"reflect"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

type Record struct {
	Year  int    `dynamodbav:"year"`
	Title string `dynamodbav:"title"`
	Songs []string
}

func main() {

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(endpoints.UsWest2RegionID),
	}))

	svc := dynamodb.New(sess)

	input := &dynamodb.GetItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"year": {
				N: aws.String("1991"),
			},
			"title": {
				S: aws.String("Generator"),
			},
		},
		TableName: aws.String("Albums"),
	}

	result, err := svc.GetItem(input)

	fmt.Println(reflect.TypeOf(result))

	if err != nil {
		panic(fmt.Sprintf("failed to get Record from DynamoDB, %v", err))
	}

	var myRecord Record

	dynamodbattribute.UnmarshalMap(result.Item, &myRecord)

	fmt.Println(reflect.TypeOf(myRecord))
	fmt.Println(myRecord)
}
