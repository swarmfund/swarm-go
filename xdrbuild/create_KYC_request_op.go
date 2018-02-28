package xdrbuild

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
	"gitlab.com/swarmfund/go/xdr"
)

type CreateKYCRequestOp struct {
	RequestID      xdr.Uint64
	AccountType    xdr.AccountType
	KYCData        xdr.Longstring
	UpdatedAccount string
	KYCLevel       xdr.Uint32
}

func (op CreateKYCRequestOp) Validate() error {
	return validation.ValidateStruct(&op,
		validation.Field(&op.AccountType, validation.Required),
		validation.Field(&op.KYCData, validation.Required),
		validation.Field(&op.UpdatedAccount, validation.Required),
		validation.Field(&op.KYCLevel, validation.Required),
	)
}

func (op CreateKYCRequestOp) XDR() (*xdr.Operation, error) {
	var updatedAccount xdr.AccountId
	if err := updatedAccount.SetAddress(op.UpdatedAccount); err != nil {
		return nil, errors.Wrap(err, "failed to set updated account")
	}
	xdrop := &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeCreateKycRequest,
			CreateKycRequestOp: &xdr.CreateKycRequestOp{
				RequestId: op.RequestID,
				ChangeKycRequest: xdr.ChangeKycRequest{
					UpdatedAccount:   updatedAccount,
					AccountTypeToSet: op.AccountType,
					KycLevel:         op.KYCLevel,
					KycData:          op.KYCData,
				},
			},
		},
	}
	return xdrop, nil
}
