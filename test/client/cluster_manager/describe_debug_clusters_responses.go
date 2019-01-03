// Code generated by go-swagger; DO NOT EDIT.

package cluster_manager

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"openpitrix.io/openpitrix/test/models"
)

// DescribeDebugClustersReader is a Reader for the DescribeDebugClusters structure.
type DescribeDebugClustersReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DescribeDebugClustersReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDescribeDebugClustersOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDescribeDebugClustersOK creates a DescribeDebugClustersOK with default headers values
func NewDescribeDebugClustersOK() *DescribeDebugClustersOK {
	return &DescribeDebugClustersOK{}
}

/*DescribeDebugClustersOK handles this case with default header values.

A successful response.
*/
type DescribeDebugClustersOK struct {
	Payload *models.OpenpitrixDescribeClustersResponse
}

func (o *DescribeDebugClustersOK) Error() string {
	return fmt.Sprintf("[GET /v1/debug_clusters][%d] describeDebugClustersOK  %+v", 200, o.Payload)
}

func (o *DescribeDebugClustersOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.OpenpitrixDescribeClustersResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
