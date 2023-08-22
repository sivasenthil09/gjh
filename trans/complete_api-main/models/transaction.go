package models

import (
	"time"
)

type Transaction struct{
	ID  int `json:"id" bson:"id"`
	Transaction_date time.Time `json:"transaction_date" bson:"transaction_date"`
	Amount int `json:"amount" bson:"amount"`
	Transaction_Type string `json:"transaction_type" bson:"transaction_type"`
	Customer_ID int `json:"customer_id" bson:"customer_id"`
}