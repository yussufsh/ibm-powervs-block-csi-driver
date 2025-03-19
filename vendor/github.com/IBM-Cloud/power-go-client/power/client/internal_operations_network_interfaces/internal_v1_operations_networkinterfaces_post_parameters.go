// Code generated by go-swagger; DO NOT EDIT.

package internal_operations_network_interfaces

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

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// NewInternalV1OperationsNetworkinterfacesPostParams creates a new InternalV1OperationsNetworkinterfacesPostParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewInternalV1OperationsNetworkinterfacesPostParams() *InternalV1OperationsNetworkinterfacesPostParams {
	return &InternalV1OperationsNetworkinterfacesPostParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewInternalV1OperationsNetworkinterfacesPostParamsWithTimeout creates a new InternalV1OperationsNetworkinterfacesPostParams object
// with the ability to set a timeout on a request.
func NewInternalV1OperationsNetworkinterfacesPostParamsWithTimeout(timeout time.Duration) *InternalV1OperationsNetworkinterfacesPostParams {
	return &InternalV1OperationsNetworkinterfacesPostParams{
		timeout: timeout,
	}
}

// NewInternalV1OperationsNetworkinterfacesPostParamsWithContext creates a new InternalV1OperationsNetworkinterfacesPostParams object
// with the ability to set a context for a request.
func NewInternalV1OperationsNetworkinterfacesPostParamsWithContext(ctx context.Context) *InternalV1OperationsNetworkinterfacesPostParams {
	return &InternalV1OperationsNetworkinterfacesPostParams{
		Context: ctx,
	}
}

// NewInternalV1OperationsNetworkinterfacesPostParamsWithHTTPClient creates a new InternalV1OperationsNetworkinterfacesPostParams object
// with the ability to set a custom HTTPClient for a request.
func NewInternalV1OperationsNetworkinterfacesPostParamsWithHTTPClient(client *http.Client) *InternalV1OperationsNetworkinterfacesPostParams {
	return &InternalV1OperationsNetworkinterfacesPostParams{
		HTTPClient: client,
	}
}

/*
InternalV1OperationsNetworkinterfacesPostParams contains all the parameters to send to the API endpoint

	for the internal v1 operations networkinterfaces post operation.

	Typically these are written to a http.Request.
*/
type InternalV1OperationsNetworkinterfacesPostParams struct {

	/* Authorization.

	   Authentication of the service token
	*/
	Authorization string

	/* CRN.

	   the CRN of the workspace
	*/
	CRN string

	/* IBMUserAuthorization.

	   Authentication of the operation account user
	*/
	IBMUserAuthorization string

	/* Body.

	   Parameters for creating a Network Interface CRN
	*/
	Body *models.InternalOperationsRequest

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the internal v1 operations networkinterfaces post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InternalV1OperationsNetworkinterfacesPostParams) WithDefaults() *InternalV1OperationsNetworkinterfacesPostParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the internal v1 operations networkinterfaces post params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *InternalV1OperationsNetworkinterfacesPostParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the internal v1 operations networkinterfaces post params
func (o *InternalV1OperationsNetworkinterfacesPostParams) WithTimeout(timeout time.Duration) *InternalV1OperationsNetworkinterfacesPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the internal v1 operations networkinterfaces post params
func (o *InternalV1OperationsNetworkinterfacesPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the internal v1 operations networkinterfaces post params
func (o *InternalV1OperationsNetworkinterfacesPostParams) WithContext(ctx context.Context) *InternalV1OperationsNetworkinterfacesPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the internal v1 operations networkinterfaces post params
func (o *InternalV1OperationsNetworkinterfacesPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the internal v1 operations networkinterfaces post params
func (o *InternalV1OperationsNetworkinterfacesPostParams) WithHTTPClient(client *http.Client) *InternalV1OperationsNetworkinterfacesPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the internal v1 operations networkinterfaces post params
func (o *InternalV1OperationsNetworkinterfacesPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithAuthorization adds the authorization to the internal v1 operations networkinterfaces post params
func (o *InternalV1OperationsNetworkinterfacesPostParams) WithAuthorization(authorization string) *InternalV1OperationsNetworkinterfacesPostParams {
	o.SetAuthorization(authorization)
	return o
}

// SetAuthorization adds the authorization to the internal v1 operations networkinterfaces post params
func (o *InternalV1OperationsNetworkinterfacesPostParams) SetAuthorization(authorization string) {
	o.Authorization = authorization
}

// WithCRN adds the cRN to the internal v1 operations networkinterfaces post params
func (o *InternalV1OperationsNetworkinterfacesPostParams) WithCRN(cRN string) *InternalV1OperationsNetworkinterfacesPostParams {
	o.SetCRN(cRN)
	return o
}

// SetCRN adds the cRN to the internal v1 operations networkinterfaces post params
func (o *InternalV1OperationsNetworkinterfacesPostParams) SetCRN(cRN string) {
	o.CRN = cRN
}

// WithIBMUserAuthorization adds the iBMUserAuthorization to the internal v1 operations networkinterfaces post params
func (o *InternalV1OperationsNetworkinterfacesPostParams) WithIBMUserAuthorization(iBMUserAuthorization string) *InternalV1OperationsNetworkinterfacesPostParams {
	o.SetIBMUserAuthorization(iBMUserAuthorization)
	return o
}

// SetIBMUserAuthorization adds the iBMUserAuthorization to the internal v1 operations networkinterfaces post params
func (o *InternalV1OperationsNetworkinterfacesPostParams) SetIBMUserAuthorization(iBMUserAuthorization string) {
	o.IBMUserAuthorization = iBMUserAuthorization
}

// WithBody adds the body to the internal v1 operations networkinterfaces post params
func (o *InternalV1OperationsNetworkinterfacesPostParams) WithBody(body *models.InternalOperationsRequest) *InternalV1OperationsNetworkinterfacesPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the internal v1 operations networkinterfaces post params
func (o *InternalV1OperationsNetworkinterfacesPostParams) SetBody(body *models.InternalOperationsRequest) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *InternalV1OperationsNetworkinterfacesPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// header param Authorization
	if err := r.SetHeaderParam("Authorization", o.Authorization); err != nil {
		return err
	}

	// header param CRN
	if err := r.SetHeaderParam("CRN", o.CRN); err != nil {
		return err
	}

	// header param IBM-UserAuthorization
	if err := r.SetHeaderParam("IBM-UserAuthorization", o.IBMUserAuthorization); err != nil {
		return err
	}
	if o.Body != nil {
		if err := r.SetBodyParam(o.Body); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
