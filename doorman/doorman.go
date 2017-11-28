package doorman

import "gitlab.com/tokend/go/doorman/types"

type Doorman interface {
	MasterSigner(address string) types.SignerConstraint
	SignerOf(address string) types.SignerConstraint
}
