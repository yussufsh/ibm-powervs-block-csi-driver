// Code generated by go-swagger; DO NOT EDIT.

package p_cloud_snapshots

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/IBM-Cloud/power-go-client/power/models"
)

// PcloudCloudinstancesSnapshotsGetReader is a Reader for the PcloudCloudinstancesSnapshotsGet structure.
type PcloudCloudinstancesSnapshotsGetReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *PcloudCloudinstancesSnapshotsGetReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewPcloudCloudinstancesSnapshotsGetOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewPcloudCloudinstancesSnapshotsGetBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 401:
		result := NewPcloudCloudinstancesSnapshotsGetUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewPcloudCloudinstancesSnapshotsGetForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewPcloudCloudinstancesSnapshotsGetNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewPcloudCloudinstancesSnapshotsGetInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewPcloudCloudinstancesSnapshotsGetOK creates a PcloudCloudinstancesSnapshotsGetOK with default headers values
func NewPcloudCloudinstancesSnapshotsGetOK() *PcloudCloudinstancesSnapshotsGetOK {
	return &PcloudCloudinstancesSnapshotsGetOK{}
}

/*
PcloudCloudinstancesSnapshotsGetOK describes a response with status code 200, with default header values.

OK
*/
type PcloudCloudinstancesSnapshotsGetOK struct {
	Payload *models.Snapshot
}

// IsSuccess returns true when this pcloud cloudinstances snapshots get o k response has a 2xx status code
func (o *PcloudCloudinstancesSnapshotsGetOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this pcloud cloudinstances snapshots get o k response has a 3xx status code
func (o *PcloudCloudinstancesSnapshotsGetOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances snapshots get o k response has a 4xx status code
func (o *PcloudCloudinstancesSnapshotsGetOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudinstances snapshots get o k response has a 5xx status code
func (o *PcloudCloudinstancesSnapshotsGetOK) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances snapshots get o k response a status code equal to that given
func (o *PcloudCloudinstancesSnapshotsGetOK) IsCode(code int) bool {
	return code == 200
}

func (o *PcloudCloudinstancesSnapshotsGetOK) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsGetOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsGetOK) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsGetOK  %+v", 200, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsGetOK) GetPayload() *models.Snapshot {
	return o.Payload
}

func (o *PcloudCloudinstancesSnapshotsGetOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Snapshot)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesSnapshotsGetBadRequest creates a PcloudCloudinstancesSnapshotsGetBadRequest with default headers values
func NewPcloudCloudinstancesSnapshotsGetBadRequest() *PcloudCloudinstancesSnapshotsGetBadRequest {
	return &PcloudCloudinstancesSnapshotsGetBadRequest{}
}

/*
PcloudCloudinstancesSnapshotsGetBadRequest describes a response with status code 400, with default header values.

Bad Request
*/
type PcloudCloudinstancesSnapshotsGetBadRequest struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances snapshots get bad request response has a 2xx status code
func (o *PcloudCloudinstancesSnapshotsGetBadRequest) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances snapshots get bad request response has a 3xx status code
func (o *PcloudCloudinstancesSnapshotsGetBadRequest) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances snapshots get bad request response has a 4xx status code
func (o *PcloudCloudinstancesSnapshotsGetBadRequest) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances snapshots get bad request response has a 5xx status code
func (o *PcloudCloudinstancesSnapshotsGetBadRequest) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances snapshots get bad request response a status code equal to that given
func (o *PcloudCloudinstancesSnapshotsGetBadRequest) IsCode(code int) bool {
	return code == 400
}

func (o *PcloudCloudinstancesSnapshotsGetBadRequest) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsGetBadRequest) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsGetBadRequest  %+v", 400, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsGetBadRequest) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesSnapshotsGetBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesSnapshotsGetUnauthorized creates a PcloudCloudinstancesSnapshotsGetUnauthorized with default headers values
func NewPcloudCloudinstancesSnapshotsGetUnauthorized() *PcloudCloudinstancesSnapshotsGetUnauthorized {
	return &PcloudCloudinstancesSnapshotsGetUnauthorized{}
}

/*
PcloudCloudinstancesSnapshotsGetUnauthorized describes a response with status code 401, with default header values.

Unauthorized
*/
type PcloudCloudinstancesSnapshotsGetUnauthorized struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances snapshots get unauthorized response has a 2xx status code
func (o *PcloudCloudinstancesSnapshotsGetUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances snapshots get unauthorized response has a 3xx status code
func (o *PcloudCloudinstancesSnapshotsGetUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances snapshots get unauthorized response has a 4xx status code
func (o *PcloudCloudinstancesSnapshotsGetUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances snapshots get unauthorized response has a 5xx status code
func (o *PcloudCloudinstancesSnapshotsGetUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances snapshots get unauthorized response a status code equal to that given
func (o *PcloudCloudinstancesSnapshotsGetUnauthorized) IsCode(code int) bool {
	return code == 401
}

func (o *PcloudCloudinstancesSnapshotsGetUnauthorized) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsGetUnauthorized) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsGetUnauthorized  %+v", 401, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsGetUnauthorized) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesSnapshotsGetUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesSnapshotsGetForbidden creates a PcloudCloudinstancesSnapshotsGetForbidden with default headers values
func NewPcloudCloudinstancesSnapshotsGetForbidden() *PcloudCloudinstancesSnapshotsGetForbidden {
	return &PcloudCloudinstancesSnapshotsGetForbidden{}
}

/*
PcloudCloudinstancesSnapshotsGetForbidden describes a response with status code 403, with default header values.

Forbidden
*/
type PcloudCloudinstancesSnapshotsGetForbidden struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances snapshots get forbidden response has a 2xx status code
func (o *PcloudCloudinstancesSnapshotsGetForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances snapshots get forbidden response has a 3xx status code
func (o *PcloudCloudinstancesSnapshotsGetForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances snapshots get forbidden response has a 4xx status code
func (o *PcloudCloudinstancesSnapshotsGetForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances snapshots get forbidden response has a 5xx status code
func (o *PcloudCloudinstancesSnapshotsGetForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances snapshots get forbidden response a status code equal to that given
func (o *PcloudCloudinstancesSnapshotsGetForbidden) IsCode(code int) bool {
	return code == 403
}

func (o *PcloudCloudinstancesSnapshotsGetForbidden) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsGetForbidden) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsGetForbidden  %+v", 403, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsGetForbidden) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesSnapshotsGetForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesSnapshotsGetNotFound creates a PcloudCloudinstancesSnapshotsGetNotFound with default headers values
func NewPcloudCloudinstancesSnapshotsGetNotFound() *PcloudCloudinstancesSnapshotsGetNotFound {
	return &PcloudCloudinstancesSnapshotsGetNotFound{}
}

/*
PcloudCloudinstancesSnapshotsGetNotFound describes a response with status code 404, with default header values.

Not Found
*/
type PcloudCloudinstancesSnapshotsGetNotFound struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances snapshots get not found response has a 2xx status code
func (o *PcloudCloudinstancesSnapshotsGetNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances snapshots get not found response has a 3xx status code
func (o *PcloudCloudinstancesSnapshotsGetNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances snapshots get not found response has a 4xx status code
func (o *PcloudCloudinstancesSnapshotsGetNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this pcloud cloudinstances snapshots get not found response has a 5xx status code
func (o *PcloudCloudinstancesSnapshotsGetNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this pcloud cloudinstances snapshots get not found response a status code equal to that given
func (o *PcloudCloudinstancesSnapshotsGetNotFound) IsCode(code int) bool {
	return code == 404
}

func (o *PcloudCloudinstancesSnapshotsGetNotFound) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsGetNotFound) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsGetNotFound  %+v", 404, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsGetNotFound) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesSnapshotsGetNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewPcloudCloudinstancesSnapshotsGetInternalServerError creates a PcloudCloudinstancesSnapshotsGetInternalServerError with default headers values
func NewPcloudCloudinstancesSnapshotsGetInternalServerError() *PcloudCloudinstancesSnapshotsGetInternalServerError {
	return &PcloudCloudinstancesSnapshotsGetInternalServerError{}
}

/*
PcloudCloudinstancesSnapshotsGetInternalServerError describes a response with status code 500, with default header values.

Internal Server Error
*/
type PcloudCloudinstancesSnapshotsGetInternalServerError struct {
	Payload *models.Error
}

// IsSuccess returns true when this pcloud cloudinstances snapshots get internal server error response has a 2xx status code
func (o *PcloudCloudinstancesSnapshotsGetInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this pcloud cloudinstances snapshots get internal server error response has a 3xx status code
func (o *PcloudCloudinstancesSnapshotsGetInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this pcloud cloudinstances snapshots get internal server error response has a 4xx status code
func (o *PcloudCloudinstancesSnapshotsGetInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this pcloud cloudinstances snapshots get internal server error response has a 5xx status code
func (o *PcloudCloudinstancesSnapshotsGetInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this pcloud cloudinstances snapshots get internal server error response a status code equal to that given
func (o *PcloudCloudinstancesSnapshotsGetInternalServerError) IsCode(code int) bool {
	return code == 500
}

func (o *PcloudCloudinstancesSnapshotsGetInternalServerError) Error() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsGetInternalServerError) String() string {
	return fmt.Sprintf("[GET /pcloud/v1/cloud-instances/{cloud_instance_id}/snapshots/{snapshot_id}][%d] pcloudCloudinstancesSnapshotsGetInternalServerError  %+v", 500, o.Payload)
}

func (o *PcloudCloudinstancesSnapshotsGetInternalServerError) GetPayload() *models.Error {
	return o.Payload
}

func (o *PcloudCloudinstancesSnapshotsGetInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}