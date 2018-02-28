package xdrbuild

import (
	"testing"
	"gitlab.com/swarmfund/go/xdr"
	"github.com/stretchr/testify/assert"
)

func TestCreateExternalSystemAccountIdPoolEntryOp_XDR(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		op := CreateExternalSystemAccountIdPoolEntryOp{
			ExternalSystemType: xdr.ExternalSystemTypeErc20Token,
			Data: "Some data",
		}
		assert.NoError(t, op.Validate())
		got, err := op.XDR()
		assert.NoError(t, err)
		body := got.Body.ManageExternalSystemAccountIdPoolEntryOp
		assert.EqualValues(t, op.ExternalSystemType, body.ActionInput.CreateExternalSystemAccountIdPoolEntryActionInput.ExternalSystemType)
		assert.EqualValues(t, op.Data, body.ActionInput.CreateExternalSystemAccountIdPoolEntryActionInput.Data)
	})

	t.Run("invalid data", func(t *testing.T) {
		op := CreateExternalSystemAccountIdPoolEntryOp{
			ExternalSystemType: xdr.ExternalSystemTypeErc20Token,
			Data: "",
		}
		assert.Error(t, op.Validate())
	})
}

