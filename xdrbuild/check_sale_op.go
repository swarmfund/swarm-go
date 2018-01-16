package xdrbuild

import "gitlab.com/swarmfund/go/xdr"

type CheckSaleOp struct {
}

func (op CheckSaleOp) XDR() (*xdr.Operation, error) {
	return &xdr.Operation{
		Body: xdr.OperationBody{
			Type:             xdr.OperationTypeCheckSaleState,
			CheckSaleStateOp: &xdr.CheckSaleStateOp{},
		},
	}, nil
}
