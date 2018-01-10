package corer

import (
	"path"
)

// RequestMaintenance - requests core to perform maintenance
func (c *connector) RequestMaintenance(coreURL string) error {
	u := c.getCoreURL()
	u.Path = path.Join(u.Path, "maintenance")
	q := u.Query()
	q.Set("queue", "true")
	u.RawQuery = q.Encode()
	request := u.String()

	return c.get(request, "Done")
}
