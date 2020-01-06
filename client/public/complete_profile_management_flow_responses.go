// Code generated by go-swagger; DO NOT EDIT.

package public

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/ory/kratos-client-go/models"
)

// CompleteProfileManagementFlowReader is a Reader for the CompleteProfileManagementFlow structure.
type CompleteProfileManagementFlowReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CompleteProfileManagementFlowReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 302:
		result := NewCompleteProfileManagementFlowFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewCompleteProfileManagementFlowInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCompleteProfileManagementFlowFound creates a CompleteProfileManagementFlowFound with default headers values
func NewCompleteProfileManagementFlowFound() *CompleteProfileManagementFlowFound {
	return &CompleteProfileManagementFlowFound{}
}

/*CompleteProfileManagementFlowFound handles this case with default header values.

Empty responses are sent when, for example, resources are deleted. The HTTP status code for empty responses is
typically 201.
*/
type CompleteProfileManagementFlowFound struct {
}

func (o *CompleteProfileManagementFlowFound) Error() string {
	return fmt.Sprintf("[POST /profiles][%d] completeProfileManagementFlowFound ", 302)
}

func (o *CompleteProfileManagementFlowFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewCompleteProfileManagementFlowInternalServerError creates a CompleteProfileManagementFlowInternalServerError with default headers values
func NewCompleteProfileManagementFlowInternalServerError() *CompleteProfileManagementFlowInternalServerError {
	return &CompleteProfileManagementFlowInternalServerError{}
}

/*CompleteProfileManagementFlowInternalServerError handles this case with default header values.

genericError
*/
type CompleteProfileManagementFlowInternalServerError struct {
	Payload *models.GenericError
}

func (o *CompleteProfileManagementFlowInternalServerError) Error() string {
	return fmt.Sprintf("[POST /profiles][%d] completeProfileManagementFlowInternalServerError  %+v", 500, o.Payload)
}

func (o *CompleteProfileManagementFlowInternalServerError) GetPayload() *models.GenericError {
	return o.Payload
}

func (o *CompleteProfileManagementFlowInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GenericError)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
