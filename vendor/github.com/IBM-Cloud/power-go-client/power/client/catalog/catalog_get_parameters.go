// Code generated by go-swagger; DO NOT EDIT.

package catalog

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

// NewCatalogGetParams creates a new CatalogGetParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewCatalogGetParams() *CatalogGetParams {
	return &CatalogGetParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewCatalogGetParamsWithTimeout creates a new CatalogGetParams object
// with the ability to set a timeout on a request.
func NewCatalogGetParamsWithTimeout(timeout time.Duration) *CatalogGetParams {
	return &CatalogGetParams{
		timeout: timeout,
	}
}

// NewCatalogGetParamsWithContext creates a new CatalogGetParams object
// with the ability to set a context for a request.
func NewCatalogGetParamsWithContext(ctx context.Context) *CatalogGetParams {
	return &CatalogGetParams{
		Context: ctx,
	}
}

// NewCatalogGetParamsWithHTTPClient creates a new CatalogGetParams object
// with the ability to set a custom HTTPClient for a request.
func NewCatalogGetParamsWithHTTPClient(client *http.Client) *CatalogGetParams {
	return &CatalogGetParams{
		HTTPClient: client,
	}
}

/*
CatalogGetParams contains all the parameters to send to the API endpoint

	for the catalog get operation.

	Typically these are written to a http.Request.
*/
type CatalogGetParams struct {

	/* XBrokerAPIVersion.

	   version number of the Service Broker API that the Platform will use
	*/
	XBrokerAPIVersion string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the catalog get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CatalogGetParams) WithDefaults() *CatalogGetParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the catalog get params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *CatalogGetParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the catalog get params
func (o *CatalogGetParams) WithTimeout(timeout time.Duration) *CatalogGetParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the catalog get params
func (o *CatalogGetParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the catalog get params
func (o *CatalogGetParams) WithContext(ctx context.Context) *CatalogGetParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the catalog get params
func (o *CatalogGetParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the catalog get params
func (o *CatalogGetParams) WithHTTPClient(client *http.Client) *CatalogGetParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the catalog get params
func (o *CatalogGetParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXBrokerAPIVersion adds the xBrokerAPIVersion to the catalog get params
func (o *CatalogGetParams) WithXBrokerAPIVersion(xBrokerAPIVersion string) *CatalogGetParams {
	o.SetXBrokerAPIVersion(xBrokerAPIVersion)
	return o
}

// SetXBrokerAPIVersion adds the xBrokerApiVersion to the catalog get params
func (o *CatalogGetParams) SetXBrokerAPIVersion(xBrokerAPIVersion string) {
	o.XBrokerAPIVersion = xBrokerAPIVersion
}

// WriteToRequest writes these params to a swagger request
func (o *CatalogGetParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param X-Broker-API-Version
	if err := r.SetHeaderParam("X-Broker-API-Version", o.XBrokerAPIVersion); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}