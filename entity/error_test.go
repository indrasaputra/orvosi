package entity_test

import (
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/indrasaputra/orvosi/entity"
)

func TestErrInternal(t *testing.T) {
	t.Run("success get internal error", func(t *testing.T) {
		err := entity.ErrInternal("")

		assert.Contains(t, err.Error(), "rpc error: code = Internal")
	})
}

func TestErrEmptyMedicalRecord(t *testing.T) {
	t.Run("success get empty medical record error", func(t *testing.T) {
		err := entity.ErrEmptyMedicalRecord()

		assert.Contains(t, err.Error(), "rpc error: code = InvalidArgument")
	})
}
