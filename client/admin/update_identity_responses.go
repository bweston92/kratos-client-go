// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/kratos-client-go/models"
)

// UpdateIdentityReader is a Reader for the UpdateIdentity structure.
type UpdateIdentityReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateIdentityReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewUpdateIdentityOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateIdentityBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateIdentityNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewUpdateIdentityInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewUpdateIdentityOK creates a UpdateIdentityOK with default headers values
func NewUpdateIdentityOK() *UpdateIdentityOK {
	return &UpdateIdentityOK{}
}

/* UpdateIdentityOK describes a response with status code 200, with default header values.

A single identity.
*/
type UpdateIdentityOK struct {
	Payload *models.Identity
}

func (o *UpdateIdentityOK) Error() string {
	return fmt.Sprintf("[PUT /identities/{id}][%d] updateIdentityOK  %+v", 200, o.Payload)
}
func (o *UpdateIdentityOK) GetPayload() *models.Identity {
	return o.Payload
}

func (o *UpdateIdentityOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Identity)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIdentityBadRequest creates a UpdateIdentityBadRequest with default headers values
func NewUpdateIdentityBadRequest() *UpdateIdentityBadRequest {
	return &UpdateIdentityBadRequest{}
}

/* UpdateIdentityBadRequest describes a response with status code 400, with default header values.

genericError
*/
type UpdateIdentityBadRequest struct {
	Payload *models.GenericError
}

func (o *UpdateIdentityBadRequest) Error() string {
	return fmt.Sprintf("[PUT /identities/{id}][%d] updateIdentityBadRequest  %+v", 400, o.Payload)
}
func (o *UpdateIdentityBadRequest) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *UpdateIdentityBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIdentityNotFound creates a UpdateIdentityNotFound with default headers values
func NewUpdateIdentityNotFound() *UpdateIdentityNotFound {
	return &UpdateIdentityNotFound{}
}

/* UpdateIdentityNotFound describes a response with status code 404, with default header values.

genericError
*/
type UpdateIdentityNotFound struct {
	Payload *models.GenericError
}

func (o *UpdateIdentityNotFound) Error() string {
	return fmt.Sprintf("[PUT /identities/{id}][%d] updateIdentityNotFound  %+v", 404, o.Payload)
}
func (o *UpdateIdentityNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *UpdateIdentityNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateIdentityInternalServerError creates a UpdateIdentityInternalServerError with default headers values
func NewUpdateIdentityInternalServerError() *UpdateIdentityInternalServerError {
	return &UpdateIdentityInternalServerError{}
}

/* UpdateIdentityInternalServerError describes a response with status code 500, with default header values.

genericError
*/
type UpdateIdentityInternalServerError struct {
	Payload *models.GenericError
}

func (o *UpdateIdentityInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /identities/{id}][%d] updateIdentityInternalServerError  %+v", 500, o.Payload)
}
func (o *UpdateIdentityInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *UpdateIdentityInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
