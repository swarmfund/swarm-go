package corer

import (
	"fmt"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"io/ioutil"
	"strings"
)

func (c *connector) get(request, expectedResponse string) (err error) {
	resp, err := c.Client.Get(request)
	if err != nil {
		return errors.Wrap(err, "Failed to perform get request")
	}

	defer func() {
		closeErr := resp.Body.Close()
		if err == nil {
			err = errors.Wrap(closeErr, "Failed to close response body")
		}
	}()

	raw, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return errors.Wrap(err, "Failed to read response body")
	}

	body := strings.TrimSpace(string(raw))
	if body != expectedResponse {
		return fmt.Errorf("reqeust to core failed. expected: %s; got: %s", expectedResponse, body)
	}

	return nil
}
