// Code generated by go-swagger; DO NOT EDIT.

package mlaadminsetting

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

// NewGetMLAAdminSettingParams creates a new GetMLAAdminSettingParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewGetMLAAdminSettingParams() *GetMLAAdminSettingParams {
	return &GetMLAAdminSettingParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewGetMLAAdminSettingParamsWithTimeout creates a new GetMLAAdminSettingParams object
// with the ability to set a timeout on a request.
func NewGetMLAAdminSettingParamsWithTimeout(timeout time.Duration) *GetMLAAdminSettingParams {
	return &GetMLAAdminSettingParams{
		timeout: timeout,
	}
}

// NewGetMLAAdminSettingParamsWithContext creates a new GetMLAAdminSettingParams object
// with the ability to set a context for a request.
func NewGetMLAAdminSettingParamsWithContext(ctx context.Context) *GetMLAAdminSettingParams {
	return &GetMLAAdminSettingParams{
		Context: ctx,
	}
}

// NewGetMLAAdminSettingParamsWithHTTPClient creates a new GetMLAAdminSettingParams object
// with the ability to set a custom HTTPClient for a request.
func NewGetMLAAdminSettingParamsWithHTTPClient(client *http.Client) *GetMLAAdminSettingParams {
	return &GetMLAAdminSettingParams{
		HTTPClient: client,
	}
}

/* GetMLAAdminSettingParams contains all the parameters to send to the API endpoint
   for the get m l a admin setting operation.

   Typically these are written to a http.Request.
*/
type GetMLAAdminSettingParams struct {

	// ClusterID.
	ClusterID string

	// ProjectID.
	ProjectID string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the get m l a admin setting params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMLAAdminSettingParams) WithDefaults() *GetMLAAdminSettingParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the get m l a admin setting params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *GetMLAAdminSettingParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the get m l a admin setting params
func (o *GetMLAAdminSettingParams) WithTimeout(timeout time.Duration) *GetMLAAdminSettingParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get m l a admin setting params
func (o *GetMLAAdminSettingParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get m l a admin setting params
func (o *GetMLAAdminSettingParams) WithContext(ctx context.Context) *GetMLAAdminSettingParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get m l a admin setting params
func (o *GetMLAAdminSettingParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get m l a admin setting params
func (o *GetMLAAdminSettingParams) WithHTTPClient(client *http.Client) *GetMLAAdminSettingParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get m l a admin setting params
func (o *GetMLAAdminSettingParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithClusterID adds the clusterID to the get m l a admin setting params
func (o *GetMLAAdminSettingParams) WithClusterID(clusterID string) *GetMLAAdminSettingParams {
	o.SetClusterID(clusterID)
	return o
}

// SetClusterID adds the clusterId to the get m l a admin setting params
func (o *GetMLAAdminSettingParams) SetClusterID(clusterID string) {
	o.ClusterID = clusterID
}

// WithProjectID adds the projectID to the get m l a admin setting params
func (o *GetMLAAdminSettingParams) WithProjectID(projectID string) *GetMLAAdminSettingParams {
	o.SetProjectID(projectID)
	return o
}

// SetProjectID adds the projectId to the get m l a admin setting params
func (o *GetMLAAdminSettingParams) SetProjectID(projectID string) {
	o.ProjectID = projectID
}

// WriteToRequest writes these params to a swagger request
func (o *GetMLAAdminSettingParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param cluster_id
	if err := r.SetPathParam("cluster_id", o.ClusterID); err != nil {
		return err
	}

	// path param project_id
	if err := r.SetPathParam("project_id", o.ProjectID); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}