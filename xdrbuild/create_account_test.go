package xdrbuild

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/tokend/go/xdr"
	"gitlab.com/tokend/keypair"
)

func TestCreateAccountOp_XDR(t *testing.T) {
	one, _ := keypair.Random()
	two, _ := keypair.Random()
	three, _ := keypair.Random()

	t.Run("valid w/o referrer", func(t *testing.T) {
		op := CreateAccountOp{
			Address:     one.Address(),
			Recovery:    two.Address(),
			AccountType: xdr.AccountTypeGeneral,
		}
		assert.NoError(t, op.Validate())
		got, err := op.XDR()
		assert.NoError(t, err)
		body := got.Body.CreateAccountOp
		assert.EqualValues(t, op.AccountType, body.AccountType)
		assert.EqualValues(t, op.Address, body.Destination.Address())
		assert.EqualValues(t, op.Recovery, body.RecoveryKey.Address())
		assert.Nil(t, body.Referrer)
	})

	t.Run("valid with referrer", func(t *testing.T) {
		referrer := three.Address()
		op := CreateAccountOp{
			Address:     one.Address(),
			Recovery:    two.Address(),
			AccountType: xdr.AccountTypeGeneral,
			Referrer:    &referrer,
		}
		assert.NoError(t, op.Validate())
		got, err := op.XDR()
		assert.NoError(t, err)
		body := got.Body.CreateAccountOp
		assert.EqualValues(t, op.AccountType, body.AccountType)
		assert.EqualValues(t, op.Address, body.Destination.Address())
		assert.EqualValues(t, op.Recovery, body.RecoveryKey.Address())
		assert.EqualValues(t, &referrer, op.Referrer)
	})

	t.Run("missing address", func(t *testing.T) {
		op := CreateAccountOp{
			Recovery:    one.Address(),
			AccountType: xdr.AccountTypeGeneral,
		}
		assert.Error(t, op.Validate())
	})

	t.Run("missing recovery", func(t *testing.T) {
		op := CreateAccountOp{
			Address:     one.Address(),
			AccountType: xdr.AccountTypeGeneral,
		}
		assert.Error(t, op.Validate())
	})

	t.Run("missing account type", func(t *testing.T) {
		op := CreateAccountOp{
			Address:  one.Address(),
			Recovery: two.Address(),
		}
		assert.Error(t, op.Validate())
	})
}
