package corer

import (
	"gitlab.com/distributed_lab/logan/v3/errors"
	"path"
)

const (
	// TxStatusError - core has rejected tx
	TxStatusError = "ERROR"
	// TxStatusPending - core still processing tx
	TxStatusPending = "PENDING"
	// TxStatusDuplicate - core already processing tx with same hash
	TxStatusDuplicate = "DUPLICATE"
)

// CoreSubmissionResponse is the json response from stellar-core's tx endpoint
type CoreSubmissionResponse struct {
	Exception string `json:"exception"`
	Error     string `json:"error"`
	Status    string `json:"status"`
}

// SubmitTx - submits tx to core
func (c *connector) SubmitTx(env string) (*CoreSubmissionResponse, error) {
	u := c.getCoreURL()
	u.Path = path.Join(u.Path, "tx")
	q := u.Query()
	q.Add("blob", env)
	u.RawQuery = q.Encode()

	var result CoreSubmissionResponse
	err := c.getJSON(u.String(), &result)
	if err != nil {
		return nil, errors.Wrap(err, "Failed to get json")
	}

	return &result, nil
}
