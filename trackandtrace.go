package postmord

import "fmt"

var (
	identifierURL = baseURL + "findByIdentifier.json?apikey=%s&id=%s&locale=%s"
	referenceURL  = baseURL + "findByReference.json?apikey=%s&customerNumber=%s&referenceValue=%s&locale=%s"
)

func (c *Client) FindByIdentifier(id string) (*Response, error) {
	return c.request(fmt.Sprintf(identifierURL, c.APIKey, id, c.Opts.Locale.String()))
}

func (c *Client) FindByReference(customerNumber, reference string) (*Response, error) {
	return c.request(fmt.Sprintf(referenceURL, c.APIKey, customerNumber, reference, c.Opts.Locale.String()))
}
