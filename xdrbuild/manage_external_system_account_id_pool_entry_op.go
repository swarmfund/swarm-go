package xdrbuild

import (
	"gitlab.com/swarmfund/go/xdr"
	"github.com/go-ozzo/ozzo-validation"
)

type CreateExternalSystemAccountIdPoolEntryOp struct {
	ExternalSystemType xdr.ExternalSystemType
	Data               xdr.Longstring
}

func (c CreateExternalSystemAccountIdPoolEntryOp) Validate() error {
	return validation.ValidateStruct(&c,
		validation.Field(&c.ExternalSystemType, validation.Required),
		validation.Field(&c.Data, validation.Required),
	)
}

func (c CreateExternalSystemAccountIdPoolEntryOp) XDR() (*xdr.Operation, error) {
	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageExternalSystemAccountIdPoolEntry,
			ManageExternalSystemAccountIdPoolEntryOp: &xdr.ManageExternalSystemAccountIdPoolEntryOp{
				ActionInput: xdr.ManageExternalSystemAccountIdPoolEntryOpActionInput{
					Action: xdr.ManageExternalSystemAccountIdPoolEntryActionCreate,
					CreateExternalSystemAccountIdPoolEntryActionInput: &xdr.CreateExternalSystemAccountIdPoolEntryActionInput{
						ExternalSystemType: c.ExternalSystemType,
						Data:               c.Data,
					},
				},
			},
		},
	}, nil
}
