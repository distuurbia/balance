// Package handler contains methods that handle requests and send them to service part
package handler

import (
	"context"
	"fmt"

	protocol "github.com/distuurbia/balance/protocol/balance"
	"github.com/go-playground/validator"
	"github.com/google/uuid"
	"github.com/sirupsen/logrus"
)

// BalanceService is an interface that contains methods of service part
type BalanceService interface {
	AddBalanceChange(ctx context.Context, profileID uuid.UUID, amount float64) error
	GetBalance(ctx context.Context, profileID uuid.UUID) (float64, error)
	DeleteProfilesBalance(ctx context.Context, profileID uuid.UUID) error
}

// BalanceHandler is a structure of handler that contains an object implemented BalanceService interface and validator
type BalanceHandler struct {
	s        BalanceService
	validate *validator.Validate
	protocol.UnimplementedBalanceServiceServer
}

// NewBalanceHandler creates an onject of *BalanceHandler fulfilled with provided fields
func NewBalanceHandler(s BalanceService, validate *validator.Validate) *BalanceHandler {
	return &BalanceHandler{s: s, validate: validate}
}

// ValidationID validate given in and parses it to uuid.UUID type
func (h *BalanceHandler) ValidationID(ctx context.Context, id string) (uuid.UUID, error) {
	err := h.validate.VarCtx(ctx, id, "required,required,uuid")
	if err != nil {
		logrus.Errorf("ValidationID -> %v", err)
		return uuid.Nil, err
	}

	profileID, err := uuid.Parse(id)
	if err != nil {
		logrus.Errorf("ValidationID -> %v", err)
		return uuid.Nil, err
	}

	if profileID == uuid.Nil {
		logrus.Errorf("ValidationID -> error: failed to use uuid")
		return uuid.Nil, fmt.Errorf("ValidationID -> error: failed to use uuid")
	}
	return profileID, nil
}

// AddBalanceChange calls lower method of service AddBalanceChange
func (h *BalanceHandler) AddBalanceChange(ctx context.Context, req *protocol.AddBalanceChangeRequest) (*protocol.AddBalanceChangeResponse, error) {
	err := h.validate.VarCtx(ctx, req.Amount, "required,ne=0")
	if err != nil {
		logrus.Errorf("ProfileHandler -> AddBalanceChange -> %v", err)
		return &protocol.AddBalanceChangeResponse{}, err
	}

	profileID, err := h.ValidationID(ctx, req.ProfileID)
	if err != nil {
		logrus.Errorf("BalanceHandler -> AddBalanceChange -> %v", err)
		return &protocol.AddBalanceChangeResponse{}, err
	}

	err = h.s.AddBalanceChange(ctx, profileID, req.Amount)
	if err != nil {
		logrus.Errorf("ProfileHandler -> AddBalanceChange -> %v", err)
		return &protocol.AddBalanceChangeResponse{}, err
	}

	return &protocol.AddBalanceChangeResponse{}, nil
}

// GetBalance calls lower method of service GetBalance
func (h *BalanceHandler) GetBalance(ctx context.Context, req *protocol.GetBalanceRequest) (*protocol.GetBalanceResponse, error) {
	profileID, err := h.ValidationID(ctx, req.ProfileID)
	if err != nil {
		logrus.Errorf("BalanceHandler -> GetBalance -> %v", err)
		return &protocol.GetBalanceResponse{}, err
	}

	totalBalance, err := h.s.GetBalance(ctx, profileID)
	if err != nil {
		logrus.Errorf("BalanceHandler -> GetBalance -> %v", err)
		return &protocol.GetBalanceResponse{}, err
	}

	return &protocol.GetBalanceResponse{TotalBalance: totalBalance}, nil
}

// DeleteProfilesBalance calls lower method of service DeleteProfilesBalance
func (h *BalanceHandler) DeleteProfilesBalance(ctx context.Context, req *protocol.DeleteProfilesBalanceRequest) (
	*protocol.DeleteProfilesBalanceResponse, error) {
	profileID, err := h.ValidationID(ctx, req.ProfileID)
	if err != nil {
		logrus.Errorf("BalanceHandler -> DeleteProfilesBalance -> %v", err)
		return &protocol.DeleteProfilesBalanceResponse{}, err
	}

	err = h.s.DeleteProfilesBalance(ctx, profileID)
	if err != nil {
		logrus.Errorf("BalanceHandler -> DeleteProfilesBalance -> %v", err)
		return &protocol.DeleteProfilesBalanceResponse{}, err
	}

	return &protocol.DeleteProfilesBalanceResponse{}, nil
}
