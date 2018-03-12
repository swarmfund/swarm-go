package xdrbuild

import (
	"github.com/go-ozzo/ozzo-validation"
	"github.com/pkg/errors"
	"gitlab.com/swarmfund/go/xdr"
)

type CreateAccountOp struct {
	Address     string
	Recovery    string
	AccountType xdr.AccountType
	Referrer    *string
}

func (op CreateAccountOp) Validate() error {
	return validation.ValidateStruct(&op,
		// TODO validate address, recovery and referrer are addresses
		validation.Field(&op.Address, validation.Required),
		validation.Field(&op.Recovery, validation.Required),
		validation.Field(&op.AccountType, validation.Required),
	)
}

func (op CreateAccountOp) XDR() (*xdr.Operation, error) {
	var destination, recovery xdr.AccountId
	if err := destination.SetAddress(op.Address); err != nil {
		return nil, errors.Wrap(err, "failed to set destination")
	}
	if err := recovery.SetAddress(op.Recovery); err != nil {
		return nil, errors.Wrap(err, "failed to set recovery")
	}
	xdrop := &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeCreateAccount,
			CreateAccountOp: &xdr.CreateAccountOp{
				Destination: destination,
				AccountType: op.AccountType,
				RecoveryKey: recovery,
			},
		},
	}

	if op.Referrer != nil {
		var referrer xdr.AccountId
		referrer.SetAddress(*op.Referrer)
		xdrop.Body.CreateAccountOp.Referrer = &referrer
	}

	return xdrop, nil
}
