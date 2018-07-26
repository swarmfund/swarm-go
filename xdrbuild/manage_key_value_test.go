package xdrbuild

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestManageKeyValueOp_XDR(t *testing.T) {
	t.Run("valid with int", func(t *testing.T) {
		v := uint32(6)
		op := ManageKeyValueOp{
			Key:    "Key",
			Uint32: &v,
		}
		assert.NoError(t, op.Validate())
		_, err := op.XDR()
		assert.NoError(t, err)
	})
	t.Run("valid with string", func(t *testing.T) {
		str := "TaskFaceValidation"
		op := ManageKeyValueOp{
			Key:    "Key",
			String: &str,
		}
		assert.NoError(t, op.Validate())
		_, err := op.XDR()
		assert.NoError(t, err)

	})
	t.Run("invalid struct", func(t *testing.T) {

		op := ManageKeyValueOp{}
		assert.Error(t, op.Validate())
	})
}
