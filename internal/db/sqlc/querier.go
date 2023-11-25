// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.20.0

package db

import (
	"context"
)

type Querier interface {
	CreateAcount(ctx context.Context, arg CreateAcountParams) (Account, error)
	DeleteAccount(ctx context.Context, id int64) error
	GetAccount(ctx context.Context, id int64) (Account, error)
	ListAccounts(ctx context.Context, arg ListAccountsParams) ([]Account, error)
	UpdateAccount(ctx context.Context, arg UpdateAccountParams) (Account, error)
}

var _ Querier = (*Queries)(nil)
