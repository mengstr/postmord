package postmord

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
)

func (c *Client) request(url string) (*Response, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return nil, err
	}

	req.Header.Add("User-Agent", userAgent)

	res, err := c.httpClient.Do(req)
	if err != nil {
		return nil, err
	}
	defer func() {
		if err = res.Body.Close(); err != nil {
			fmt.Printf("failed to close body: %s\n", err.Error())
		}
	}()

	// according to swagger it only responds 400 or 200 - so check for != 200
	if res.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("request failed: %s", res.Status)
	}

	b, err := ioutil.ReadAll(res.Body)
	if err != nil {
		return nil, err
	}

	response := Response{}
	if err = json.Unmarshal(b, &response); err != nil {
		return nil, err
	}

	if faults := response.TrackingInformation.CompositeFault.Faults; len(faults) > 0 {
		return &response, fmt.Errorf("%s: %s", faults[0].FaultCode, faults[0].ExplanationText)
	}

	return &response, nil
}
