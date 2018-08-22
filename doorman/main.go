package doorman

import (
	"net/http"

	"gitlab.com/tokend/go/doorman/data"
)

type SignerExtension int

const (
	SignerExtensionUsersList SignerExtension = 1 << iota
	SignerExternsionUser
	SignerExternsionOperationsList
	SignerExternsionPendingIssuance
	SignerExternsionIssuanceHistory
	SignerExternsionPendingKYC
	SignerExternsionKYCHistory
	SignerExternsionCrowdfundingCampaign
)

type SignerConstraint func(*http.Request, Doorman) error

func New(passAllChecks bool, accountQ data.AccountQ) Doorman {
	return &doorman{
		PassAllChecks: passAllChecks,
		AccountQ:      accountQ,
	}
}
