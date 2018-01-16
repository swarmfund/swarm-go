package xdrbuild

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"gitlab.com/swarmfund/go/xdr"
)

func TestCheckSaleOp_XDR(t *testing.T) {
	op := CheckSaleOp{}
	got, err := op.XDR()
	assert.NoError(t, err)
	assert.Equal(t, xdr.OperationTypeCheckSaleState, got.Body.Type)
}
