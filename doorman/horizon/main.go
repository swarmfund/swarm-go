package horizon

import (
	"net/http"

	"gitlab.com/tokend/go/doorman/data"
	"gitlab.com/tokend/go/doorman/types"
	"gitlab.com/tokend/go/signcontrol"
)

type Doorman struct {
	// SkipSignatureCheck disable signature validation
	SkipSignatureCheck bool
	// PassAllChecks disable constraints validation completely, any request will succeed
	PassAllChecks bool
	// AccountQ used to get account details during constraint checks
	AccountQ data.AccountQ
	// MasterAddress master account address used for admin signature checks
	MasterAddress string
}

func (doorman *Doorman) MasterSigner(_ string) types.SignerConstraint {
	return doorman.SignerOf(doorman.MasterAddress)
}

func (doorman *Doorman) SignerOf(address string) types.SignerConstraint {
	return func(r *http.Request) error {
		if doorman.PassAllChecks {
			return nil
		}

		var signer string
		var err error

		if doorman.SkipSignatureCheck {
			if signer = r.Header.Get(signcontrol.PublicKeyHeader); signer == "" {
				return signcontrol.ErrNotSigned
			}
		} else {
			signer, err = signcontrol.CheckSignature(r)
			if err != nil {
				return err
			}
		}

		if signer == address {
			return nil
		}

		signers, err := doorman.AccountQ.Signers(address)
		if err != nil {
			return err
		}

		// TODO make it readable
		for _, accountSigner := range signers {
			if accountSigner.AccountID == signer && accountSigner.Weight > 0 {
				return nil
			}
		}
		return signcontrol.ErrNotAllowed
	}
}
