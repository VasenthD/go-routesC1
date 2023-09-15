package services

import (
	"context"
	"errors"
	"fmt"

	"mongodb1/models"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/mongo"
	//"go.mongodb.org/mongo-driver/mongo/options"
)

type Transactionservice struct {
	TransactionCollection *mongo.Collection
	ctx                   context.Context
}

// Initialize the transaction service
func InitProfileServiceInit(collection *mongo.Collection, ctx context.Context) *Transactionservice {
	return &Transactionservice{collection, ctx}
}

func (p *Transactionservice) Createtransaction(user *models.Transaction) (*models.DBResponse, error) {

	res, err := p.TransactionCollection.InsertOne(p.ctx, &user)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("user with that email already exist")
		}
		return nil, err
	}

	var newUser *models.DBResponse
	query := bson.M{"_id": res.InsertedID}

	err = p.TransactionCollection.FindOne(p.ctx, query).Decode(&newUser)
	if err != nil {
		return nil, err
	}
	return newUser, nil
}
func (p *Transactionservice) Gettransaction(user *models.Transaction) int {

	fmt.Println("i got it")
	return 1

}
