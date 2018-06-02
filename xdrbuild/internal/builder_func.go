package internal

import "gitlab.com/lbhack/go/xdr"

type OperationFunc func() (*xdr.Operation, error)

func (f OperationFunc) XDR() (*xdr.Operation, error) {
	return f()
}
