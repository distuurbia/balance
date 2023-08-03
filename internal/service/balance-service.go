// Package service contains the bisnes logic of app
package service

import (
	"context"
	"fmt"

	"github.com/google/uuid"
)

// BalanceRepository is an interface that contains methods of repository part
type BalanceRepository interface {
	AddBalanceChange(ctx context.Context, profileID uuid.UUID, amount float64) error
	GetBalance(ctx context.Context, profileID uuid.UUID) (float64, error)
	DeleteProfilesBalance(ctx context.Context, profileID uuid.UUID) error
}

// BalanceService contains an object that gonna implement BalanceRepository interface
type BalanceService struct {
	r BalanceRepository
}

// zero contains zero value
const zero = 0

// NewBalanceService creates *BalanceService object, filles it and returns
func NewBalanceService(r BalanceRepository) *BalanceService {
	return &BalanceService{r: r}
}

// AddBalanceChange calls lower method of repository AddBalanceChange
func (s *BalanceService) AddBalanceChange(ctx context.Context, profileID uuid.UUID, amount float64) error {
	if profileID == uuid.Nil {
		return fmt.Errorf("BalanceService -> AddBalanceChange -> error: failed to use uuid")
	}
	if amount == zero {
		return fmt.Errorf("BalanceService -> AddBalanceChange -> error: amount cannot be zero")
	}
	err := s.r.AddBalanceChange(ctx, profileID, amount)
	if err != nil {
		return fmt.Errorf("BalanceService -> AddBalanceChange -> %w", err)
	}
	return nil
}

// GetBalance calls lower method of reposirory GetBalance
func (s *BalanceService) GetBalance(ctx context.Context, profileID uuid.UUID) (float64, error) {
	if profileID == uuid.Nil {
		return zero, fmt.Errorf("BalanceService -> GetBalance -> error: failed to use uuid")
	}
	totalBalance, err := s.r.GetBalance(ctx, profileID)
	if err != nil {
		return zero, fmt.Errorf("BalanceService -> GetBalance -> %w", err)
	}
	return totalBalance, nil
}

// DeleteProfilesBalance calls lower method of repository DeleteProfilesBalance
func (s *BalanceService) DeleteProfilesBalance(ctx context.Context, profileID uuid.UUID) error {
	if profileID == uuid.Nil {
		return fmt.Errorf("BalanceService -> DeleteProfilesBalance -> error: failed to use uuid")
	}
	err := s.r.DeleteProfilesBalance(ctx, profileID)
	if err != nil {
		return fmt.Errorf("BalanceService -> DeleteProfilesBalance -> %w", err)
	}
	return nil
}
