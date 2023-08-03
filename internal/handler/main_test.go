package handler

import (
	"os"
	"testing"

	"github.com/go-playground/validator"
)

var (
	validate *validator.Validate
)

func TestMain(m *testing.M) {
	validate = validator.New()
	exitCode := m.Run()
	os.Exit(exitCode)
}
