package xdrbuild

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/swarmfund/go/xdr"

)

func TestManageSaleOp_XDR(t *testing.T) {
	t.Run("valid", func(t *testing.T) {
		op := ManageSaleOp{
			SaleID:           5,
			ManageSaleAction: xdr.ManageSaleActionBlock,
		}

		assert.NoError(t, op.Validate())
		got, err := op.XDR()
		assert.NoError(t, err)
		body := got.Body.ManageSaleOp
		assert.EqualValues(t, op.SaleID, body.SaleId)
		assert.EqualValues(t, op.ManageSaleAction, body.Action)

	})

	t.Run("missing sale id", func(t *testing.T) {
		op := ManageSaleOp{
			ManageSaleAction: xdr.ManageSaleActionDelete,
		}
		assert.Error(t, op.Validate())
	})

	t.Run("missing action", func(t *testing.T) {
		op := ManageSaleOp{
			SaleID:           2,
		}
		assert.Error(t, op.Validate())
	})


}
