// Code generated by go-swagger; DO NOT EDIT.

package resource_quota

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

// NewCreateResourceQuotaParams creates a new CreateResourceQuotaParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCreateResourceQuotaParams() *CreateResourceQuotaParams {
	return &CreateResourceQuotaParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCreateResourceQuotaParamsWithTimeout creates a new CreateResourceQuotaParams object
// with the ability to set a timeout on a request.
func NewCreateResourceQuotaParamsWithTimeout(timeout time.Duration) *CreateResourceQuotaParams {
	return &CreateResourceQuotaParams{
		timeout: timeout,
	}
}

// NewCreateResourceQuotaParamsWithContext creates a new CreateResourceQuotaParams object
// with the ability to set a context for a request.
func NewCreateResourceQuotaParamsWithContext(ctx context.Context) *CreateResourceQuotaParams {
	return &CreateResourceQuotaParams{
		Context: ctx,
	}
}

// NewCreateResourceQuotaParamsWithHTTPClient creates a new CreateResourceQuotaParams object
// with the ability to set a custom HTTPClient for a request.
func NewCreateResourceQuotaParamsWithHTTPClient(client *http.Client) *CreateResourceQuotaParams {
	return &CreateResourceQuotaParams{
		HTTPClient: client,
	}
}

/* CreateResourceQuotaParams contains all the parameters to send to the API endpoint
   for the create resource quota operation.

   Typically these are written to a http.Request.
*/
type CreateResourceQuotaParams struct {

	// Body.
	Body CreateResourceQuotaBody

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the create resource quota params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateResourceQuotaParams) WithDefaults() *CreateResourceQuotaParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the create resource quota params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CreateResourceQuotaParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the create resource quota params
func (o *CreateResourceQuotaParams) WithTimeout(timeout time.Duration) *CreateResourceQuotaParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the create resource quota params
func (o *CreateResourceQuotaParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the create resource quota params
func (o *CreateResourceQuotaParams) WithContext(ctx context.Context) *CreateResourceQuotaParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the create resource quota params
func (o *CreateResourceQuotaParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the create resource quota params
func (o *CreateResourceQuotaParams) WithHTTPClient(client *http.Client) *CreateResourceQuotaParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the create resource quota params
func (o *CreateResourceQuotaParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the create resource quota params
func (o *CreateResourceQuotaParams) WithBody(body CreateResourceQuotaBody) *CreateResourceQuotaParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the create resource quota params
func (o *CreateResourceQuotaParams) SetBody(body CreateResourceQuotaBody) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *CreateResourceQuotaParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error
	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}