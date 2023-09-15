package models

import (
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
)

type Transaction struct {
	ID              primitive.ObjectID `json:"id" bson:"_id"`
	TransactionDate string             `json:"transaction_date" bson:"transaction_date"`
	Amount          float64            `json:"amount" bson:"amount"`
	TransactionType Type               `json:"transaction_type" bson:"transaction_type"`
	CustomerID      primitive.ObjectID `json:"customer_id" bson:"customer_id"`
}

type Type struct {
	Transfer   string `json:"transfer" bson:"transfer"`
	Deposit    string `json:"deposit" bson:"deposit"`
	Withdrawal string `json:"withdrawal" bson:"withdrawal"`
}

type DBResponse struct {
	ID                 primitive.ObjectID `json:"id" bson:"_id"`
	Name               string             `json:"name" bson:"name"`
	Email              string             `json:"email" bson:"email"`
	Password           string             `json:"password" bson:"password"`
	PasswordConfirm    string             `json:"passwordConfirm,omitempty" bson:"passwordConfirm,omitempty"`
	Role               string             `json:"role" bson:"role"`
	VerificationCode   string             `json:"verificationCode,omitempty" bson:"verificationCode"`
	ResetPasswordToken string             `json:"resetPasswordToken,omitempty" bson:"resetPasswordToken,omitempty"`
	ResetPasswordAt    time.Time          `json:"resetPasswordAt,omitempty" bson:"resetPasswordAt,omitempty"`
	Verified           bool               `json:"verified" bson:"verified"`
	CreatedAt          time.Time          `json:"created_at" bson:"created_at"`
	UpdatedAt          time.Time          `json:"updated_at" bson:"updated_at"`
}
