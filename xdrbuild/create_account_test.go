package xdrbuild

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/swarmfund/go/xdr"
	"gitlab.com/tokend/keypair"
)

func TestCreateAccountOp_XDR(t *testing.T) {
	kp, _ := keypair.Random()
	t.Run("valid", func(t *testing.T) {
		op := CreateAccountOp{
			Address:     kp.Address(),
			Recovery:    kp.Address(),
			AccountType: xdr.AccountTypeGeneral,
		}
		assert.NoError(t, op.Validate())
		got, err := op.XDR()
		assert.NoError(t, err)
		body := got.Body.CreateAccountOp
		assert.EqualValues(t, op.AccountType, body.AccountType)
		assert.EqualValues(t, op.Address, body.Destination.Address())
		assert.EqualValues(t, op.Recovery, body.RecoveryKey.Address())
	})

	t.Run("missing address", func(t *testing.T) {
		op := CreateAccountOp{
			Recovery:    kp.Address(),
			AccountType: xdr.AccountTypeGeneral,
		}
		assert.Error(t, op.Validate())
	})

	t.Run("missing recovery", func(t *testing.T) {
		op := CreateAccountOp{
			Address:     kp.Address(),
			AccountType: xdr.AccountTypeGeneral,
		}
		assert.Error(t, op.Validate())
	})

	t.Run("missing account type", func(t *testing.T) {
		op := CreateAccountOp{
			Address:  kp.Address(),
			Recovery: kp.Address(),
		}
		assert.Error(t, op.Validate())
	})
}
