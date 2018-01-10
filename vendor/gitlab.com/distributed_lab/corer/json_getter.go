package corer

import (
	"encoding/json"
	"gitlab.com/distributed_lab/logan/v3/errors"
	"io/ioutil"
	"net/http"
)

func (c *connector) getJSON(request string, response interface{}) (err error) {
	var httpResponse *http.Response
	httpResponse, err = c.Client.Get(request)
	if err != nil {
		return errors.Wrap(err, "Failed to perform get request")
	}

	defer func() {
		closeErr := httpResponse.Body.Close()
		if err == nil {
			err = errors.Wrap(closeErr, "Failed to close response body")
		}
	}()

	contents, err := ioutil.ReadAll(httpResponse.Body)
	if err != nil {
		return errors.Wrap(err, "Failed to read all response body")
	}

	err = json.Unmarshal(contents, response)
	if err != nil {
		return errors.Wrap(err, "Failed to unmarshal response")
	}

	return nil
}
