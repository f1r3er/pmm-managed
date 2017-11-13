// Code generated by go-swagger; DO NOT EDIT.

package scrape_configs

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/percona/pmm-managed/api/swagger/models"
)

// GetMixin2Reader is a Reader for the GetMixin2 structure.
type GetMixin2Reader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetMixin2Reader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetMixin2OK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetMixin2OK creates a GetMixin2OK with default headers values
func NewGetMixin2OK() *GetMixin2OK {
	return &GetMixin2OK{}
}

/*GetMixin2OK handles this case with default header values.

(empty)
*/
type GetMixin2OK struct {
	Payload *models.APIScrapeConfigsGetResponse
}

func (o *GetMixin2OK) Error() string {
	return fmt.Sprintf("[GET /v0/scrape-configs/{job_name}][%d] getMixin2OK  %+v", 200, o.Payload)
}

func (o *GetMixin2OK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.APIScrapeConfigsGetResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}