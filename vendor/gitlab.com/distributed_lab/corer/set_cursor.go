package corer

import (
	"fmt"
	"path"
)

// SetCursor reports to core last ledger been processed, so that core could release resources during maintenance
func (c *connector) SetCursor(id string, lastLedger int32) error {
	u := c.getCoreURL()
	u.Path = path.Join(u.Path, "setcursor")
	q := u.Query()
	q.Set("id", id)
	q.Set("cursor", fmt.Sprintf("%d", lastLedger))
	u.RawQuery = q.Encode()

	return c.get(u.String(), "Done")
}
