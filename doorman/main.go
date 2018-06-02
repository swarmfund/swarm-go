package doorman

import (
	"net/http"

	"gitlab.com/tokend/go/doorman/data"
)

type SignerConstraint func(*http.Request, Doorman) error

func New(passAllChecks bool, accountQ data.AccountQ) Doorman {
	return &doorman{
		PassAllChecks: passAllChecks,
		AccountQ:      accountQ,
	}
}
