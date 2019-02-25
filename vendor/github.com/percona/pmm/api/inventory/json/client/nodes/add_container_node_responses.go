// Code generated by go-swagger; DO NOT EDIT.

package nodes

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// AddContainerNodeReader is a Reader for the AddContainerNode structure.
type AddContainerNodeReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *AddContainerNodeReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewAddContainerNodeOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewAddContainerNodeOK creates a AddContainerNodeOK with default headers values
func NewAddContainerNodeOK() *AddContainerNodeOK {
	return &AddContainerNodeOK{}
}

/*AddContainerNodeOK handles this case with default header values.

A successful response.
*/
type AddContainerNodeOK struct {
	Payload *AddContainerNodeOKBody
}

func (o *AddContainerNodeOK) Error() string {
	return fmt.Sprintf("[POST /v1/inventory/Nodes/AddContainer][%d] addContainerNodeOK  %+v", 200, o.Payload)
}

func (o *AddContainerNodeOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(AddContainerNodeOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

/*AddContainerNodeBody add container node body
swagger:model AddContainerNodeBody
*/
type AddContainerNodeBody struct {

	// Custom user-assigned labels. Keys must start with "_".
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Docker container identifier. If specified, must be a unique Docker container identifier.
	DockerContainerID string `json:"docker_container_id,omitempty"`

	// Container name.
	DockerContainerName string `json:"docker_container_name,omitempty"`

	// Linux machine-id of the Generic Node where this Container Node runs. If defined, Generic Node with that machine_id must exist.
	MachineID string `json:"machine_id,omitempty"`

	// Unique across all Nodes user-defined name.
	NodeName string `json:"node_name,omitempty"`
}

// Validate validates this add container node body
func (o *AddContainerNodeBody) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddContainerNodeBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddContainerNodeBody) UnmarshalBinary(b []byte) error {
	var res AddContainerNodeBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddContainerNodeOKBody add container node o k body
swagger:model AddContainerNodeOKBody
*/
type AddContainerNodeOKBody struct {

	// container
	Container *AddContainerNodeOKBodyContainer `json:"container,omitempty"`
}

// Validate validates this add container node o k body
func (o *AddContainerNodeOKBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := o.validateContainer(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *AddContainerNodeOKBody) validateContainer(formats strfmt.Registry) error {

	if swag.IsZero(o.Container) { // not required
		return nil
	}

	if o.Container != nil {
		if err := o.Container.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("addContainerNodeOK" + "." + "container")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (o *AddContainerNodeOKBody) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddContainerNodeOKBody) UnmarshalBinary(b []byte) error {
	var res AddContainerNodeOKBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}

/*AddContainerNodeOKBodyContainer ContainerNode represents a Docker container.
swagger:model AddContainerNodeOKBodyContainer
*/
type AddContainerNodeOKBodyContainer struct {

	// Custom user-assigned labels. Keys must start with "_". Can be changed.
	CustomLabels map[string]string `json:"custom_labels,omitempty"`

	// Docker container identifier. If specified, must be a unique Docker container identifier. Can't be changed.
	DockerContainerID string `json:"docker_container_id,omitempty"`

	// Container name. Can be changed.
	DockerContainerName string `json:"docker_container_name,omitempty"`

	// Linux machine-id of the Generic Node where this Container Node runs. If defined, Generic Node with that machine_id must exist. Can't be changed.
	MachineID string `json:"machine_id,omitempty"`

	// Unique randomly generated instance identifier, can't be changed.
	NodeID string `json:"node_id,omitempty"`

	// Unique across all Nodes user-defined name, can be changed.
	NodeName string `json:"node_name,omitempty"`
}

// Validate validates this add container node o k body container
func (o *AddContainerNodeOKBodyContainer) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (o *AddContainerNodeOKBodyContainer) MarshalBinary() ([]byte, error) {
	if o == nil {
		return nil, nil
	}
	return swag.WriteJSON(o)
}

// UnmarshalBinary interface implementation
func (o *AddContainerNodeOKBodyContainer) UnmarshalBinary(b []byte) error {
	var res AddContainerNodeOKBodyContainer
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*o = res
	return nil
}