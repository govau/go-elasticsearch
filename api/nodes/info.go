// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package nodes

import (
	"fmt"
	"net/http"
	"net/url"
)

// Info - the cluster nodes info API allows to retrieve one or more (or all) of the cluster nodes information. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/cluster-nodes-info.html for more info.
//
// options: optional parameters. Supports the following functional options: WithMetric, WithNodeID, WithFlatSettings, WithTimeout, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (n *Nodes) Info(options ...*Option) (*InfoResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: n.transport.URL.Scheme,
			Host:   n.transport.URL.Host,
		},
		Method: "GET",
	}
	methodOptions := supportedOptions["Info"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := n.transport.Do(req)
	return &InfoResponse{resp}, err
}

// InfoResponse is the response for Info
type InfoResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}
