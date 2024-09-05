package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID          primitive.ObjectID `json:"id" bson:"_id,omitempty"`
	Amount      float64           `json:"amount" bson:"amount"`
	Type        string            `json:"type" bson:"type"` // "income" or "expense"
	Category    string            `json:"category" bson:"category"`
	Subcategory string            `json:"subcategory,omitempty" bson:"subcategory,omitempty"`
	Account     string            `json:"account" bson:"account"`
	Description string            `json:"description" bson:"description"`
	Date        time.Time         `json:"date" bson:"date"`
	CreatedAt   time.Time         `json:"createdAt" bson:"createdAt"`
	UpdatedAt   time.Time         `json:"updatedAt" bson:"updatedAt"`
