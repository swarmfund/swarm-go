package corer

import (
	"gitlab.com/distributed_lab/logan/v3/errors"
	"path"
	"strconv"
)

// CloseLedger - requests core to close ledger at specified time
func (c *connector) CloseLedger(closeTime int64) error {
	u := c.getCoreURL()
	u.Path = path.Join(u.Path, "manualclose")
	q := u.Query()
	q.Add("close_time", strconv.FormatInt(closeTime, 10))
	u.RawQuery = q.Encode()

	err := c.get(u.String(), "Forcing ledger to close...")
	if err != nil {
		return errors.Wrap(err, "Failed to perform request")
	}

	return nil
}
