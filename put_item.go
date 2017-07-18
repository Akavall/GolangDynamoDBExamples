package main

import (
	"fmt"

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

	r := Record{
		Year:  1991,
		Title: "Generator",
		Songs: []string{"No Direction", "The Answer"},
	}

	av, err := dynamodbattribute.MarshalMap(r)
	if err != nil {
		panic(fmt.Sprintf("failed to DynamoDB marshal Record, %v", err))
	}

	_, err = svc.PutItem(&dynamodb.PutItemInput{
		TableName: aws.String("Albums"),
		Item:      av,
	})

	if err != nil {
		panic(fmt.Sprintf("failed to put Record to DynamoDB, %v", err))
	}

}
