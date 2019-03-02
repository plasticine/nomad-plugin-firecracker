// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2018 Amazon.com, Inc. or its affiliates. All Rights Reserved.
//
// Licensed under the Apache License, Version 2.0 (the "License"). You may
// not use this file except in compliance with the License. A copy of the
// License is located at
//
// 	http://aws.amazon.com/apache2.0/
//
// or in the "license" file accompanying this file. This file is distributed
// on an "AS IS" BASIS, WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either
// express or implied. See the License for the specific language governing
// permissions and limitations under the License.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	client_models "github.com/firecracker-microvm/firecracker-go-sdk/client/models"
)

// NewPatchGuestNetworkInterfaceByIDParams creates a new PatchGuestNetworkInterfaceByIDParams object
// with the default values initialized.
func NewPatchGuestNetworkInterfaceByIDParams() *PatchGuestNetworkInterfaceByIDParams {
	var ()
	return &PatchGuestNetworkInterfaceByIDParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewPatchGuestNetworkInterfaceByIDParamsWithTimeout creates a new PatchGuestNetworkInterfaceByIDParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewPatchGuestNetworkInterfaceByIDParamsWithTimeout(timeout time.Duration) *PatchGuestNetworkInterfaceByIDParams {
	var ()
	return &PatchGuestNetworkInterfaceByIDParams{

		timeout: timeout,
	}
}

// NewPatchGuestNetworkInterfaceByIDParamsWithContext creates a new PatchGuestNetworkInterfaceByIDParams object
// with the default values initialized, and the ability to set a context for a request
func NewPatchGuestNetworkInterfaceByIDParamsWithContext(ctx context.Context) *PatchGuestNetworkInterfaceByIDParams {
	var ()
	return &PatchGuestNetworkInterfaceByIDParams{

		Context: ctx,
	}
}

// NewPatchGuestNetworkInterfaceByIDParamsWithHTTPClient creates a new PatchGuestNetworkInterfaceByIDParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewPatchGuestNetworkInterfaceByIDParamsWithHTTPClient(client *http.Client) *PatchGuestNetworkInterfaceByIDParams {
	var ()
	return &PatchGuestNetworkInterfaceByIDParams{
		HTTPClient: client,
	}
}

/*PatchGuestNetworkInterfaceByIDParams contains all the parameters to send to the API endpoint
for the patch guest network interface by ID operation typically these are written to a http.Request
*/
type PatchGuestNetworkInterfaceByIDParams struct {

	/*Body
	  A subset of the guest network interface properties

	*/
	Body *client_models.PartialNetworkInterface
	/*IfaceID
	  The id of the guest network interface

	*/
	IfaceID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the patch guest network interface by ID params
func (o *PatchGuestNetworkInterfaceByIDParams) WithTimeout(timeout time.Duration) *PatchGuestNetworkInterfaceByIDParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch guest network interface by ID params
func (o *PatchGuestNetworkInterfaceByIDParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch guest network interface by ID params
func (o *PatchGuestNetworkInterfaceByIDParams) WithContext(ctx context.Context) *PatchGuestNetworkInterfaceByIDParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch guest network interface by ID params
func (o *PatchGuestNetworkInterfaceByIDParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch guest network interface by ID params
func (o *PatchGuestNetworkInterfaceByIDParams) WithHTTPClient(client *http.Client) *PatchGuestNetworkInterfaceByIDParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch guest network interface by ID params
func (o *PatchGuestNetworkInterfaceByIDParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the patch guest network interface by ID params
func (o *PatchGuestNetworkInterfaceByIDParams) WithBody(body *client_models.PartialNetworkInterface) *PatchGuestNetworkInterfaceByIDParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the patch guest network interface by ID params
func (o *PatchGuestNetworkInterfaceByIDParams) SetBody(body *client_models.PartialNetworkInterface) {
	o.Body = body
}

// WithIfaceID adds the ifaceID to the patch guest network interface by ID params
func (o *PatchGuestNetworkInterfaceByIDParams) WithIfaceID(ifaceID string) *PatchGuestNetworkInterfaceByIDParams {
	o.SetIfaceID(ifaceID)
	return o
}

// SetIfaceID adds the ifaceId to the patch guest network interface by ID params
func (o *PatchGuestNetworkInterfaceByIDParams) SetIfaceID(ifaceID string) {
	o.IfaceID = ifaceID
}

// WriteToRequest writes these params to a swagger request
func (o *PatchGuestNetworkInterfaceByIDParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	// path param iface_id
	if err := r.SetPathParam("iface_id", o.IfaceID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
