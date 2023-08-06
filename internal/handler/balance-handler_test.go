package handler

import (
	"context"
	"testing"

	"github.com/distuurbia/balance/internal/handler/mocks"
	protocol "github.com/distuurbia/balance/protocol/balance"
	"github.com/google/uuid"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

func TestValidationID(t *testing.T) {
	s := new(mocks.BalanceService)
	h := NewBalanceHandler(s, validate)

	testID := uuid.New()
	validatedID, err := h.ValidationID(context.Background(), testID.String())
	require.NoError(t, err)
	require.Equal(t, testID, validatedID)
}

func TestAddBalanceChange(t *testing.T) {
	s := new(mocks.BalanceService)

	s.On("AddBalanceChange", mock.Anything, mock.AnythingOfType("uuid.UUID"), mock.AnythingOfType("float64")).Return(nil)

	h := NewBalanceHandler(s, validate)

	_, err := h.AddBalanceChange(context.Background(), &protocol.AddBalanceChangeRequest{
		ProfileID: uuid.New().String(),
		Amount:    float64(505.23),
	})
	require.NoError(t, err)
}

func TestGetBalance(t *testing.T) {
	s := new(mocks.BalanceService)

	s.On("GetBalance", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(float64(505), nil)

	h := NewBalanceHandler(s, validate)

	resp, err := h.GetBalance(context.Background(), &protocol.GetBalanceRequest{
		ProfileID: uuid.New().String(),
	})
	require.NoError(t, err)
	require.Equal(t, resp.TotalBalance, float64(505))
}

func TestDeleteProfilesBalance(t *testing.T) {
	s := new(mocks.BalanceService)

	s.On("DeleteProfilesBalance", mock.Anything, mock.AnythingOfType("uuid.UUID")).Return(nil)

	h := NewBalanceHandler(s, validate)

	_, err := h.DeleteProfilesBalance(context.Background(), &protocol.DeleteProfilesBalanceRequest{
		ProfileID: uuid.New().String(),
	})
	require.NoError(t, err)
}
