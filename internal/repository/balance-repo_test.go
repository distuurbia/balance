package repository

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestAddBalanceChange(t *testing.T) {
	err := r.AddBalanceChange(context.Background(), uuid.New(), 102.3)
	require.NoError(t, err)
}

func TestAddBalanceChangeNilProfileID(t *testing.T) {
	err := r.AddBalanceChange(context.Background(), uuid.Nil, 102.3)
	require.Error(t, err)
}

func TestAddBalanceChangeZeroAmount(t *testing.T) {
	err := r.AddBalanceChange(context.Background(), uuid.New(), 0)
	require.Error(t, err)
}

func TestGetBalance(t *testing.T) {
	dep := 102.3
	withdraw := -56.7
	profileID := uuid.New()
	err := r.AddBalanceChange(context.Background(), profileID, dep)
	require.NoError(t, err)
	err = r.AddBalanceChange(context.Background(), profileID, withdraw)
	require.NoError(t, err)
	totalBalance, err := r.GetBalance(context.Background(), profileID)
	require.NoError(t, err)
	require.Equal(t, dep+withdraw, totalBalance)
}

func TestGetBalanceNilID(t *testing.T) {
	_, err := r.GetBalance(context.Background(), uuid.Nil)
	require.Error(t, err)
}

func TestDeleteProfilesBalance(t *testing.T) {
	dep := 102.3
	withdraw := -56.7
	profileID := uuid.New()

	err := r.AddBalanceChange(context.Background(), profileID, dep)
	require.NoError(t, err)
	err = r.AddBalanceChange(context.Background(), profileID, withdraw)
	require.NoError(t, err)

	err = r.DeleteProfilesBalance(context.Background(), profileID)
	require.NoError(t, err)

	totalBalance, err := r.GetBalance(context.Background(), profileID)
	require.NoError(t, err)
	require.Equal(t, float64(0), totalBalance)
}

func TestDeleteProfilesBalanceNilID(t *testing.T) {
	err := r.DeleteProfilesBalance(context.Background(), uuid.Nil)
	require.Error(t, err)
}

func TestDeleteProfilesBalanceNotExist(t *testing.T) {
	err := r.DeleteProfilesBalance(context.Background(), uuid.New())
	require.Error(t, err)
}
