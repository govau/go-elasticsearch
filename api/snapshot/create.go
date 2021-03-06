// generated by github.com/elastic/go-elasticsearch/cmd/generator; DO NOT EDIT

package snapshot

import (
	"fmt"
	"net/http"
	"net/url"
)

// Create - the snapshot and restore module allows to create snapshots of individual indices or an entire cluster into a remote repository like shared file system, S3, or HDFS. See https://www.elastic.co/guide/en/elasticsearch/reference/5.x/modules-snapshots.html for more info.
//
// repository: a repository name.
//
// snapshot: a snapshot name.
//
// options: optional parameters. Supports the following functional options: WithMasterTimeout, WithWaitForCompletion, WithBody, WithErrorTrace, WithFilterPath, WithHuman, WithPretty, WithSourceParam, see the Option type in this package for more info.
func (s *Snapshot) Create(repository string, snapshot string, options ...*Option) (*CreateResponse, error) {
	req := &http.Request{
		URL: &url.URL{
			Scheme: s.transport.URL.Scheme,
			Host:   s.transport.URL.Host,
		},
		Method: "PUT",
	}
	methodOptions := supportedOptions["Create"]
	for _, option := range options {
		if _, ok := methodOptions[option.name]; !ok {
			return nil, fmt.Errorf("unsupported option: %s", option.name)
		}
		option.apply(req)
	}
	resp, err := s.transport.Do(req)
	return &CreateResponse{resp}, err
}

// CreateResponse is the response for Create
type CreateResponse struct {
	Response *http.Response
	// TODO: fill in structured response
}
