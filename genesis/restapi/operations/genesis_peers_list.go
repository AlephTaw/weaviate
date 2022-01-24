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

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the generate command

import (
	"net/http"

	"github.com/go-openapi/runtime/middleware"
)

// GenesisPeersListHandlerFunc turns a function with the right signature into a genesis peers list handler
type GenesisPeersListHandlerFunc func(GenesisPeersListParams) middleware.Responder

// Handle executing the request and returning a response
func (fn GenesisPeersListHandlerFunc) Handle(params GenesisPeersListParams) middleware.Responder {
	return fn(params)
}

// GenesisPeersListHandler interface for that can handle valid genesis peers list params
type GenesisPeersListHandler interface {
	Handle(GenesisPeersListParams) middleware.Responder
}

// NewGenesisPeersList creates a new http.Handler for the genesis peers list operation
func NewGenesisPeersList(ctx *middleware.Context, handler GenesisPeersListHandler) *GenesisPeersList {
	return &GenesisPeersList{Context: ctx, Handler: handler}
}

/*GenesisPeersList swagger:route GET /peers genesisPeersList

List the registered peers

*/
type GenesisPeersList struct {
	Context *middleware.Context
	Handler GenesisPeersListHandler
}

func (o *GenesisPeersList) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, rCtx, _ := o.Context.RouteInfo(r)
	if rCtx != nil {
		r = rCtx
	}
	Params := NewGenesisPeersListParams()

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)
}
