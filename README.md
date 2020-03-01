# postmord
A Go API wrapper for the PostNord API

## Example
```bash
$ go get github.com/antonjah/postmord
```

```go
package main

import (
	"fmt"
	"github.com/antonjah/postmord"
	"io/ioutil"
)

func main() {
	c := postmord.NewClient("MyConsumerAPIKey", nil, postmord.DefaultOptions)
	postmordResponse, err := c.FindByIdentifier("MyShipmentID")
	if err != nil {
		panic(err)
}

	// Output status
	for _, shipment := range postmordResponse.Shipments {
		fmt.Println(shipment.ShipmentID, shipment.StatusText)
	}

	// Write to JSON/XML/YAML file
	j, _ := postmordResponse.JSON()
	if err := ioutil.WriteFile("MyShipments.json", j, 0755); err != nil {
		panic(err)
	}
}
```

## Author
@antonjah
