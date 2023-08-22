package services

import (
	"context"
	"errors"
	"rest-api/interfaces"
	"rest-api/models"
	"go.mongodb.org/mongo-driver/mongo"
	"time"
)

type TransactionService struct {
	TransactionCollection *mongo.Collection
	ctx               context.Context
}

func NewTransactionServiceInit(collection *mongo.Collection, ctx context.Context) interfaces.ITransaction {
	return &TransactionService{ collection , ctx}
}

func (p *TransactionService) CreateTransaction(user *models.Transaction) (*models.DBResponse, error) {
	
	
	user.Transaction_date=time.Now()
	_, err := p.TransactionCollection.InsertOne(p.ctx, &user)

	if err != nil {
		if er, ok := err.(mongo.WriteException); ok && er.WriteErrors[0].Code == 11000 {
			return nil, errors.New("user with that email already exist")
		}
		return nil, err
	}
	return nil, err
}



