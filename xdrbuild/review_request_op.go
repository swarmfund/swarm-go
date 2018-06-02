package xdrbuild

import (
	"encoding/hex"

	"github.com/pkg/errors"
	"gitlab.com/tokend/go/xdr"
)

type ReviewRequestDetails interface {
	ReviewRequestDetails() xdr.ReviewRequestOpRequestDetails
}

type ReviewRequestOpDetails struct {
	Type       xdr.ReviewableRequestType
	Withdrawal *ReviewRequestOpWithdrawalDetails
}

type ReviewRequestOpWithdrawalDetails struct {
	ExternalDetails string
}

type ReviewRequestOp struct {
	ID      uint64
	Hash    string
	Action  xdr.ReviewRequestOpAction
	Details ReviewRequestDetails
	Reason  string
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
	hash, err := hex.DecodeString(op.Hash)
	if err != nil {
		return nil, errors.Wrap(err, "failed to decode hash")
	}
	var xdrhash xdr.Hash
	copy(xdrhash[:], hash[:32])

	var details xdr.ReviewRequestOpRequestDetails
	if op.Details != nil {
		details = op.Details.ReviewRequestDetails()
	}

	xdrop := &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeReviewRequest,
			ReviewRequestOp: &xdr.ReviewRequestOp{
				RequestId:      xdr.Uint64(op.ID),
				RequestHash:    xdrhash,
				Action:         op.Action,
				RequestDetails: details,
				Reason:         xdr.String256(op.Reason),
			},
		},
	}

	return xdrop, nil
}
