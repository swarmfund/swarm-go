package doorman

import (
	"net/http"

	"strconv"
	"strings"

	"gitlab.com/tokend/go/resources"
	"gitlab.com/tokend/go/signcontrol"
)

func SignerOf(address string) SignerConstraint {
	return func(r *http.Request, doorman Doorman) error {
		signer, err := signcontrol.CheckSignature(r)
		if err != nil {
			return err
		}

		if signer == address {
			return nil
		}
		signers, err := doorman.AccountSigners(address)
		if err != nil {
			return err
		}

		for _, accountSigner := range signers {
			if accountSigner.AccountID == signer && accountSigner.Weight > 0 {
				return nil
			}
		}
		return signcontrol.ErrNotAllowed
	}
}

func SignerOfWithPermission(address string, signerPermission SignerExtension) SignerConstraint {
	return func(r *http.Request, doorman Doorman) error {
		signer, err := signcontrol.CheckSignature(r)
		if err != nil {
			return err
		}

		if signer == address {
			return nil
		}
		signers, err := doorman.AccountSigners(address)
		if err != nil {
			return err
		}

		for _, accountSigner := range signers {
			if accountSigner.AccountID == signer && accountSigner.Weight > 0 && checkPermission(accountSigner, int(signerPermission)) {
				return nil
			}
		}
		return signcontrol.ErrNotAllowed
	}
}

func SignatureOf(address string) SignerConstraint {
	return func(r *http.Request, doorman Doorman) error {
		signer, err := signcontrol.CheckSignature(r)
		if err != nil {
			return err
		}

		if signer == address {
			return nil
		}

		return signcontrol.ErrNotAllowed
	}
}

func checkPermission(accountSigner resources.Signer, permission int) bool {
	s := strings.Split(accountSigner.Name, ":")
	prms := s[len(s)-1]
	signerPermission, err := strconv.Atoi(prms)
	if err != nil {
		return true
	}
	return signerPermission&permission != 0
}
