
package main

import (
	"log"
	"strconv"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// var params = {
// 	TableName: 'accounts',
// 	Item: { // a map of attribute name to AttributeValue
  
// 		AccountId: 3,
// 		ContractId: 2,
// 		MainAccountId: 1,
// 		PersonId: 2,
// 		Name: "Alex Bispo",
// 		Description: "Desenvolvedor"
// 	},
//   };
//   docClient.put(params, function(err, data) {
// 	if (err) ppJson(err); // an error occurred
// 	else ppJson(data); // successful response
//   });

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

func RetrieveAccountById(id int) (account Account, err error) {
	result, err := dbSvc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("accounts"),
		Key: map[string]*dynamodb.AttributeValue{
			"AccountId": {
				N: aws.String(strconv.Itoa(id)),
			},
		},
	})
	if err != nil {
		log.Println(err.Error())
		return
	}

	account = Account{}

	log.Println(result)

	log.Println(result.Item)

	err = dynamodbattribute.UnmarshalMap(result.Item, &account)
	if err != nil {
		log.Println(err)
	}

	log.Println(account)

	return
}

// func RetrieveAccountByID(id int) (account Account, err error) {
// 	account = Account{ID: id, Name: "Alex"}
// 	return
// }
