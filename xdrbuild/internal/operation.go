package internal

import "gitlab.com/swarmfund/go/xdr"

type Operation interface {
	XDR() (*xdr.Operation, error)
}
