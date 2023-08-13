package token

import (
	"time"
)

type PayLoad struct {
	UserID    string    `json:"account_id"`
	UserName  string    `json:"user_name"`
	UserEmail string    `json:"user_email"`
	IssuedAt  time.Time `json:"issued_at"`
	Expiry    time.Time `json:"expiry"`
}

func newPayload(userID, userName, userEmail string, issuedAt, expiry time.Time) *PayLoad {
	return &PayLoad{
		UserID:    userID,
		UserName:  userName,
		UserEmail: userEmail,
		IssuedAt:  issuedAt,
		Expiry:    expiry,
	}
}
