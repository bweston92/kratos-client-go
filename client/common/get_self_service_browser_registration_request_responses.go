// Code generated by go-swagger; DO NOT EDIT.

package common

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/ory/kratos-client-go/models"
)

// GetSelfServiceBrowserRegistrationRequestReader is a Reader for the GetSelfServiceBrowserRegistrationRequest structure.
type GetSelfServiceBrowserRegistrationRequestReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSelfServiceBrowserRegistrationRequestReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSelfServiceBrowserRegistrationRequestOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetSelfServiceBrowserRegistrationRequestForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetSelfServiceBrowserRegistrationRequestNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 410:
		result := NewGetSelfServiceBrowserRegistrationRequestGone()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewGetSelfServiceBrowserRegistrationRequestInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetSelfServiceBrowserRegistrationRequestOK creates a GetSelfServiceBrowserRegistrationRequestOK with default headers values
func NewGetSelfServiceBrowserRegistrationRequestOK() *GetSelfServiceBrowserRegistrationRequestOK {
	return &GetSelfServiceBrowserRegistrationRequestOK{}
}

/*GetSelfServiceBrowserRegistrationRequestOK handles this case with default header values.

registrationRequest
*/
type GetSelfServiceBrowserRegistrationRequestOK struct {
	Payload *models.RegistrationRequest
}

func (o *GetSelfServiceBrowserRegistrationRequestOK) Error() string {
	return fmt.Sprintf("[GET /self-service/browser/flows/requests/registration][%d] getSelfServiceBrowserRegistrationRequestOK  %+v", 200, o.Payload)
}

func (o *GetSelfServiceBrowserRegistrationRequestOK) GetPayload() *models.RegistrationRequest {
	return o.Payload
}

func (o *GetSelfServiceBrowserRegistrationRequestOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RegistrationRequest)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSelfServiceBrowserRegistrationRequestForbidden creates a GetSelfServiceBrowserRegistrationRequestForbidden with default headers values
func NewGetSelfServiceBrowserRegistrationRequestForbidden() *GetSelfServiceBrowserRegistrationRequestForbidden {
	return &GetSelfServiceBrowserRegistrationRequestForbidden{}
}

/*GetSelfServiceBrowserRegistrationRequestForbidden handles this case with default header values.

genericError
*/
type GetSelfServiceBrowserRegistrationRequestForbidden struct {
	Payload *models.GenericError
}

func (o *GetSelfServiceBrowserRegistrationRequestForbidden) Error() string {
	return fmt.Sprintf("[GET /self-service/browser/flows/requests/registration][%d] getSelfServiceBrowserRegistrationRequestForbidden  %+v", 403, o.Payload)
}

func (o *GetSelfServiceBrowserRegistrationRequestForbidden) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetSelfServiceBrowserRegistrationRequestForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSelfServiceBrowserRegistrationRequestNotFound creates a GetSelfServiceBrowserRegistrationRequestNotFound with default headers values
func NewGetSelfServiceBrowserRegistrationRequestNotFound() *GetSelfServiceBrowserRegistrationRequestNotFound {
	return &GetSelfServiceBrowserRegistrationRequestNotFound{}
}

/*GetSelfServiceBrowserRegistrationRequestNotFound handles this case with default header values.

genericError
*/
type GetSelfServiceBrowserRegistrationRequestNotFound struct {
	Payload *models.GenericError
}

func (o *GetSelfServiceBrowserRegistrationRequestNotFound) Error() string {
	return fmt.Sprintf("[GET /self-service/browser/flows/requests/registration][%d] getSelfServiceBrowserRegistrationRequestNotFound  %+v", 404, o.Payload)
}

func (o *GetSelfServiceBrowserRegistrationRequestNotFound) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetSelfServiceBrowserRegistrationRequestNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSelfServiceBrowserRegistrationRequestGone creates a GetSelfServiceBrowserRegistrationRequestGone with default headers values
func NewGetSelfServiceBrowserRegistrationRequestGone() *GetSelfServiceBrowserRegistrationRequestGone {
	return &GetSelfServiceBrowserRegistrationRequestGone{}
}

/*GetSelfServiceBrowserRegistrationRequestGone handles this case with default header values.

genericError
*/
type GetSelfServiceBrowserRegistrationRequestGone struct {
	Payload *models.GenericError
}

func (o *GetSelfServiceBrowserRegistrationRequestGone) Error() string {
	return fmt.Sprintf("[GET /self-service/browser/flows/requests/registration][%d] getSelfServiceBrowserRegistrationRequestGone  %+v", 410, o.Payload)
}

func (o *GetSelfServiceBrowserRegistrationRequestGone) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetSelfServiceBrowserRegistrationRequestGone) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSelfServiceBrowserRegistrationRequestInternalServerError creates a GetSelfServiceBrowserRegistrationRequestInternalServerError with default headers values
func NewGetSelfServiceBrowserRegistrationRequestInternalServerError() *GetSelfServiceBrowserRegistrationRequestInternalServerError {
	return &GetSelfServiceBrowserRegistrationRequestInternalServerError{}
}

/*GetSelfServiceBrowserRegistrationRequestInternalServerError handles this case with default header values.

genericError
*/
type GetSelfServiceBrowserRegistrationRequestInternalServerError struct {
	Payload *models.GenericError
}

func (o *GetSelfServiceBrowserRegistrationRequestInternalServerError) Error() string {
	return fmt.Sprintf("[GET /self-service/browser/flows/requests/registration][%d] getSelfServiceBrowserRegistrationRequestInternalServerError  %+v", 500, o.Payload)
}

func (o *GetSelfServiceBrowserRegistrationRequestInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *GetSelfServiceBrowserRegistrationRequestInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}