package db

import (
	"context"
	"fmt"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	*Queries
	db *pgxpool.Pool
}

func NewStore(db *pgxpool.Pool) *Store {
	return &Store{
		db:      db,
		Queries: New(db),
	}
}

func (store *Store) execTx(ctx context.Context, fn func(*Queries) error) error {
	tx, err := store.db.Begin(ctx)
	if err != nil {
		return err
	}
	q := New(tx)
	err = fn(q)
	if err != nil {
		if rbErr := tx.Rollback(ctx); rbErr != nil {
			return fmt.Errorf("tx err: %v, rb err: %v", rbErr, rbErr)
		}
		return err
	}
	return tx.Commit(ctx)
}

type TransferTxParams struct {
	FromAccountId int64 `json:"from_account_id"`
	ToAccountID   int64 `json:"to_account_id"`
	Amount        int64 `json:"amount"`
}

type TransferTxResult struct {
	Transfer   Transfer `json:"transfer"`
	FromAmount Account  `json:"from_account"`
	ToAccount  Account  `json:"to_account"`
	FromEntry  Entry    `json:"from_entry"`
	ToEntry    Entry    `json:"to_entry"`
}

// func (store *Store) TransferTx(ctx context.Context, args TransferTxParams) (TransferTxResult, error) {
// 	var result TransferTxResult
// 	store.execTx(ctx, func(q *Queries) error {
// 		result.Transfer, err = q.CreateAcount(ctx, CreateTra{})
// 		return nil
// 	})
// }
