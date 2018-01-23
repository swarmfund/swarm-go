package corer

import (
	"gitlab.com/distributed_lab/logan/v3/errors"
	"path"
)

// GetCoreInfo - returns error if failed to get info
func (c *connector) GetCoreInfo(infoResponse interface{}) error {
	coreURL := c.getCoreURL()
	coreURL.Path = path.Join(coreURL.Path, "info")
	err := c.getJSON(coreURL.String(), infoResponse)
	if err != nil {
		return errors.Wrap(err, "Failed to get json")
	}

	return nil
}
