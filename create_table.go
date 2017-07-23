package main

import (
	"fmt"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/endpoints"
	"github.com/aws/aws-sdk-go/aws/session"

	"github.com/aws/aws-sdk-go/service/dynamodb"
)

func main() {

	sess := session.Must(session.NewSession(&aws.Config{
		Region: aws.String(endpoints.UsWest2RegionID),
	}))

	svc := dynamodb.New(sess)

	// Notes on Hash and Range Keys:
	// Dynamodb, partitions and does look ups based on the hash Key
	// therefore, we need Hash Key to be well distributed
	// Artist name seems like in is very well distributed,
	// Artist gender would be a bad choice

	// Range Key would allows to make queries for a given HashKey,
	// For example: Give me all for Artist: Bad Religion give me all albums
	// after 1995

	// We don't have to have Range Key at all, but we have to have a HashKey.

	// The combination of HashKey and Range Key (Primary Key), has to be unique,
	// otherwise we end up losing records,
	// The combiantion of Artist and Year, is not ideal, because it relies on the
	// assumption that no Artist released two Albums in the same year.

	input := &dynamodb.CreateTableInput{
		AttributeDefinitions: []*dynamodb.AttributeDefinition{
			{
				AttributeName: aws.String("artist"),
				AttributeType: aws.String("S"),
			},
			{
				AttributeName: aws.String("year"),
				AttributeType: aws.String("N"),
			},
		},
		KeySchema: []*dynamodb.KeySchemaElement{
			{
				AttributeName: aws.String("artist"), // Hash Key or Partition Key
				KeyType:       aws.String("HASH"),
			},
			{
				AttributeName: aws.String("year"),
				KeyType:       aws.String("RANGE"),
			},
		},
		ProvisionedThroughput: &dynamodb.ProvisionedThroughput{
			ReadCapacityUnits:  aws.Int64(1),
			WriteCapacityUnits: aws.Int64(1),
		},
		TableName: aws.String("Albums"),
	}

	result, err := svc.CreateTable(input)

	if err != nil {
		panic(fmt.Sprintf("failed to create table, %v", err))
	}

	fmt.Println("result: ", result)
}
