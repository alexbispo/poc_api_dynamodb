package main

import (
	"log"

	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
)

var dbSvc *dynamodb.DynamoDB

func init() {
	sess, err := session.NewSession(&aws.Config{
		Region:   aws.String("us-west-2"),
		Endpoint: aws.String("http://localhost:8000")})
	if err != nil {
		log.Println(err)
		return
	}
	dbSvc = dynamodb.New(sess)
}

// func RetrieveAccountById(id int) (account Account, err error) {
// 	result, err := dbSvc.GetItem(&dynamodb.GetItemInput{
// 		TableName: aws.String("Accounts"),
// 		Key: map[string]*dynamodb.AttributeValue{
// 			"AccountId": {
// 				N: aws.IntValue(id),
// 			},
// 		},
// 	})
// 	if err != nil {
// 		fmt.Println(err.Error())
// 		return
// 	}
// }

func RetrieveAccountByID(id int) (account Account, err error) {
	account = Account{ID: id, Name: "Alex"}
	return
}
