from __future__ import print_function # Python 2/3 compatibility
import boto3

dynamodb = boto3.resource('dynamodb', region_name='us-west-2')

table = dynamodb.create_table(
    TableName='Albums',
    KeySchema=[
        {
            'AttributeName': 'artist',
            'KeyType': 'HASH'  #Partition key
        },
        {
            'AttributeName': 'year',
            'KeyType': 'RANGE'  #Sort key
        }
    ],
    AttributeDefinitions=[
        {
            'AttributeName': 'artist',
            'AttributeType': 'S'
        },
        {
            'AttributeName': 'year',
            'AttributeType': 'N'
        },

    ],
    ProvisionedThroughput={
        'ReadCapacityUnits': 1,
        'WriteCapacityUnits': 1
    }
)
