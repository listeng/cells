// Code generated by go-swagger; DO NOT EDIT.

package provisioning

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/pydio/pydio-sdk-go/models"
)

// CreatePeopleReader is a Reader for the CreatePeople structure.
type CreatePeopleReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *CreatePeopleReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewCreatePeopleOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewCreatePeopleOK creates a CreatePeopleOK with default headers values
func NewCreatePeopleOK() *CreatePeopleOK {
	return &CreatePeopleOK{}
}

/*CreatePeopleOK handles this case with default header values.

Successful response
*/
type CreatePeopleOK struct {
	Payload *models.PydioResponse
}

func (o *CreatePeopleOK) Error() string {
	return fmt.Sprintf("[POST /admin/people/{path}][%d] createPeopleOK  %+v", 200, o.Payload)
}

func (o *CreatePeopleOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.PydioResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}