package doorman

import (
	"net/http"

	"gitlab.com/swarmfund/go/doorman/types"
	"gitlab.com/swarmfund/go/signcontrol"
)

// Check ensures request passes at least one constraint
func Check(r *http.Request, constraints ...types.SignerConstraint) error {
	for _, constraint := range constraints {
		switch err := constraint(r); err {
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
