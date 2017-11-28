package doorman

import (
	"gitlab.com/tokend/go/doorman/data"
	"gitlab.com/tokend/go/doorman/horizon"
)

func NewDoormanHorizon(skipSignatureCheck bool, accountQ data.AccountQ) Doorman {
	return &horizon.Doorman{
		SkipSignatureCheck: skipSignatureCheck,
		AccountQ:           accountQ,
	}
}
