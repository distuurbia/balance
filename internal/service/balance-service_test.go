package service

import (
	"context"
	"testing"

	"github.com/distuurbia/balance/internal/service/mocks"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestAddBalanceChange(t *testing.T) {
	r := new(mocks.BalanceRepository)
	r.On("AddBalanceChange", mock.Anything, mock.AnythingOfType("uuid.UUID"), mock.AnythingOfType("float64")).Return(nil)
	s := NewBalanceService(r)

	err := s.AddBalanceChange(context.Background(), uuid.New(), 102.3)
	require.NoError(t, err)
}

func TestGetBalance(t *testing.T) {
	r := new(mocks.BalanceRepository)
	r.On("GetBalance", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(float64(2.1), nil)
	s := NewBalanceService(r)

	profileID := uuid.New()

	totalBalance, err := s.GetBalance(context.Background(), profileID)
	require.NoError(t, err)
	require.Equal(t, float64(2.1), totalBalance)
}

func TestDeleteProfilesBalance(t *testing.T) {
	r := new(mocks.BalanceRepository)
	r.On("DeleteProfilesBalance", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(nil)
	s := NewBalanceService(r)

	profileID := uuid.New()

	err := s.DeleteProfilesBalance(context.Background(), profileID)
	require.NoError(t, err)
}
