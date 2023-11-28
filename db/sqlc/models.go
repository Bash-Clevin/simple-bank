// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.24.0

package db

import (
	"github.com/jackc/pgx/v5/pgtype"
)

type Account struct {
	ID          int64              `json:"id"`
	Owner       string             `json:"owner"`
	Balance     int64              `json:"balance"`
	Currency    string             `json:"currency"`
	CreatedAt   pgtype.Timestamptz `json:"created_at"`
	CountryCode pgtype.Int4        `json:"country_code"`
}

type Entry struct {
	ID        int64 `json:"id"`
	AccountID int64 `json:"account_id"`
	// can be negative or positive
	Amount    int64              `json:"amount"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}

type Transfer struct {
	ID            int64 `json:"id"`
	FromAccountID int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	// must be positive
	Amount    int64              `json:"amount"`
	CreatedAt pgtype.Timestamptz `json:"created_at"`
}

type User struct {
	Username          string             `json:"username"`
	HashedPassword    string             `json:"hashed_password"`
	FullName          string             `json:"full_name"`
	Email             string             `json:"email"`
	PasswordChangedAt pgtype.Timestamptz `json:"password_changed_at"`
	CreatedAt         pgtype.Timestamptz `json:"created_at"`
}
