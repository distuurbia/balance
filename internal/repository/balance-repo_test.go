package repository

import (
	"context"
	"testing"

	"github.com/google/uuid"
	"github.com/stretchr/testify/require"
)

func TestAddBalanceChange(t *testing.T) {
	profileID := uuid.New()
	amount := float64(102.3)
	err := r.AddBalanceChange(context.Background(), profileID, amount)
	require.NoError(t, err)

	var testAmount float64
	err = r.pool.QueryRow(context.Background(), "SELECT amount FROM balances WHERE profileid = $1", profileID).Scan(&testAmount)
	require.NoError(t, err)
	require.Equal(t, amount, testAmount)
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
