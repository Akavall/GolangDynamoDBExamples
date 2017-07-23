package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

type Record struct {
	Artist string   `dynamodbav:"artist"`
	Year   int      `dynamodbav:"year"`
	Title  string   `dynamodbav:"title"`
	Songs  []string `dynamodbav:"songs"`
}

func main() {

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(endpoints.UsWest2RegionID),
	}))

	svc := dynamodb.New(sess)

	input := &dynamodb.DeleteItemInput{
		Key: map[string]*dynamodb.AttributeValue{
			"artist": {
				S: aws.String("Bad Religion"),
			},
			"year": {
				N: aws.String("1991"),
			},
		},
		TableName: aws.String("Albums"),
	}

	result, err := svc.DeleteItem(input)

	fmt.Println(result)

	if err != nil {
		panic(fmt.Sprintf("failed to delete Record from DynamoDB, %v", err))
	} else {
		fmt.Println("Item deleted successfully")
	}
}
