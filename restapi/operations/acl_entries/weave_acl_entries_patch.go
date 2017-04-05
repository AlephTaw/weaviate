/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 Weaviate. All rights reserved.
 * LICENSE: https://github.com/weaviate/weaviate/blob/master/LICENSE
 * AUTHOR: Bob van Luijt (bob@weaviate.com)
 * See www.weaviate.com for details
 * See package.json for author and maintainer info
 * Contact: @weaviate_iot / yourfriends@weaviate.com
 */
 package acl_entries




import (
	"net/http"

	middleware "github.com/go-openapi/runtime/middleware"
)

// WeaveACLEntriesPatchHandlerFunc turns a function with the right signature into a weave acl entries patch handler
type WeaveACLEntriesPatchHandlerFunc func(WeaveACLEntriesPatchParams, interface{}) middleware.Responder

// Handle executing the request and returning a response
func (fn WeaveACLEntriesPatchHandlerFunc) Handle(params WeaveACLEntriesPatchParams, principal interface{}) middleware.Responder {
	return fn(params, principal)
}

// WeaveACLEntriesPatchHandler interface for that can handle valid weave acl entries patch params
type WeaveACLEntriesPatchHandler interface {
	Handle(WeaveACLEntriesPatchParams, interface{}) middleware.Responder
}

// NewWeaveACLEntriesPatch creates a new http.Handler for the weave acl entries patch operation
func NewWeaveACLEntriesPatch(ctx *middleware.Context, handler WeaveACLEntriesPatchHandler) *WeaveACLEntriesPatch {
	return &WeaveACLEntriesPatch{Context: ctx, Handler: handler}
}

/*WeaveACLEntriesPatch swagger:route PATCH /devices/{deviceId}/aclEntries/{aclEntryId} aclEntries weaveAclEntriesPatch

Update an ACL entry. This method supports patch semantics.

*/
type WeaveACLEntriesPatch struct {
	Context *middleware.Context
	Handler WeaveACLEntriesPatchHandler
}

func (o *WeaveACLEntriesPatch) ServeHTTP(rw http.ResponseWriter, r *http.Request) {
	route, _ := o.Context.RouteInfo(r)
	var Params = NewWeaveACLEntriesPatchParams()

	uprinc, err := o.Context.Authorize(r, route)
	if err != nil {
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}
	var principal interface{}
	if uprinc != nil {
		principal = uprinc
	}

	if err := o.Context.BindValidRequest(r, route, &Params); err != nil { // bind params
		o.Context.Respond(rw, r, route.Produces, route, err)
		return
	}

	res := o.Handler.Handle(Params, principal) // actually handle the request

	o.Context.Respond(rw, r, route.Produces, route, res)

}