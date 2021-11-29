# postmord
A Go API wrapper for the PostNord track and trace API

## Example
```bash
$ go get -u github.com/antonjah/postmord
```

```go
package main

import (
	"fmt"
	"github.com/antonjah/postmord"
)

func main() {
	client := postmord.NewClient("MyConsumerAPIKey", nil, postmord.DefaultOptions)
	response, err := client.FindByIdentifier("MyShipmentID")
	if err != nil {
		// handle error
	}

	for _, shipment := range response.TrackingInformation.Shipments {
		fmt.Println(shipment.ShipmentID, shipment.StatusText)
	}

	// JSON
	j, _ := response.JSON()

	// YAML
	y, _ := response.YAML()

	// XML
	x, _ := response.XML()
}
```

## Author
@antonjah
