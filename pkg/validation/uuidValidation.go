package validation

import (
	"errors"

	"github.com/google/uuid"
)

func IsValidUUID(id string) bool {
	_, err := uuid.Parse(id)
	return err == nil
}

const (
	InvalidUUIDErrorMessage = "is not a valid uuid"
)

func CreateInvalidError(attribute string) error {
	return errors.New(attribute + " " + InvalidUUIDErrorMessage)
}
