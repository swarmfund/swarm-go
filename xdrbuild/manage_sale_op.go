package xdrbuild

import (
	"github.com/go-ozzo/ozzo-validation"
	"gitlab.com/swarmfund/go/xdr"
)

type ManageSaleOp struct {
	SaleID           uint64
	ManageSaleAction xdr.ManageSaleAction
}

func (op ManageSaleOp) Validate() error {
	return validation.ValidateStruct(&op,
		validation.Field(&op.SaleID, validation.Required),
		validation.Field(&op.ManageSaleAction, validation.Required),
	)
}

func (op ManageSaleOp) XDR() (*xdr.Operation, error) {
	xdrop := &xdr.Operation{
		Body: xdr.OperationBody{
			Type: xdr.OperationTypeManageSale,
			ManageSaleOp: &xdr.ManageSaleOp{
				SaleId: xdr.Uint64(op.SaleID),
				Action: op.ManageSaleAction,
			},
		},
	}
	return xdrop, nil
}
