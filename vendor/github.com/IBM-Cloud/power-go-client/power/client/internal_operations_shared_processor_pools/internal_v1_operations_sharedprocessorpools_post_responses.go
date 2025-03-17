// Code generated by go-swagger; DO NOT EDIT.

package internal_operations_shared_processor_pools

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

// InternalV1OperationsSharedprocessorpoolsPostReader is a Reader for the InternalV1OperationsSharedprocessorpoolsPost structure.
type InternalV1OperationsSharedprocessorpoolsPostReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *InternalV1OperationsSharedprocessorpoolsPostReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 201:
		result := NewInternalV1OperationsSharedprocessorpoolsPostCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewInternalV1OperationsSharedprocessorpoolsPostBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewInternalV1OperationsSharedprocessorpoolsPostUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewInternalV1OperationsSharedprocessorpoolsPostForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 429:
		result := NewInternalV1OperationsSharedprocessorpoolsPostTooManyRequests()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewInternalV1OperationsSharedprocessorpoolsPostInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("[POST /internal/v1/operations/shared-processor-pools] internal.v1.operations.sharedprocessorpools.post", response, response.Code())
	}
}

// NewInternalV1OperationsSharedprocessorpoolsPostCreated creates a InternalV1OperationsSharedprocessorpoolsPostCreated with default headers values
func NewInternalV1OperationsSharedprocessorpoolsPostCreated() *InternalV1OperationsSharedprocessorpoolsPostCreated {
	return &InternalV1OperationsSharedprocessorpoolsPostCreated{}
}

/*
InternalV1OperationsSharedprocessorpoolsPostCreated describes a response with status code 201, with default header values.

Created
*/
type InternalV1OperationsSharedprocessorpoolsPostCreated struct {
	Payload *models.InternalOperationsResponse
}

// IsSuccess returns true when this internal v1 operations sharedprocessorpools post created response has a 2xx status code
func (o *InternalV1OperationsSharedprocessorpoolsPostCreated) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this internal v1 operations sharedprocessorpools post created response has a 3xx status code
func (o *InternalV1OperationsSharedprocessorpoolsPostCreated) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 operations sharedprocessorpools post created response has a 4xx status code
func (o *InternalV1OperationsSharedprocessorpoolsPostCreated) IsClientError() bool {
	return false
}

// IsServerError returns true when this internal v1 operations sharedprocessorpools post created response has a 5xx status code
func (o *InternalV1OperationsSharedprocessorpoolsPostCreated) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 operations sharedprocessorpools post created response a status code equal to that given
func (o *InternalV1OperationsSharedprocessorpoolsPostCreated) IsCode(code int) bool {
	return code == 201
}

// Code gets the status code for the internal v1 operations sharedprocessorpools post created response
func (o *InternalV1OperationsSharedprocessorpoolsPostCreated) Code() int {
	return 201
}

func (o *InternalV1OperationsSharedprocessorpoolsPostCreated) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /internal/v1/operations/shared-processor-pools][%d] internalV1OperationsSharedprocessorpoolsPostCreated %s", 201, payload)
}

func (o *InternalV1OperationsSharedprocessorpoolsPostCreated) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /internal/v1/operations/shared-processor-pools][%d] internalV1OperationsSharedprocessorpoolsPostCreated %s", 201, payload)
}

func (o *InternalV1OperationsSharedprocessorpoolsPostCreated) GetPayload() *models.InternalOperationsResponse {
	return o.Payload
}

func (o *InternalV1OperationsSharedprocessorpoolsPostCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InternalOperationsResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1OperationsSharedprocessorpoolsPostBadRequest creates a InternalV1OperationsSharedprocessorpoolsPostBadRequest with default headers values
func NewInternalV1OperationsSharedprocessorpoolsPostBadRequest() *InternalV1OperationsSharedprocessorpoolsPostBadRequest {
	return &InternalV1OperationsSharedprocessorpoolsPostBadRequest{}
}

/*
InternalV1OperationsSharedprocessorpoolsPostBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type InternalV1OperationsSharedprocessorpoolsPostBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 operations sharedprocessorpools post bad request response has a 2xx status code
func (o *InternalV1OperationsSharedprocessorpoolsPostBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 operations sharedprocessorpools post bad request response has a 3xx status code
func (o *InternalV1OperationsSharedprocessorpoolsPostBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 operations sharedprocessorpools post bad request response has a 4xx status code
func (o *InternalV1OperationsSharedprocessorpoolsPostBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 operations sharedprocessorpools post bad request response has a 5xx status code
func (o *InternalV1OperationsSharedprocessorpoolsPostBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 operations sharedprocessorpools post bad request response a status code equal to that given
func (o *InternalV1OperationsSharedprocessorpoolsPostBadRequest) IsCode(code int) bool {
	return code == 400
}

// Code gets the status code for the internal v1 operations sharedprocessorpools post bad request response
func (o *InternalV1OperationsSharedprocessorpoolsPostBadRequest) Code() int {
	return 400
}

func (o *InternalV1OperationsSharedprocessorpoolsPostBadRequest) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /internal/v1/operations/shared-processor-pools][%d] internalV1OperationsSharedprocessorpoolsPostBadRequest %s", 400, payload)
}

func (o *InternalV1OperationsSharedprocessorpoolsPostBadRequest) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /internal/v1/operations/shared-processor-pools][%d] internalV1OperationsSharedprocessorpoolsPostBadRequest %s", 400, payload)
}

func (o *InternalV1OperationsSharedprocessorpoolsPostBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1OperationsSharedprocessorpoolsPostBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1OperationsSharedprocessorpoolsPostUnauthorized creates a InternalV1OperationsSharedprocessorpoolsPostUnauthorized with default headers values
func NewInternalV1OperationsSharedprocessorpoolsPostUnauthorized() *InternalV1OperationsSharedprocessorpoolsPostUnauthorized {
	return &InternalV1OperationsSharedprocessorpoolsPostUnauthorized{}
}

/*
InternalV1OperationsSharedprocessorpoolsPostUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type InternalV1OperationsSharedprocessorpoolsPostUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 operations sharedprocessorpools post unauthorized response has a 2xx status code
func (o *InternalV1OperationsSharedprocessorpoolsPostUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 operations sharedprocessorpools post unauthorized response has a 3xx status code
func (o *InternalV1OperationsSharedprocessorpoolsPostUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 operations sharedprocessorpools post unauthorized response has a 4xx status code
func (o *InternalV1OperationsSharedprocessorpoolsPostUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 operations sharedprocessorpools post unauthorized response has a 5xx status code
func (o *InternalV1OperationsSharedprocessorpoolsPostUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 operations sharedprocessorpools post unauthorized response a status code equal to that given
func (o *InternalV1OperationsSharedprocessorpoolsPostUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the internal v1 operations sharedprocessorpools post unauthorized response
func (o *InternalV1OperationsSharedprocessorpoolsPostUnauthorized) Code() int {
	return 401
}

func (o *InternalV1OperationsSharedprocessorpoolsPostUnauthorized) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /internal/v1/operations/shared-processor-pools][%d] internalV1OperationsSharedprocessorpoolsPostUnauthorized %s", 401, payload)
}

func (o *InternalV1OperationsSharedprocessorpoolsPostUnauthorized) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /internal/v1/operations/shared-processor-pools][%d] internalV1OperationsSharedprocessorpoolsPostUnauthorized %s", 401, payload)
}

func (o *InternalV1OperationsSharedprocessorpoolsPostUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1OperationsSharedprocessorpoolsPostUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1OperationsSharedprocessorpoolsPostForbidden creates a InternalV1OperationsSharedprocessorpoolsPostForbidden with default headers values
func NewInternalV1OperationsSharedprocessorpoolsPostForbidden() *InternalV1OperationsSharedprocessorpoolsPostForbidden {
	return &InternalV1OperationsSharedprocessorpoolsPostForbidden{}
}

/*
InternalV1OperationsSharedprocessorpoolsPostForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type InternalV1OperationsSharedprocessorpoolsPostForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 operations sharedprocessorpools post forbidden response has a 2xx status code
func (o *InternalV1OperationsSharedprocessorpoolsPostForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 operations sharedprocessorpools post forbidden response has a 3xx status code
func (o *InternalV1OperationsSharedprocessorpoolsPostForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 operations sharedprocessorpools post forbidden response has a 4xx status code
func (o *InternalV1OperationsSharedprocessorpoolsPostForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 operations sharedprocessorpools post forbidden response has a 5xx status code
func (o *InternalV1OperationsSharedprocessorpoolsPostForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 operations sharedprocessorpools post forbidden response a status code equal to that given
func (o *InternalV1OperationsSharedprocessorpoolsPostForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the internal v1 operations sharedprocessorpools post forbidden response
func (o *InternalV1OperationsSharedprocessorpoolsPostForbidden) Code() int {
	return 403
}

func (o *InternalV1OperationsSharedprocessorpoolsPostForbidden) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /internal/v1/operations/shared-processor-pools][%d] internalV1OperationsSharedprocessorpoolsPostForbidden %s", 403, payload)
}

func (o *InternalV1OperationsSharedprocessorpoolsPostForbidden) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /internal/v1/operations/shared-processor-pools][%d] internalV1OperationsSharedprocessorpoolsPostForbidden %s", 403, payload)
}

func (o *InternalV1OperationsSharedprocessorpoolsPostForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1OperationsSharedprocessorpoolsPostForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1OperationsSharedprocessorpoolsPostTooManyRequests creates a InternalV1OperationsSharedprocessorpoolsPostTooManyRequests with default headers values
func NewInternalV1OperationsSharedprocessorpoolsPostTooManyRequests() *InternalV1OperationsSharedprocessorpoolsPostTooManyRequests {
	return &InternalV1OperationsSharedprocessorpoolsPostTooManyRequests{}
}

/*
InternalV1OperationsSharedprocessorpoolsPostTooManyRequests describes a response with status code 429, with default header values.

Too Many Requests
*/
type InternalV1OperationsSharedprocessorpoolsPostTooManyRequests struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 operations sharedprocessorpools post too many requests response has a 2xx status code
func (o *InternalV1OperationsSharedprocessorpoolsPostTooManyRequests) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 operations sharedprocessorpools post too many requests response has a 3xx status code
func (o *InternalV1OperationsSharedprocessorpoolsPostTooManyRequests) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 operations sharedprocessorpools post too many requests response has a 4xx status code
func (o *InternalV1OperationsSharedprocessorpoolsPostTooManyRequests) IsClientError() bool {
	return true
}

// IsServerError returns true when this internal v1 operations sharedprocessorpools post too many requests response has a 5xx status code
func (o *InternalV1OperationsSharedprocessorpoolsPostTooManyRequests) IsServerError() bool {
	return false
}

// IsCode returns true when this internal v1 operations sharedprocessorpools post too many requests response a status code equal to that given
func (o *InternalV1OperationsSharedprocessorpoolsPostTooManyRequests) IsCode(code int) bool {
	return code == 429
}

// Code gets the status code for the internal v1 operations sharedprocessorpools post too many requests response
func (o *InternalV1OperationsSharedprocessorpoolsPostTooManyRequests) Code() int {
	return 429
}

func (o *InternalV1OperationsSharedprocessorpoolsPostTooManyRequests) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /internal/v1/operations/shared-processor-pools][%d] internalV1OperationsSharedprocessorpoolsPostTooManyRequests %s", 429, payload)
}

func (o *InternalV1OperationsSharedprocessorpoolsPostTooManyRequests) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /internal/v1/operations/shared-processor-pools][%d] internalV1OperationsSharedprocessorpoolsPostTooManyRequests %s", 429, payload)
}

func (o *InternalV1OperationsSharedprocessorpoolsPostTooManyRequests) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1OperationsSharedprocessorpoolsPostTooManyRequests) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewInternalV1OperationsSharedprocessorpoolsPostInternalServerError creates a InternalV1OperationsSharedprocessorpoolsPostInternalServerError with default headers values
func NewInternalV1OperationsSharedprocessorpoolsPostInternalServerError() *InternalV1OperationsSharedprocessorpoolsPostInternalServerError {
	return &InternalV1OperationsSharedprocessorpoolsPostInternalServerError{}
}

/*
InternalV1OperationsSharedprocessorpoolsPostInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type InternalV1OperationsSharedprocessorpoolsPostInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this internal v1 operations sharedprocessorpools post internal server error response has a 2xx status code
func (o *InternalV1OperationsSharedprocessorpoolsPostInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this internal v1 operations sharedprocessorpools post internal server error response has a 3xx status code
func (o *InternalV1OperationsSharedprocessorpoolsPostInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this internal v1 operations sharedprocessorpools post internal server error response has a 4xx status code
func (o *InternalV1OperationsSharedprocessorpoolsPostInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this internal v1 operations sharedprocessorpools post internal server error response has a 5xx status code
func (o *InternalV1OperationsSharedprocessorpoolsPostInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this internal v1 operations sharedprocessorpools post internal server error response a status code equal to that given
func (o *InternalV1OperationsSharedprocessorpoolsPostInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the internal v1 operations sharedprocessorpools post internal server error response
func (o *InternalV1OperationsSharedprocessorpoolsPostInternalServerError) Code() int {
	return 500
}

func (o *InternalV1OperationsSharedprocessorpoolsPostInternalServerError) Error() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /internal/v1/operations/shared-processor-pools][%d] internalV1OperationsSharedprocessorpoolsPostInternalServerError %s", 500, payload)
}

func (o *InternalV1OperationsSharedprocessorpoolsPostInternalServerError) String() string {
	payload, _ := json.Marshal(o.Payload)
	return fmt.Sprintf("[POST /internal/v1/operations/shared-processor-pools][%d] internalV1OperationsSharedprocessorpoolsPostInternalServerError %s", 500, payload)
}

func (o *InternalV1OperationsSharedprocessorpoolsPostInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *InternalV1OperationsSharedprocessorpoolsPostInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
