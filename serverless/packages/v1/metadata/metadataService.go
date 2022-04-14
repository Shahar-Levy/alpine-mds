package main

import (
	"errors"
	"fmt"
	"os"

	httptransport "github.com/go-openapi/runtime/client"
	"github.com/netbox-community/go-netbox/netbox/client"
	"github.com/netbox-community/go-netbox/netbox/client/dcim"
)

type Request struct {
	Location string `json:"location"`
}

type Response struct {
	StatusCode int               `json:"statusCode,omitempty"`
	Headers    map[string]string `json:"headers,omitempty"`
	Body       string            `json:"body,omitempty"`
}

func Main(in Request) (*Response, error) {
	token := os.Getenv("NETBOX_TOKEN")
	if token == "" {
		return nil, errors.New("Please provide netbox API token via env var NETBOX_TOKEN")
	}

	netboxHost := os.Getenv("NETBOX_HOST")
	if netboxHost == "" {
		return nil, errors.New("Please provide netbox host via env var NETBOX_HOST")
	}

	transport := httptransport.New(netboxHost, client.DefaultBasePath, []string{"https"})
	transport.DefaultAuthentication = httptransport.APIKeyAuth("Authorization", "header", "Token "+token)

	c := client.New(transport, nil)

	req := dcim.NewDcimSitesListParams()
	res, err := c.Dcim.DcimSitesList(req, nil)

	if err != nil {
		return nil, errors.New("Cannot get sites list: %v", err)
	}

	return &Response{
		Body: fmt.Sprintf("Sites %s", res.Payload.Results),
	}, nil
}
