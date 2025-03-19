// Code generated by go-swagger; DO NOT EDIT.

package internal_operations_volumes

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

// InternalV1OperationsVolumesDeleteReader is a Reader for the InternalV1OperationsVolumesDelete structure.
type InternalV1OperationsVolumesDeleteReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InternalV1OperationsVolumesDeleteReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewInternalV1OperationsVolumesDeleteNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInternalV1OperationsVolumesDeleteBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewInternalV1OperationsVolumesDeleteUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewInternalV1OperationsVolumesDeleteForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewInternalV1OperationsVolumesDeleteNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewInternalV1OperationsVolumesDeleteGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewInternalV1OperationsVolumesDeleteTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewInternalV1OperationsVolumesDeleteInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[DELETE /internal/v1/operations/volumes/{resource_crn}] internal.v1.operations.volumes.delete", response, response.Code())
	}
}

// NewInternalV1OperationsVolumesDeleteNoContent creates a InternalV1OperationsVolumesDeleteNoContent with default headers values
func NewInternalV1OperationsVolumesDeleteNoContent() *InternalV1OperationsVolumesDeleteNoContent {
	return &InternalV1OperationsVolumesDeleteNoContent{}
}

/*
InternalV1OperationsVolumesDeleteNoContent describes a response with status code 204, with default header values.

Deleted
*/
type InternalV1OperationsVolumesDeleteNoContent struct {
}

// IsSuccess returns true when this internal v1 operations volumes delete no content response has a 2xx status code
func (o *InternalV1OperationsVolumesDeleteNoContent) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this internal v1 operations volumes delete no content response has a 3xx status code
func (o *InternalV1OperationsVolumesDeleteNoContent) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 operations volumes delete no content response has a 4xx status code
func (o *InternalV1OperationsVolumesDeleteNoContent) IsClientError() bool {
	return false
}

// IsServerError returns true when this internal v1 operations volumes delete no content response has a 5xx status code
func (o *InternalV1OperationsVolumesDeleteNoContent) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 operations volumes delete no content response a status code equal to that given
func (o *InternalV1OperationsVolumesDeleteNoContent) IsCode(code int) bool {
	return code == 204
}

// Code gets the status code for the internal v1 operations volumes delete no content response
func (o *InternalV1OperationsVolumesDeleteNoContent) Code() int {
	return 204
}

func (o *InternalV1OperationsVolumesDeleteNoContent) Error() string {
	return fmt.Sprintf("[DELETE /internal/v1/operations/volumes/{resource_crn}][%d] internalV1OperationsVolumesDeleteNoContent", 204)
}

func (o *InternalV1OperationsVolumesDeleteNoContent) String() string {
	return fmt.Sprintf("[DELETE /internal/v1/operations/volumes/{resource_crn}][%d] internalV1OperationsVolumesDeleteNoContent", 204)
}

func (o *InternalV1OperationsVolumesDeleteNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewInternalV1OperationsVolumesDeleteBadRequest creates a InternalV1OperationsVolumesDeleteBadRequest with default headers values
func NewInternalV1OperationsVolumesDeleteBadRequest() *InternalV1OperationsVolumesDeleteBadRequest {
	return &InternalV1OperationsVolumesDeleteBadRequest{}
}

/*
InternalV1OperationsVolumesDeleteBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type InternalV1OperationsVolumesDeleteBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 operations volumes delete bad request response has a 2xx status code
func (o *InternalV1OperationsVolumesDeleteBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 operations volumes delete bad request response has a 3xx status code
func (o *InternalV1OperationsVolumesDeleteBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 operations volumes delete bad request response has a 4xx status code
func (o *InternalV1OperationsVolumesDeleteBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 operations volumes delete bad request response has a 5xx status code
func (o *InternalV1OperationsVolumesDeleteBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 operations volumes delete bad request response a status code equal to that given
func (o *InternalV1OperationsVolumesDeleteBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the internal v1 operations volumes delete bad request response
func (o *InternalV1OperationsVolumesDeleteBadRequest) Code() int {
	return 400
}

func (o *InternalV1OperationsVolumesDeleteBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/volumes/{resource_crn}][%d] internalV1OperationsVolumesDeleteBadRequest %s", 400, payload)
}

func (o *InternalV1OperationsVolumesDeleteBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/volumes/{resource_crn}][%d] internalV1OperationsVolumesDeleteBadRequest %s", 400, payload)
}

func (o *InternalV1OperationsVolumesDeleteBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1OperationsVolumesDeleteBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1OperationsVolumesDeleteUnauthorized creates a InternalV1OperationsVolumesDeleteUnauthorized with default headers values
func NewInternalV1OperationsVolumesDeleteUnauthorized() *InternalV1OperationsVolumesDeleteUnauthorized {
	return &InternalV1OperationsVolumesDeleteUnauthorized{}
}

/*
InternalV1OperationsVolumesDeleteUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type InternalV1OperationsVolumesDeleteUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 operations volumes delete unauthorized response has a 2xx status code
func (o *InternalV1OperationsVolumesDeleteUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 operations volumes delete unauthorized response has a 3xx status code
func (o *InternalV1OperationsVolumesDeleteUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 operations volumes delete unauthorized response has a 4xx status code
func (o *InternalV1OperationsVolumesDeleteUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 operations volumes delete unauthorized response has a 5xx status code
func (o *InternalV1OperationsVolumesDeleteUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 operations volumes delete unauthorized response a status code equal to that given
func (o *InternalV1OperationsVolumesDeleteUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the internal v1 operations volumes delete unauthorized response
func (o *InternalV1OperationsVolumesDeleteUnauthorized) Code() int {
	return 401
}

func (o *InternalV1OperationsVolumesDeleteUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/volumes/{resource_crn}][%d] internalV1OperationsVolumesDeleteUnauthorized %s", 401, payload)
}

func (o *InternalV1OperationsVolumesDeleteUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/volumes/{resource_crn}][%d] internalV1OperationsVolumesDeleteUnauthorized %s", 401, payload)
}

func (o *InternalV1OperationsVolumesDeleteUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1OperationsVolumesDeleteUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1OperationsVolumesDeleteForbidden creates a InternalV1OperationsVolumesDeleteForbidden with default headers values
func NewInternalV1OperationsVolumesDeleteForbidden() *InternalV1OperationsVolumesDeleteForbidden {
	return &InternalV1OperationsVolumesDeleteForbidden{}
}

/*
InternalV1OperationsVolumesDeleteForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type InternalV1OperationsVolumesDeleteForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 operations volumes delete forbidden response has a 2xx status code
func (o *InternalV1OperationsVolumesDeleteForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 operations volumes delete forbidden response has a 3xx status code
func (o *InternalV1OperationsVolumesDeleteForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 operations volumes delete forbidden response has a 4xx status code
func (o *InternalV1OperationsVolumesDeleteForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 operations volumes delete forbidden response has a 5xx status code
func (o *InternalV1OperationsVolumesDeleteForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 operations volumes delete forbidden response a status code equal to that given
func (o *InternalV1OperationsVolumesDeleteForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the internal v1 operations volumes delete forbidden response
func (o *InternalV1OperationsVolumesDeleteForbidden) Code() int {
	return 403
}

func (o *InternalV1OperationsVolumesDeleteForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/volumes/{resource_crn}][%d] internalV1OperationsVolumesDeleteForbidden %s", 403, payload)
}

func (o *InternalV1OperationsVolumesDeleteForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/volumes/{resource_crn}][%d] internalV1OperationsVolumesDeleteForbidden %s", 403, payload)
}

func (o *InternalV1OperationsVolumesDeleteForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1OperationsVolumesDeleteForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1OperationsVolumesDeleteNotFound creates a InternalV1OperationsVolumesDeleteNotFound with default headers values
func NewInternalV1OperationsVolumesDeleteNotFound() *InternalV1OperationsVolumesDeleteNotFound {
	return &InternalV1OperationsVolumesDeleteNotFound{}
}

/*
InternalV1OperationsVolumesDeleteNotFound describes a response with status code 404, with default header values.

Not Found
*/
type InternalV1OperationsVolumesDeleteNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 operations volumes delete not found response has a 2xx status code
func (o *InternalV1OperationsVolumesDeleteNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 operations volumes delete not found response has a 3xx status code
func (o *InternalV1OperationsVolumesDeleteNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 operations volumes delete not found response has a 4xx status code
func (o *InternalV1OperationsVolumesDeleteNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 operations volumes delete not found response has a 5xx status code
func (o *InternalV1OperationsVolumesDeleteNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 operations volumes delete not found response a status code equal to that given
func (o *InternalV1OperationsVolumesDeleteNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the internal v1 operations volumes delete not found response
func (o *InternalV1OperationsVolumesDeleteNotFound) Code() int {
	return 404
}

func (o *InternalV1OperationsVolumesDeleteNotFound) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/volumes/{resource_crn}][%d] internalV1OperationsVolumesDeleteNotFound %s", 404, payload)
}

func (o *InternalV1OperationsVolumesDeleteNotFound) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/volumes/{resource_crn}][%d] internalV1OperationsVolumesDeleteNotFound %s", 404, payload)
}

func (o *InternalV1OperationsVolumesDeleteNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1OperationsVolumesDeleteNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1OperationsVolumesDeleteGone creates a InternalV1OperationsVolumesDeleteGone with default headers values
func NewInternalV1OperationsVolumesDeleteGone() *InternalV1OperationsVolumesDeleteGone {
	return &InternalV1OperationsVolumesDeleteGone{}
}

/*
InternalV1OperationsVolumesDeleteGone describes a response with status code 410, with default header values.

Gone
*/
type InternalV1OperationsVolumesDeleteGone struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 operations volumes delete gone response has a 2xx status code
func (o *InternalV1OperationsVolumesDeleteGone) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 operations volumes delete gone response has a 3xx status code
func (o *InternalV1OperationsVolumesDeleteGone) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 operations volumes delete gone response has a 4xx status code
func (o *InternalV1OperationsVolumesDeleteGone) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 operations volumes delete gone response has a 5xx status code
func (o *InternalV1OperationsVolumesDeleteGone) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 operations volumes delete gone response a status code equal to that given
func (o *InternalV1OperationsVolumesDeleteGone) IsCode(code int) bool {
	return code == 410
}

// Code gets the status code for the internal v1 operations volumes delete gone response
func (o *InternalV1OperationsVolumesDeleteGone) Code() int {
	return 410
}

func (o *InternalV1OperationsVolumesDeleteGone) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/volumes/{resource_crn}][%d] internalV1OperationsVolumesDeleteGone %s", 410, payload)
}

func (o *InternalV1OperationsVolumesDeleteGone) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/volumes/{resource_crn}][%d] internalV1OperationsVolumesDeleteGone %s", 410, payload)
}

func (o *InternalV1OperationsVolumesDeleteGone) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1OperationsVolumesDeleteGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1OperationsVolumesDeleteTooManyRequests creates a InternalV1OperationsVolumesDeleteTooManyRequests with default headers values
func NewInternalV1OperationsVolumesDeleteTooManyRequests() *InternalV1OperationsVolumesDeleteTooManyRequests {
	return &InternalV1OperationsVolumesDeleteTooManyRequests{}
}

/*
InternalV1OperationsVolumesDeleteTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type InternalV1OperationsVolumesDeleteTooManyRequests struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 operations volumes delete too many requests response has a 2xx status code
func (o *InternalV1OperationsVolumesDeleteTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 operations volumes delete too many requests response has a 3xx status code
func (o *InternalV1OperationsVolumesDeleteTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 operations volumes delete too many requests response has a 4xx status code
func (o *InternalV1OperationsVolumesDeleteTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 operations volumes delete too many requests response has a 5xx status code
func (o *InternalV1OperationsVolumesDeleteTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 operations volumes delete too many requests response a status code equal to that given
func (o *InternalV1OperationsVolumesDeleteTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the internal v1 operations volumes delete too many requests response
func (o *InternalV1OperationsVolumesDeleteTooManyRequests) Code() int {
	return 429
}

func (o *InternalV1OperationsVolumesDeleteTooManyRequests) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/volumes/{resource_crn}][%d] internalV1OperationsVolumesDeleteTooManyRequests %s", 429, payload)
}

func (o *InternalV1OperationsVolumesDeleteTooManyRequests) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/volumes/{resource_crn}][%d] internalV1OperationsVolumesDeleteTooManyRequests %s", 429, payload)
}

func (o *InternalV1OperationsVolumesDeleteTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1OperationsVolumesDeleteTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1OperationsVolumesDeleteInternalServerError creates a InternalV1OperationsVolumesDeleteInternalServerError with default headers values
func NewInternalV1OperationsVolumesDeleteInternalServerError() *InternalV1OperationsVolumesDeleteInternalServerError {
	return &InternalV1OperationsVolumesDeleteInternalServerError{}
}

/*
InternalV1OperationsVolumesDeleteInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type InternalV1OperationsVolumesDeleteInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 operations volumes delete internal server error response has a 2xx status code
func (o *InternalV1OperationsVolumesDeleteInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 operations volumes delete internal server error response has a 3xx status code
func (o *InternalV1OperationsVolumesDeleteInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 operations volumes delete internal server error response has a 4xx status code
func (o *InternalV1OperationsVolumesDeleteInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this internal v1 operations volumes delete internal server error response has a 5xx status code
func (o *InternalV1OperationsVolumesDeleteInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this internal v1 operations volumes delete internal server error response a status code equal to that given
func (o *InternalV1OperationsVolumesDeleteInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the internal v1 operations volumes delete internal server error response
func (o *InternalV1OperationsVolumesDeleteInternalServerError) Code() int {
	return 500
}

func (o *InternalV1OperationsVolumesDeleteInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/volumes/{resource_crn}][%d] internalV1OperationsVolumesDeleteInternalServerError %s", 500, payload)
}

func (o *InternalV1OperationsVolumesDeleteInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[DELETE /internal/v1/operations/volumes/{resource_crn}][%d] internalV1OperationsVolumesDeleteInternalServerError %s", 500, payload)
}

func (o *InternalV1OperationsVolumesDeleteInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1OperationsVolumesDeleteInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
