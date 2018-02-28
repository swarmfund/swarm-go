package xdrbuild

import (
	"testing"
	"gitlab.com/swarmfund/go/keypair"
	"gitlab.com/swarmfund/go/xdr"
	"github.com/stretchr/testify/assert"
)

func TestCreateKYCRequestOp_XDR(t *testing.T) {
	kp, _ := keypair.Random()
	t.Run("valid", func(t *testing.T) {
		op := CreateKYCRequestOp{
			RequestID:      0,
			AccountType:    xdr.AccountTypeGeneral,
			KYCData:        "Some KYC data",
			UpdatedAccount: kp.Address(),
			KYCLevel:       1,
		}
		assert.NoError(t, op.Validate())
		got, err := op.XDR()
		assert.NoError(t, err)
		body := got.Body.CreateKycRequestOp
		assert.EqualValues(t, op.RequestID, body.RequestId)
		assert.EqualValues(t, op.AccountType, body.ChangeKycRequest.AccountTypeToSet)
		assert.EqualValues(t, op.KYCData, body.ChangeKycRequest.KycData)
		assert.EqualValues(t, op.UpdatedAccount, body.ChangeKycRequest.UpdatedAccount.Address())
		assert.EqualValues(t, op.KYCLevel, body.ChangeKycRequest.KycLevel)
	})

	t.Run("missing account type", func(t *testing.T) {
		op := CreateKYCRequestOp{
			RequestID:      0,
			KYCData:        "Some KYC data",
			UpdatedAccount: kp.Address(),
			KYCLevel:       1,
		}
		assert.Error(t, op.Validate())
	})

	t.Run("missing KYC data", func(t *testing.T) {
		op := CreateKYCRequestOp{
			RequestID:      0,
			AccountType:    xdr.AccountTypeGeneral,
			UpdatedAccount: kp.Address(),
			KYCLevel:       1,
		}
		assert.Error(t, op.Validate())
	})

	t.Run("missing updated account", func(t *testing.T) {
		op := CreateKYCRequestOp{
			RequestID:      0,
			AccountType:    xdr.AccountTypeGeneral,
			KYCData:        "Some KYC data",
			KYCLevel:       1,
		}
		assert.Error(t, op.Validate())
	})
}
