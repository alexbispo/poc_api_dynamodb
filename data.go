
package main

import (
	"log"
	// "strconv"
	"github.com/aws/aws-sdk-go/aws"
	"github.com/aws/aws-sdk-go/aws/session"
	"github.com/aws/aws-sdk-go/service/dynamodb"
	"github.com/aws/aws-sdk-go/service/dynamodb/dynamodbattribute"
)

// var params = {
//     TableName: 'card_details',
//     KeySchema: [ // The type of of schema.  Must start with a HASH type, with an optional second RANGE.
//         { // Required HASH type attribute
//             AttributeName: 'indice_pk',
//             KeyType: 'HASH',
//         },
//         { // Required HASH type attribute
//             AttributeName: 'indice_sk',
//             KeyType: 'RANGE',
//         }
//     ],
//     AttributeDefinitions: [ // The names and types of all primary and index key attributes only
//         {
//             AttributeName: 'indice_pk',
//             AttributeType: 'S', // (S | N | B) for string, number, binary
//         },
//         {
//             AttributeName: 'indice_sk',
//             AttributeType: 'S', // (S | N | B) for string, number, binary
//         }

//     ],
//     ProvisionedThroughput: { // required provisioned throughput for the table
//         ReadCapacityUnits: 5,
//         WriteCapacityUnits: 5,
//     }
// };
// dynamodb.createTable(params, function(err, data) {
//     if (err) ppJson(err); // an error occurred
//     else ppJson(data); // successful response

// });

// var params = {
// 	TableName: 'card_details',
// 	Item: { // a map of attribute name to AttributeValue
  
	// indice_pk: "card_id",
	// indice_sk: "1234",
	// account_id: 120,
	// card_id: 1234,
	// contract_id: 100,
	// main_account_id: 10,
	// external_code: "1xxx"
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

func RetrieveCardById(id string) (Card, error) {
	result, err := dbSvc.GetItem(&dynamodb.GetItemInput{
		TableName: aws.String("cards_detail"),
		Key: map[string]*dynamodb.AttributeValue{
			"indice_pk": {
				S: aws.String("card_id"),
			},
			"indice_sk": {
				S: aws.String(id),
			},
		},
	})
	if err != nil {
		log.Println(err.Error())
		return Card{}, err
	}

	card := Card{}

	log.Println(result)

	log.Println(result.Item)

	err = dynamodbattribute.UnmarshalMap(result.Item, &card)
	if err != nil {
		log.Println(err)
	}

	if  card.account_id != 0 {
		log.Println(card.account_id)
	}

	log.Println(card)

	return card, nil
}

// func RetrieveAccountByID(id int) (card Account, err error) {
// 	card = Account{ID: id, Name: "Alex"}
// 	return
// }
