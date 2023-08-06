// Package repository contains methods working with db
package repository

import (
	"context"
	"fmt"
	"time"

	"github.com/google/uuid"
	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/sirupsen/logrus"
)

// BalanceRepository contains pgxpool
type BalanceRepository struct {
	pool *pgxpool.Pool
}

// zero contains zero value
const zero = 0

// NewBalanceRepository creates an object of *ProfileRepository
func NewBalanceRepository(pool *pgxpool.Pool) *BalanceRepository {
	return &BalanceRepository{pool: pool}
}

// AddBalanceChange adds row of balances table using profileID and positive/negative value of amount
func (r *BalanceRepository) AddBalanceChange(ctx context.Context, profileID uuid.UUID, amount float64) error {
	tx, err := r.pool.Begin(ctx)

	defer func() {
		if err != nil {
			errRollback := tx.Rollback(ctx)
			if errRollback != nil {
				logrus.Errorf("BalanceRepository -> Deposit -> %v", errRollback)
			}
		} else {
			errCommit := tx.Commit(ctx)
			if errCommit != nil {
				logrus.Errorf("BalanceRepository -> Deposit -> %v", errCommit)
			}
		}
	}()

	if err != nil {
		return fmt.Errorf("BalanceRepository -> Deposit -> %w", err)
	}
	_, err = r.pool.Exec(ctx, "INSERT INTO balances (balanceID, profileID, amount, tsTime) VALUES($1, $2, $3, $4)", uuid.New(), profileID, amount, time.Now())
	if err != nil {
		return fmt.Errorf("BalanceRepository -> Deposit -> %w", err)
	}
	return nil
}

// GetBalance gets all deposits or withdraws of exact profile from balances table and counts total balance
func (r *BalanceRepository) GetBalance(ctx context.Context, profileID uuid.UUID) (float64, error) {
	tx, err := r.pool.BeginTx(ctx, pgx.TxOptions{IsoLevel: pgx.RepeatableRead})

	defer func() {
		if err != nil {
			errRollback := tx.Rollback(ctx)
			if errRollback != nil {
				logrus.Errorf("GetBalance -> GetBalance -> %v", errRollback)
			}
		} else {
			errCommit := tx.Commit(ctx)
			if errCommit != nil {
				logrus.Errorf("GetBalance -> GetBalance -> %v", errCommit)
			}
		}
	}()

	if err != nil {
		return zero, fmt.Errorf("GetBalance -> GetBalance -> %w", err)
	}

	var totalBalance float64
	var value float64
	rows, err := r.pool.Query(ctx, "SELECT amount FROM balances WHERE profileID = $1 FOR UPDATE", profileID)
	if err != nil {
		return zero, fmt.Errorf("BalanceRepository -> GetBalance -> %w", err)
	}
	defer rows.Close()
	for rows.Next() {
		err = rows.Scan(&value)
		if err != nil {
			return zero, fmt.Errorf("BalanceRepository -> GetBalance -> %w", err)
		}
		totalBalance += value
	}
	return totalBalance, nil
}

// DeleteProfilesBalance deletes all rows from balances table related to exact profile using profileID
func (r *BalanceRepository) DeleteProfilesBalance(ctx context.Context, profileID uuid.UUID) error {
	tx, err := r.pool.Begin(ctx)

	defer func() {
		if err != nil {
			errRollback := tx.Rollback(ctx)
			if errRollback != nil {
				logrus.Errorf("BalanceRepository -> DeleteProfilesBalance -> %v", errRollback)
			}
		} else {
			errCommit := tx.Commit(ctx)
			if errCommit != nil {
				logrus.Errorf("BalanceRepository -> DeleteProfilesBalance -> %v", errCommit)
			}
		}
	}()

	if err != nil {
		return fmt.Errorf("BalanceRepository -> DeleteProfilesBalance -> %w", err)
	}
	res, err := r.pool.Exec(ctx, "DELETE FROM balances WHERE profileID = $1", profileID)
	if err != nil {
		return fmt.Errorf("BalanceRepository -> DeleteProfilesBalance -> %w", err)
	}
	if res.RowsAffected() == 0 {
		return fmt.Errorf("BalanceRepository -> DeleteProfilesBalance -> %w", pgx.ErrNoRows)
	}

	return nil
}
