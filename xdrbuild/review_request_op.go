package xdrbuild

import (
	"encoding/hex"

	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdr"
)

type ReviewRequestDetails interface {
	ReviewRequestDetails() xdr.ReviewRequestOpRequestDetails
}

// TODO research why does it exist
type ReviewRequestOpDetails struct {
	Type       xdr.ReviewableRequestType
	Withdrawal *ReviewRequestOpWithdrawalDetails
	Issuance   *ReviewRequestOpIssuanceDetails
}

type ReviewRequestOpWithdrawalDetails struct {
	ExternalDetails string
}

type ReviewRequestOpIssuanceDetails struct{}
type IssuanceDetails struct{}

func (d IssuanceDetails) ReviewRequestDetails() xdr.ReviewRequestOpRequestDetails {
	return xdr.ReviewRequestOpRequestDetails{
		RequestType: xdr.ReviewableRequestTypeIssuanceCreate,
	}
}

type ReviewDetails struct {
	TasksToAdd      uint32
	TasksToRemove   uint32
	ExternalDetails string
}

type ReviewRequestOp struct {
	ID uint64
	// Hash optional, not a pointer for backwards compatibility
	Hash          string
	Action        xdr.ReviewRequestOpAction
	Details       ReviewRequestDetails
	Reason        string
	ReviewDetails *ReviewDetails
}

type WithdrawalDetails struct {
	ExternalDetails string
}

func (d WithdrawalDetails) ReviewRequestDetails() xdr.ReviewRequestOpRequestDetails {
	return xdr.ReviewRequestOpRequestDetails{
		RequestType: xdr.ReviewableRequestTypeWithdraw,
		Withdrawal: &xdr.WithdrawalDetails{
			ExternalDetails: d.ExternalDetails,
		},
	}
}

type TwoStepWithdrawalDetails struct {
	ExternalDetails string
}

func (d TwoStepWithdrawalDetails) ReviewRequestDetails() xdr.ReviewRequestOpRequestDetails {
	return xdr.ReviewRequestOpRequestDetails{
		RequestType: xdr.ReviewableRequestTypeTwoStepWithdrawal,
		TwoStepWithdrawal: &xdr.WithdrawalDetails{
			ExternalDetails: d.ExternalDetails,
		},
	}
}

type UpdateKYCDetails struct {
	TasksToAdd      uint32
	TasksToRemove   uint32
	ExternalDetails string
}

func (d UpdateKYCDetails) ReviewRequestDetails() xdr.ReviewRequestOpRequestDetails {
	return xdr.ReviewRequestOpRequestDetails{
		RequestType: xdr.ReviewableRequestTypeUpdateKyc,
		UpdateKyc: &xdr.UpdateKycDetails{
			TasksToAdd:      xdr.Uint32(d.TasksToAdd),
			TasksToRemove:   xdr.Uint32(d.TasksToRemove),
			ExternalDetails: d.ExternalDetails,
		},
	}
}

func (op ReviewRequestOp) XDR() (*xdr.Operation, error) {
	xdrop := &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeReviewRequest,
			ReviewRequestOp: &xdr.ReviewRequestOp{
				RequestId: xdr.Uint64(op.ID),
				Action:    op.Action,
				Reason:    xdr.Longstring(op.Reason),
			},
		},
	}

	if op.Hash != "" {
		hash, err := hex.DecodeString(op.Hash)
		if err != nil {
			return nil, errors.Wrap(err, "failed to decode hash")
		}
		copy(xdrop.Body.ReviewRequestOp.RequestHash[:], hash[:32])
	}

	if op.Details != nil {
		xdrop.Body.ReviewRequestOp.RequestDetails = op.Details.ReviewRequestDetails()
	}

	if op.ReviewDetails != nil {
		xdrop.Body.ReviewRequestOp.Ext = xdr.ReviewRequestOpExt{
			V: xdr.LedgerVersionAddTasksToReviewableRequest,
			ReviewDetails: &xdr.ReviewDetails{
				TasksToAdd:      xdr.Uint32(op.ReviewDetails.TasksToAdd),
				TasksToRemove:   xdr.Uint32(op.ReviewDetails.TasksToRemove),
				ExternalDetails: op.ReviewDetails.ExternalDetails,
			},
		}
	}

	return xdrop, nil
}
