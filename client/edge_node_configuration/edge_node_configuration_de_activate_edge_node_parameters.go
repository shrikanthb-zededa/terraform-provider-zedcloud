// Code generated by go-swagger; DO NOT EDIT.

package edge_node_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
)

// DeactivateEdgeNodeParams creates a new EdgeNodeConfigurationDeActivateEdgeNodeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func DeactivateEdgeNodeParams() *EdgeNodeConfigurationDeActivateEdgeNodeParams {
	return &EdgeNodeConfigurationDeActivateEdgeNodeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewEdgeNodeConfigurationDeActivateEdgeNodeParamsWithTimeout creates a new EdgeNodeConfigurationDeActivateEdgeNodeParams object
// with the ability to set a timeout on a request.
func NewEdgeNodeConfigurationDeActivateEdgeNodeParamsWithTimeout(timeout time.Duration) *EdgeNodeConfigurationDeActivateEdgeNodeParams {
	return &EdgeNodeConfigurationDeActivateEdgeNodeParams{
		timeout: timeout,
	}
}

// NewEdgeNodeConfigurationDeActivateEdgeNodeParamsWithContext creates a new EdgeNodeConfigurationDeActivateEdgeNodeParams object
// with the ability to set a context for a request.
func NewEdgeNodeConfigurationDeActivateEdgeNodeParamsWithContext(ctx context.Context) *EdgeNodeConfigurationDeActivateEdgeNodeParams {
	return &EdgeNodeConfigurationDeActivateEdgeNodeParams{
		Context: ctx,
	}
}

// NewEdgeNodeConfigurationDeActivateEdgeNodeParamsWithHTTPClient creates a new EdgeNodeConfigurationDeActivateEdgeNodeParams object
// with the ability to set a custom HTTPClient for a request.
func NewEdgeNodeConfigurationDeActivateEdgeNodeParamsWithHTTPClient(client *http.Client) *EdgeNodeConfigurationDeActivateEdgeNodeParams {
	return &EdgeNodeConfigurationDeActivateEdgeNodeParams{
		HTTPClient: client,
	}
}

/*
EdgeNodeConfigurationDeActivateEdgeNodeParams contains all the parameters to send to the API endpoint

	for the edge node configuration de activate edge node operation.

	Typically these are written to a http.Request.
*/
type EdgeNodeConfigurationDeActivateEdgeNodeParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* ID.

	   system generated unique id for a device
	*/
	ID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the edge node configuration de activate edge node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNodeConfigurationDeActivateEdgeNodeParams) WithDefaults() *EdgeNodeConfigurationDeActivateEdgeNodeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the edge node configuration de activate edge node params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *EdgeNodeConfigurationDeActivateEdgeNodeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the edge node configuration de activate edge node params
func (o *EdgeNodeConfigurationDeActivateEdgeNodeParams) WithTimeout(timeout time.Duration) *EdgeNodeConfigurationDeActivateEdgeNodeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the edge node configuration de activate edge node params
func (o *EdgeNodeConfigurationDeActivateEdgeNodeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the edge node configuration de activate edge node params
func (o *EdgeNodeConfigurationDeActivateEdgeNodeParams) WithContext(ctx context.Context) *EdgeNodeConfigurationDeActivateEdgeNodeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the edge node configuration de activate edge node params
func (o *EdgeNodeConfigurationDeActivateEdgeNodeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the edge node configuration de activate edge node params
func (o *EdgeNodeConfigurationDeActivateEdgeNodeParams) WithHTTPClient(client *http.Client) *EdgeNodeConfigurationDeActivateEdgeNodeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the edge node configuration de activate edge node params
func (o *EdgeNodeConfigurationDeActivateEdgeNodeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the edge node configuration de activate edge node params
func (o *EdgeNodeConfigurationDeActivateEdgeNodeParams) WithXRequestID(xRequestID *string) *EdgeNodeConfigurationDeActivateEdgeNodeParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the edge node configuration de activate edge node params
func (o *EdgeNodeConfigurationDeActivateEdgeNodeParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithID adds the id to the edge node configuration de activate edge node params
func (o *EdgeNodeConfigurationDeActivateEdgeNodeParams) WithID(id string) *EdgeNodeConfigurationDeActivateEdgeNodeParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the edge node configuration de activate edge node params
func (o *EdgeNodeConfigurationDeActivateEdgeNodeParams) SetID(id string) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *EdgeNodeConfigurationDeActivateEdgeNodeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", o.ID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
