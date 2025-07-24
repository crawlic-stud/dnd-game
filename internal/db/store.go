package db

import (
	"context"

	"github.com/jackc/pgx/v5/pgxpool"
)

type Store struct {
	*Queries
	Pool *pgxpool.Pool
}

func (s *Store) Transaction(ctx context.Context, txFunc func(tx *Queries) error) error {
	tx, err := s.Pool.Begin(ctx)
	if err != nil {
		return err
	}
	defer tx.Rollback(ctx)

	qTx := s.Queries.WithTx(tx)

	err = txFunc(qTx)
	if err != nil {
		return err
	}

	return tx.Commit(ctx)
}
