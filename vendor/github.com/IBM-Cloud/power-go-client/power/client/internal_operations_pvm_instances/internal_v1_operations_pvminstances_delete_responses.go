// Code generated by go-swagger; DO NOT EDIT.

package internal_operations_pvm_instances

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// InternalV1OperationsPvminstancesDeleteReader is a Reader for the InternalV1OperationsPvminstancesDelete structure.
type InternalV1OperationsPvminstancesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InternalV1OperationsPvminstancesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewInternalV1OperationsPvminstancesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInternalV1OperationsPvminstancesDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewInternalV1OperationsPvminstancesDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewInternalV1OperationsPvminstancesDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewInternalV1OperationsPvminstancesDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewInternalV1OperationsPvminstancesDeleteGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewInternalV1OperationsPvminstancesDeleteTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewInternalV1OperationsPvminstancesDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /internal/v1/operations/pvm-instances/{resource_crn}] internal.v1.operations.pvminstances.delete", response, response.Code())
	}
}

// NewInternalV1OperationsPvminstancesDeleteNoContent creates a InternalV1OperationsPvminstancesDeleteNoContent with default headers values
func NewInternalV1OperationsPvminstancesDeleteNoContent() *InternalV1OperationsPvminstancesDeleteNoContent {
	return &InternalV1OperationsPvminstancesDeleteNoContent{}
}

/*
InternalV1OperationsPvminstancesDeleteNoContent describes a response with status code 204, with default header values.

Deleted
*/
type InternalV1OperationsPvminstancesDeleteNoContent struct {
}

// IsSuccess returns true when this internal v1 operations pvminstances delete no content response has a 2xx status code
func (o *InternalV1OperationsPvminstancesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this internal v1 operations pvminstances delete no content response has a 3xx status code
func (o *InternalV1OperationsPvminstancesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 operations pvminstances delete no content response has a 4xx status code
func (o *InternalV1OperationsPvminstancesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this internal v1 operations pvminstances delete no content response has a 5xx status code
func (o *InternalV1OperationsPvminstancesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 operations pvminstances delete no content response a status code equal to that given
func (o *InternalV1OperationsPvminstancesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the internal v1 operations pvminstances delete no content response
func (o *InternalV1OperationsPvminstancesDeleteNoContent) Code() int {
	return 204
}

func (o *InternalV1OperationsPvminstancesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /internal/v1/operations/pvm-instances/{resource_crn}][%d] internalV1OperationsPvminstancesDeleteNoContent", 204)
}

func (o *InternalV1OperationsPvminstancesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /internal/v1/operations/pvm-instances/{resource_crn}][%d] internalV1OperationsPvminstancesDeleteNoContent", 204)
}

func (o *InternalV1OperationsPvminstancesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInternalV1OperationsPvminstancesDeleteBadRequest creates a InternalV1OperationsPvminstancesDeleteBadRequest with default headers values
func NewInternalV1OperationsPvminstancesDeleteBadRequest() *InternalV1OperationsPvminstancesDeleteBadRequest {
	return &InternalV1OperationsPvminstancesDeleteBadRequest{}
}

/*
InternalV1OperationsPvminstancesDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type InternalV1OperationsPvminstancesDeleteBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 operations pvminstances delete bad request response has a 2xx status code
func (o *InternalV1OperationsPvminstancesDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 operations pvminstances delete bad request response has a 3xx status code
func (o *InternalV1OperationsPvminstancesDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 operations pvminstances delete bad request response has a 4xx status code
func (o *InternalV1OperationsPvminstancesDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 operations pvminstances delete bad request response has a 5xx status code
func (o *InternalV1OperationsPvminstancesDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 operations pvminstances delete bad request response a status code equal to that given
func (o *InternalV1OperationsPvminstancesDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the internal v1 operations pvminstances delete bad request response
func (o *InternalV1OperationsPvminstancesDeleteBadRequest) Code() int {
	return 400
}

func (o *InternalV1OperationsPvminstancesDeleteBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/pvm-instances/{resource_crn}][%d] internalV1OperationsPvminstancesDeleteBadRequest %s", 400, payload)
}

func (o *InternalV1OperationsPvminstancesDeleteBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/pvm-instances/{resource_crn}][%d] internalV1OperationsPvminstancesDeleteBadRequest %s", 400, payload)
}

func (o *InternalV1OperationsPvminstancesDeleteBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1OperationsPvminstancesDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1OperationsPvminstancesDeleteUnauthorized creates a InternalV1OperationsPvminstancesDeleteUnauthorized with default headers values
func NewInternalV1OperationsPvminstancesDeleteUnauthorized() *InternalV1OperationsPvminstancesDeleteUnauthorized {
	return &InternalV1OperationsPvminstancesDeleteUnauthorized{}
}

/*
InternalV1OperationsPvminstancesDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type InternalV1OperationsPvminstancesDeleteUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 operations pvminstances delete unauthorized response has a 2xx status code
func (o *InternalV1OperationsPvminstancesDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 operations pvminstances delete unauthorized response has a 3xx status code
func (o *InternalV1OperationsPvminstancesDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 operations pvminstances delete unauthorized response has a 4xx status code
func (o *InternalV1OperationsPvminstancesDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 operations pvminstances delete unauthorized response has a 5xx status code
func (o *InternalV1OperationsPvminstancesDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 operations pvminstances delete unauthorized response a status code equal to that given
func (o *InternalV1OperationsPvminstancesDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the internal v1 operations pvminstances delete unauthorized response
func (o *InternalV1OperationsPvminstancesDeleteUnauthorized) Code() int {
	return 401
}

func (o *InternalV1OperationsPvminstancesDeleteUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/pvm-instances/{resource_crn}][%d] internalV1OperationsPvminstancesDeleteUnauthorized %s", 401, payload)
}

func (o *InternalV1OperationsPvminstancesDeleteUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/pvm-instances/{resource_crn}][%d] internalV1OperationsPvminstancesDeleteUnauthorized %s", 401, payload)
}

func (o *InternalV1OperationsPvminstancesDeleteUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1OperationsPvminstancesDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1OperationsPvminstancesDeleteForbidden creates a InternalV1OperationsPvminstancesDeleteForbidden with default headers values
func NewInternalV1OperationsPvminstancesDeleteForbidden() *InternalV1OperationsPvminstancesDeleteForbidden {
	return &InternalV1OperationsPvminstancesDeleteForbidden{}
}

/*
InternalV1OperationsPvminstancesDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type InternalV1OperationsPvminstancesDeleteForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 operations pvminstances delete forbidden response has a 2xx status code
func (o *InternalV1OperationsPvminstancesDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 operations pvminstances delete forbidden response has a 3xx status code
func (o *InternalV1OperationsPvminstancesDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 operations pvminstances delete forbidden response has a 4xx status code
func (o *InternalV1OperationsPvminstancesDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 operations pvminstances delete forbidden response has a 5xx status code
func (o *InternalV1OperationsPvminstancesDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 operations pvminstances delete forbidden response a status code equal to that given
func (o *InternalV1OperationsPvminstancesDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the internal v1 operations pvminstances delete forbidden response
func (o *InternalV1OperationsPvminstancesDeleteForbidden) Code() int {
	return 403
}

func (o *InternalV1OperationsPvminstancesDeleteForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/pvm-instances/{resource_crn}][%d] internalV1OperationsPvminstancesDeleteForbidden %s", 403, payload)
}

func (o *InternalV1OperationsPvminstancesDeleteForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/pvm-instances/{resource_crn}][%d] internalV1OperationsPvminstancesDeleteForbidden %s", 403, payload)
}

func (o *InternalV1OperationsPvminstancesDeleteForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1OperationsPvminstancesDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1OperationsPvminstancesDeleteNotFound creates a InternalV1OperationsPvminstancesDeleteNotFound with default headers values
func NewInternalV1OperationsPvminstancesDeleteNotFound() *InternalV1OperationsPvminstancesDeleteNotFound {
	return &InternalV1OperationsPvminstancesDeleteNotFound{}
}

/*
InternalV1OperationsPvminstancesDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type InternalV1OperationsPvminstancesDeleteNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 operations pvminstances delete not found response has a 2xx status code
func (o *InternalV1OperationsPvminstancesDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 operations pvminstances delete not found response has a 3xx status code
func (o *InternalV1OperationsPvminstancesDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 operations pvminstances delete not found response has a 4xx status code
func (o *InternalV1OperationsPvminstancesDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 operations pvminstances delete not found response has a 5xx status code
func (o *InternalV1OperationsPvminstancesDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 operations pvminstances delete not found response a status code equal to that given
func (o *InternalV1OperationsPvminstancesDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the internal v1 operations pvminstances delete not found response
func (o *InternalV1OperationsPvminstancesDeleteNotFound) Code() int {
	return 404
}

func (o *InternalV1OperationsPvminstancesDeleteNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/pvm-instances/{resource_crn}][%d] internalV1OperationsPvminstancesDeleteNotFound %s", 404, payload)
}

func (o *InternalV1OperationsPvminstancesDeleteNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/pvm-instances/{resource_crn}][%d] internalV1OperationsPvminstancesDeleteNotFound %s", 404, payload)
}

func (o *InternalV1OperationsPvminstancesDeleteNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1OperationsPvminstancesDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1OperationsPvminstancesDeleteGone creates a InternalV1OperationsPvminstancesDeleteGone with default headers values
func NewInternalV1OperationsPvminstancesDeleteGone() *InternalV1OperationsPvminstancesDeleteGone {
	return &InternalV1OperationsPvminstancesDeleteGone{}
}

/*
InternalV1OperationsPvminstancesDeleteGone describes a response with status code 410, with default header values.

Gone
*/
type InternalV1OperationsPvminstancesDeleteGone struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 operations pvminstances delete gone response has a 2xx status code
func (o *InternalV1OperationsPvminstancesDeleteGone) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 operations pvminstances delete gone response has a 3xx status code
func (o *InternalV1OperationsPvminstancesDeleteGone) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 operations pvminstances delete gone response has a 4xx status code
func (o *InternalV1OperationsPvminstancesDeleteGone) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 operations pvminstances delete gone response has a 5xx status code
func (o *InternalV1OperationsPvminstancesDeleteGone) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 operations pvminstances delete gone response a status code equal to that given
func (o *InternalV1OperationsPvminstancesDeleteGone) IsCode(code int) bool {
	return code == 410
}

// Code gets the status code for the internal v1 operations pvminstances delete gone response
func (o *InternalV1OperationsPvminstancesDeleteGone) Code() int {
	return 410
}

func (o *InternalV1OperationsPvminstancesDeleteGone) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/pvm-instances/{resource_crn}][%d] internalV1OperationsPvminstancesDeleteGone %s", 410, payload)
}

func (o *InternalV1OperationsPvminstancesDeleteGone) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/pvm-instances/{resource_crn}][%d] internalV1OperationsPvminstancesDeleteGone %s", 410, payload)
}

func (o *InternalV1OperationsPvminstancesDeleteGone) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1OperationsPvminstancesDeleteGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1OperationsPvminstancesDeleteTooManyRequests creates a InternalV1OperationsPvminstancesDeleteTooManyRequests with default headers values
func NewInternalV1OperationsPvminstancesDeleteTooManyRequests() *InternalV1OperationsPvminstancesDeleteTooManyRequests {
	return &InternalV1OperationsPvminstancesDeleteTooManyRequests{}
}

/*
InternalV1OperationsPvminstancesDeleteTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type InternalV1OperationsPvminstancesDeleteTooManyRequests struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 operations pvminstances delete too many requests response has a 2xx status code
func (o *InternalV1OperationsPvminstancesDeleteTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 operations pvminstances delete too many requests response has a 3xx status code
func (o *InternalV1OperationsPvminstancesDeleteTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 operations pvminstances delete too many requests response has a 4xx status code
func (o *InternalV1OperationsPvminstancesDeleteTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 operations pvminstances delete too many requests response has a 5xx status code
func (o *InternalV1OperationsPvminstancesDeleteTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 operations pvminstances delete too many requests response a status code equal to that given
func (o *InternalV1OperationsPvminstancesDeleteTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the internal v1 operations pvminstances delete too many requests response
func (o *InternalV1OperationsPvminstancesDeleteTooManyRequests) Code() int {
	return 429
}

func (o *InternalV1OperationsPvminstancesDeleteTooManyRequests) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/pvm-instances/{resource_crn}][%d] internalV1OperationsPvminstancesDeleteTooManyRequests %s", 429, payload)
}

func (o *InternalV1OperationsPvminstancesDeleteTooManyRequests) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/pvm-instances/{resource_crn}][%d] internalV1OperationsPvminstancesDeleteTooManyRequests %s", 429, payload)
}

func (o *InternalV1OperationsPvminstancesDeleteTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1OperationsPvminstancesDeleteTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1OperationsPvminstancesDeleteInternalServerError creates a InternalV1OperationsPvminstancesDeleteInternalServerError with default headers values
func NewInternalV1OperationsPvminstancesDeleteInternalServerError() *InternalV1OperationsPvminstancesDeleteInternalServerError {
	return &InternalV1OperationsPvminstancesDeleteInternalServerError{}
}

/*
InternalV1OperationsPvminstancesDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type InternalV1OperationsPvminstancesDeleteInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 operations pvminstances delete internal server error response has a 2xx status code
func (o *InternalV1OperationsPvminstancesDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 operations pvminstances delete internal server error response has a 3xx status code
func (o *InternalV1OperationsPvminstancesDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 operations pvminstances delete internal server error response has a 4xx status code
func (o *InternalV1OperationsPvminstancesDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this internal v1 operations pvminstances delete internal server error response has a 5xx status code
func (o *InternalV1OperationsPvminstancesDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this internal v1 operations pvminstances delete internal server error response a status code equal to that given
func (o *InternalV1OperationsPvminstancesDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the internal v1 operations pvminstances delete internal server error response
func (o *InternalV1OperationsPvminstancesDeleteInternalServerError) Code() int {
	return 500
}

func (o *InternalV1OperationsPvminstancesDeleteInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/pvm-instances/{resource_crn}][%d] internalV1OperationsPvminstancesDeleteInternalServerError %s", 500, payload)
}

func (o *InternalV1OperationsPvminstancesDeleteInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/pvm-instances/{resource_crn}][%d] internalV1OperationsPvminstancesDeleteInternalServerError %s", 500, payload)
}

func (o *InternalV1OperationsPvminstancesDeleteInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1OperationsPvminstancesDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
