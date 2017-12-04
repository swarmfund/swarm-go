package doorman

import (
	"net/http"

	"gitlab.com/swarmfund/go/doorman/data"
	"gitlab.com/swarmfund/go/resources"
	"gitlab.com/swarmfund/go/signcontrol"
)

// Doorman interface only purpose is to simplify tests
type Doorman interface {
	Check(*http.Request, ...SignerConstraint) error
	AccountSigners(string) ([]resources.Signer, error)
}

type doorman struct {
	// PassAllChecks disable constraints validation completely, any request will succeed
	PassAllChecks bool
	// AccountQ used to get account details during constraint checks
	AccountQ data.AccountQ
}

func (d *doorman) AccountSigners(address string) ([]resources.Signer, error) {
	return d.AccountQ.Signers(address)
}

// Check ensures request passes at least one constraint
func (d *doorman) Check(r *http.Request, constraints ...SignerConstraint) error {
	if d.PassAllChecks {
		return nil
	}

	for _, constraint := range constraints {
		switch err := constraint(r, d); err {
		case nil:
			// request passed constraint check
			return nil
		case signcontrol.ErrNotAllowed:
			// check failed, let's try next one
			continue
		default:
			// probably runtime issue
			return err
		}
	}

	// request failed all checks
	return signcontrol.ErrNotAllowed
}
