package doorman

import (
	"gitlab.com/swarmfund/go/doorman/data"
	"gitlab.com/swarmfund/go/doorman/horizon"
)

func NewDoormanHorizon(skipSignatureCheck bool, accountQ data.AccountQ) Doorman {
	return &horizon.Doorman{
		SkipSignatureCheck: skipSignatureCheck,
		AccountQ:           accountQ,
	}
}
