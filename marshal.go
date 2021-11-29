package postmord

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"github.com/go-yaml/yaml"
)

type encoding string

const (
	encodingYAML encoding = "yaml"
	encodingXML  encoding = "xml"
	encodingJSON encoding = "json"
)

func (r *Response) JSON() ([]byte, error) { return r.marshal(encodingJSON) }
func (r *Response) YAML() ([]byte, error) { return r.marshal(encodingYAML) }
func (r *Response) XML() ([]byte, error)  { return r.marshal(encodingXML) }

func (r *Response) marshal(e encoding) ([]byte, error) {
	switch e {
	case encodingYAML:
		return yaml.Marshal(r)
	case encodingJSON:
		return json.Marshal(r)
	case encodingXML:
		return xml.Marshal(r)
	default:
		return nil, fmt.Errorf("unknown encoding: %s", e)
	}
}
