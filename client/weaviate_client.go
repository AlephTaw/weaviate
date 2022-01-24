//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package client

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/runtime"
	httptransport "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"

	"github.com/semi-technologies/weaviate/client/batch"
	"github.com/semi-technologies/weaviate/client/classifications"
	"github.com/semi-technologies/weaviate/client/graphql"
	"github.com/semi-technologies/weaviate/client/meta"
	"github.com/semi-technologies/weaviate/client/objects"
	"github.com/semi-technologies/weaviate/client/operations"
	"github.com/semi-technologies/weaviate/client/schema"
	"github.com/semi-technologies/weaviate/client/well_known"
)

// Default weaviate HTTP client.
var Default = NewHTTPClient(nil)

const (
	// DefaultHost is the default Host
	// found in Meta (info) section of spec file
	DefaultHost string = "localhost"
	// DefaultBasePath is the default BasePath
	// found in Meta (info) section of spec file
	DefaultBasePath string = "/v1"
)

// DefaultSchemes are the default schemes found in Meta (info) section of spec file
var DefaultSchemes = []string{"https"}

// NewHTTPClient creates a new weaviate HTTP client.
func NewHTTPClient(formats strfmt.Registry) *Weaviate {
	return NewHTTPClientWithConfig(formats, nil)
}

// NewHTTPClientWithConfig creates a new weaviate HTTP client,
// using a customizable transport config.
func NewHTTPClientWithConfig(formats strfmt.Registry, cfg *TransportConfig) *Weaviate {
	// ensure nullable parameters have default
	if cfg == nil {
		cfg = DefaultTransportConfig()
	}

	// create transport and client
	transport := httptransport.New(cfg.Host, cfg.BasePath, cfg.Schemes)
	return New(transport, formats)
}

// New creates a new weaviate client
func New(transport runtime.ClientTransport, formats strfmt.Registry) *Weaviate {
	// ensure nullable parameters have default
	if formats == nil {
		formats = strfmt.Default
	}

	cli := new(Weaviate)
	cli.Transport = transport
	cli.Batch = batch.New(transport, formats)
	cli.Classifications = classifications.New(transport, formats)
	cli.Graphql = graphql.New(transport, formats)
	cli.Meta = meta.New(transport, formats)
	cli.Objects = objects.New(transport, formats)
	cli.Operations = operations.New(transport, formats)
	cli.Schema = schema.New(transport, formats)
	cli.WellKnown = well_known.New(transport, formats)
	return cli
}

// DefaultTransportConfig creates a TransportConfig with the
// default settings taken from the meta section of the spec file.
func DefaultTransportConfig() *TransportConfig {
	return &TransportConfig{
		Host:     DefaultHost,
		BasePath: DefaultBasePath,
		Schemes:  DefaultSchemes,
	}
}

// TransportConfig contains the transport related info,
// found in the meta section of the spec file.
type TransportConfig struct {
	Host     string
	BasePath string
	Schemes  []string
}

// WithHost overrides the default host,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithHost(host string) *TransportConfig {
	cfg.Host = host
	return cfg
}

// WithBasePath overrides the default basePath,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithBasePath(basePath string) *TransportConfig {
	cfg.BasePath = basePath
	return cfg
}

// WithSchemes overrides the default schemes,
// provided by the meta section of the spec file.
func (cfg *TransportConfig) WithSchemes(schemes []string) *TransportConfig {
	cfg.Schemes = schemes
	return cfg
}

// Weaviate is a client for weaviate
type Weaviate struct {
	Batch batch.ClientService

	Classifications classifications.ClientService

	Graphql graphql.ClientService

	Meta meta.ClientService

	Objects objects.ClientService

	Operations operations.ClientService

	Schema schema.ClientService

	WellKnown well_known.ClientService

	Transport runtime.ClientTransport
}

// SetTransport changes the transport on the client and all its subresources
func (c *Weaviate) SetTransport(transport runtime.ClientTransport) {
	c.Transport = transport
	c.Batch.SetTransport(transport)
	c.Classifications.SetTransport(transport)
	c.Graphql.SetTransport(transport)
	c.Meta.SetTransport(transport)
	c.Objects.SetTransport(transport)
	c.Operations.SetTransport(transport)
	c.Schema.SetTransport(transport)
	c.WellKnown.SetTransport(transport)
}
