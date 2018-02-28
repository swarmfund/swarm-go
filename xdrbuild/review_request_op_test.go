package xdrbuild

import (
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
	"gitlab.com/swarmfund/go/xdr"
)

func TestReviewRequestOp_XDR(t *testing.T) {
	hash := [32]byte{
		0xde, 0xad, 0xbe, 0xaf, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
		0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00,
	}

	t.Run("approve", func(t *testing.T) {
		op := ReviewRequestOp{
			ID:     1,
			Hash:   fmt.Sprintf("%x", hash[:]),
			Action: xdr.ReviewRequestOpActionApprove,
		}
		got, err := op.XDR()
		if err != nil {
			t.Fatal(err)
		}
		body := got.Body.ReviewRequestOp
		assert.EqualValues(t, op.ID, body.RequestId)
		assert.EqualValues(t, hash, body.RequestHash)
		assert.EqualValues(t, op.Action, body.Action)
	})

	t.Run("reject", func(t *testing.T) {
		op := ReviewRequestOp{
			ID:     1,
			Hash:   fmt.Sprintf("%x", hash[:]),
			Action: xdr.ReviewRequestOpActionReject,
			Reason: "yoba",
		}
		got, err := op.XDR()
		if err != nil {
			t.Fatal(err)
		}
		body := got.Body.ReviewRequestOp
		assert.EqualValues(t, op.ID, body.RequestId)
		assert.EqualValues(t, hash, body.RequestHash)
		assert.EqualValues(t, op.Action, body.Action)
		assert.EqualValues(t, op.Reason, body.Reason)
	})

	t.Run("approve withdraw", func(t *testing.T) {
		op := ReviewRequestOp{
			ID:     1,
			Hash:   fmt.Sprintf("%x", hash[:]),
			Action: xdr.ReviewRequestOpActionApprove,
		}
		got, err := op.XDR()
		if err != nil {
			t.Fatal(err)
		}
		body := got.Body.ReviewRequestOp
		assert.EqualValues(t, op.ID, body.RequestId)
		assert.EqualValues(t, hash, body.RequestHash)
		assert.EqualValues(t, op.Action, body.Action)
	})

	t.Run("reject withdraw", func(t *testing.T) {
		op := ReviewRequestOp{
			ID:     1,
			Hash:   fmt.Sprintf("%x", hash[:]),
			Action: xdr.ReviewRequestOpActionPermanentReject,
			Reason: "yoba",
			Details: WithdrawalDetails{
				ExternalDetails: "foobar",
			},
		}
		got, err := op.XDR()
		if err != nil {
			t.Fatal(err)
		}
		body := got.Body.ReviewRequestOp
		assert.EqualValues(t, op.ID, body.RequestId)
		assert.EqualValues(t, hash, body.RequestHash)
		assert.EqualValues(t, op.Action, body.Action)
		assert.EqualValues(t, xdr.ReviewableRequestTypeWithdraw, body.RequestDetails.RequestType)
		assert.NotNil(t, body.RequestDetails.Withdrawal)
		assert.EqualValues(t, "foobar", body.RequestDetails.Withdrawal.ExternalDetails)
	})
}
