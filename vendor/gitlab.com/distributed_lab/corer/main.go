package corer

import (
	"gitlab.com/distributed_lab/logan/v3/errors"
	"net/http"
	"net/url"
)

// Connector -- provides method to communicate with core
type Connector interface {
	// CloseLedger - requests core to close ledger at specified time
	CloseLedger(closeTime int64) error
	// GetCoreInfo - returns error if failed to get info
	GetCoreInfo(infoResponse interface{}) error
	// RequestMaintenance - requests core to perform maintenance
	RequestMaintenance(coreURL string) error
	// SetCursor reports to core last ledger been processed, so that core could release resources during maintenance
	SetCursor(id string, lastLedger int32) error
	// SubmitTx - submits tx to core
	SubmitTx(env string) (*CoreSubmissionResponse, error)
}

type connector struct {
	Client *http.Client

	coreURL url.URL
}

// NewConnector creates instance of Connector
func NewConnector(httpClient *http.Client, coreURL string) (Connector, error) {
	parsedCoreURL, err := url.Parse(coreURL)
	if err != nil {
		return nil, errors.Wrap(err, "Invalid core URL")
	}

	return &connector{
		Client:  httpClient,
		coreURL: *parsedCoreURL,
	}, nil
}

func (c *connector) getCoreURL() url.URL {
	return c.coreURL
}
