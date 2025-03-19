// Code generated by go-swagger; DO NOT EDIT.

package internal_operations_snapshots

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

// InternalV1OperationsSnapshotsPostReader is a Reader for the InternalV1OperationsSnapshotsPost structure.
type InternalV1OperationsSnapshotsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InternalV1OperationsSnapshotsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewInternalV1OperationsSnapshotsPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInternalV1OperationsSnapshotsPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewInternalV1OperationsSnapshotsPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewInternalV1OperationsSnapshotsPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewInternalV1OperationsSnapshotsPostTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewInternalV1OperationsSnapshotsPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /internal/v1/operations/snapshots] internal.v1.operations.snapshots.post", response, response.Code())
	}
}

// NewInternalV1OperationsSnapshotsPostCreated creates a InternalV1OperationsSnapshotsPostCreated with default headers values
func NewInternalV1OperationsSnapshotsPostCreated() *InternalV1OperationsSnapshotsPostCreated {
	return &InternalV1OperationsSnapshotsPostCreated{}
}

/*
InternalV1OperationsSnapshotsPostCreated describes a response with status code 201, with default header values.

Created
*/
type InternalV1OperationsSnapshotsPostCreated struct {
	Payload *models.InternalOperationsResponse
}

// IsSuccess returns true when this internal v1 operations snapshots post created response has a 2xx status code
func (o *InternalV1OperationsSnapshotsPostCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this internal v1 operations snapshots post created response has a 3xx status code
func (o *InternalV1OperationsSnapshotsPostCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 operations snapshots post created response has a 4xx status code
func (o *InternalV1OperationsSnapshotsPostCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this internal v1 operations snapshots post created response has a 5xx status code
func (o *InternalV1OperationsSnapshotsPostCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 operations snapshots post created response a status code equal to that given
func (o *InternalV1OperationsSnapshotsPostCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the internal v1 operations snapshots post created response
func (o *InternalV1OperationsSnapshotsPostCreated) Code() int {
	return 201
}

func (o *InternalV1OperationsSnapshotsPostCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /internal/v1/operations/snapshots][%d] internalV1OperationsSnapshotsPostCreated %s", 201, payload)
}

func (o *InternalV1OperationsSnapshotsPostCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /internal/v1/operations/snapshots][%d] internalV1OperationsSnapshotsPostCreated %s", 201, payload)
}

func (o *InternalV1OperationsSnapshotsPostCreated) GetPayload() *models.InternalOperationsResponse {
	return o.Payload
}

func (o *InternalV1OperationsSnapshotsPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalOperationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1OperationsSnapshotsPostBadRequest creates a InternalV1OperationsSnapshotsPostBadRequest with default headers values
func NewInternalV1OperationsSnapshotsPostBadRequest() *InternalV1OperationsSnapshotsPostBadRequest {
	return &InternalV1OperationsSnapshotsPostBadRequest{}
}

/*
InternalV1OperationsSnapshotsPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type InternalV1OperationsSnapshotsPostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 operations snapshots post bad request response has a 2xx status code
func (o *InternalV1OperationsSnapshotsPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 operations snapshots post bad request response has a 3xx status code
func (o *InternalV1OperationsSnapshotsPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 operations snapshots post bad request response has a 4xx status code
func (o *InternalV1OperationsSnapshotsPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 operations snapshots post bad request response has a 5xx status code
func (o *InternalV1OperationsSnapshotsPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 operations snapshots post bad request response a status code equal to that given
func (o *InternalV1OperationsSnapshotsPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the internal v1 operations snapshots post bad request response
func (o *InternalV1OperationsSnapshotsPostBadRequest) Code() int {
	return 400
}

func (o *InternalV1OperationsSnapshotsPostBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /internal/v1/operations/snapshots][%d] internalV1OperationsSnapshotsPostBadRequest %s", 400, payload)
}

func (o *InternalV1OperationsSnapshotsPostBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /internal/v1/operations/snapshots][%d] internalV1OperationsSnapshotsPostBadRequest %s", 400, payload)
}

func (o *InternalV1OperationsSnapshotsPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1OperationsSnapshotsPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1OperationsSnapshotsPostUnauthorized creates a InternalV1OperationsSnapshotsPostUnauthorized with default headers values
func NewInternalV1OperationsSnapshotsPostUnauthorized() *InternalV1OperationsSnapshotsPostUnauthorized {
	return &InternalV1OperationsSnapshotsPostUnauthorized{}
}

/*
InternalV1OperationsSnapshotsPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type InternalV1OperationsSnapshotsPostUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 operations snapshots post unauthorized response has a 2xx status code
func (o *InternalV1OperationsSnapshotsPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 operations snapshots post unauthorized response has a 3xx status code
func (o *InternalV1OperationsSnapshotsPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 operations snapshots post unauthorized response has a 4xx status code
func (o *InternalV1OperationsSnapshotsPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 operations snapshots post unauthorized response has a 5xx status code
func (o *InternalV1OperationsSnapshotsPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 operations snapshots post unauthorized response a status code equal to that given
func (o *InternalV1OperationsSnapshotsPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the internal v1 operations snapshots post unauthorized response
func (o *InternalV1OperationsSnapshotsPostUnauthorized) Code() int {
	return 401
}

func (o *InternalV1OperationsSnapshotsPostUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /internal/v1/operations/snapshots][%d] internalV1OperationsSnapshotsPostUnauthorized %s", 401, payload)
}

func (o *InternalV1OperationsSnapshotsPostUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /internal/v1/operations/snapshots][%d] internalV1OperationsSnapshotsPostUnauthorized %s", 401, payload)
}

func (o *InternalV1OperationsSnapshotsPostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1OperationsSnapshotsPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1OperationsSnapshotsPostForbidden creates a InternalV1OperationsSnapshotsPostForbidden with default headers values
func NewInternalV1OperationsSnapshotsPostForbidden() *InternalV1OperationsSnapshotsPostForbidden {
	return &InternalV1OperationsSnapshotsPostForbidden{}
}

/*
InternalV1OperationsSnapshotsPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type InternalV1OperationsSnapshotsPostForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 operations snapshots post forbidden response has a 2xx status code
func (o *InternalV1OperationsSnapshotsPostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 operations snapshots post forbidden response has a 3xx status code
func (o *InternalV1OperationsSnapshotsPostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 operations snapshots post forbidden response has a 4xx status code
func (o *InternalV1OperationsSnapshotsPostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 operations snapshots post forbidden response has a 5xx status code
func (o *InternalV1OperationsSnapshotsPostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 operations snapshots post forbidden response a status code equal to that given
func (o *InternalV1OperationsSnapshotsPostForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the internal v1 operations snapshots post forbidden response
func (o *InternalV1OperationsSnapshotsPostForbidden) Code() int {
	return 403
}

func (o *InternalV1OperationsSnapshotsPostForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /internal/v1/operations/snapshots][%d] internalV1OperationsSnapshotsPostForbidden %s", 403, payload)
}

func (o *InternalV1OperationsSnapshotsPostForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /internal/v1/operations/snapshots][%d] internalV1OperationsSnapshotsPostForbidden %s", 403, payload)
}

func (o *InternalV1OperationsSnapshotsPostForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1OperationsSnapshotsPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1OperationsSnapshotsPostTooManyRequests creates a InternalV1OperationsSnapshotsPostTooManyRequests with default headers values
func NewInternalV1OperationsSnapshotsPostTooManyRequests() *InternalV1OperationsSnapshotsPostTooManyRequests {
	return &InternalV1OperationsSnapshotsPostTooManyRequests{}
}

/*
InternalV1OperationsSnapshotsPostTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type InternalV1OperationsSnapshotsPostTooManyRequests struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 operations snapshots post too many requests response has a 2xx status code
func (o *InternalV1OperationsSnapshotsPostTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 operations snapshots post too many requests response has a 3xx status code
func (o *InternalV1OperationsSnapshotsPostTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 operations snapshots post too many requests response has a 4xx status code
func (o *InternalV1OperationsSnapshotsPostTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 operations snapshots post too many requests response has a 5xx status code
func (o *InternalV1OperationsSnapshotsPostTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 operations snapshots post too many requests response a status code equal to that given
func (o *InternalV1OperationsSnapshotsPostTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the internal v1 operations snapshots post too many requests response
func (o *InternalV1OperationsSnapshotsPostTooManyRequests) Code() int {
	return 429
}

func (o *InternalV1OperationsSnapshotsPostTooManyRequests) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /internal/v1/operations/snapshots][%d] internalV1OperationsSnapshotsPostTooManyRequests %s", 429, payload)
}

func (o *InternalV1OperationsSnapshotsPostTooManyRequests) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /internal/v1/operations/snapshots][%d] internalV1OperationsSnapshotsPostTooManyRequests %s", 429, payload)
}

func (o *InternalV1OperationsSnapshotsPostTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1OperationsSnapshotsPostTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1OperationsSnapshotsPostInternalServerError creates a InternalV1OperationsSnapshotsPostInternalServerError with default headers values
func NewInternalV1OperationsSnapshotsPostInternalServerError() *InternalV1OperationsSnapshotsPostInternalServerError {
	return &InternalV1OperationsSnapshotsPostInternalServerError{}
}

/*
InternalV1OperationsSnapshotsPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type InternalV1OperationsSnapshotsPostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 operations snapshots post internal server error response has a 2xx status code
func (o *InternalV1OperationsSnapshotsPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 operations snapshots post internal server error response has a 3xx status code
func (o *InternalV1OperationsSnapshotsPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 operations snapshots post internal server error response has a 4xx status code
func (o *InternalV1OperationsSnapshotsPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this internal v1 operations snapshots post internal server error response has a 5xx status code
func (o *InternalV1OperationsSnapshotsPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this internal v1 operations snapshots post internal server error response a status code equal to that given
func (o *InternalV1OperationsSnapshotsPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the internal v1 operations snapshots post internal server error response
func (o *InternalV1OperationsSnapshotsPostInternalServerError) Code() int {
	return 500
}

func (o *InternalV1OperationsSnapshotsPostInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /internal/v1/operations/snapshots][%d] internalV1OperationsSnapshotsPostInternalServerError %s", 500, payload)
}

func (o *InternalV1OperationsSnapshotsPostInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /internal/v1/operations/snapshots][%d] internalV1OperationsSnapshotsPostInternalServerError %s", 500, payload)
}

func (o *InternalV1OperationsSnapshotsPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1OperationsSnapshotsPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
