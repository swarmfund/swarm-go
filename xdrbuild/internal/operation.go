package internal

import "gitlab.com/lbhack/go/xdr"

type Operation interface {
	XDR() (*xdr.Operation, error)
}
