// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/netbox-community/go-netbox/netbox/models"
)

// DcimRearPortsPathsReader is a Reader for the DcimRearPortsPaths structure.
type DcimRearPortsPathsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRearPortsPathsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewDcimRearPortsPathsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	default:
		return nil, runtime.NewAPIError("response status code does not match any response statuses defined for this endpoint in the swagger spec", response, response.Code())
	}
}

// NewDcimRearPortsPathsOK creates a DcimRearPortsPathsOK with default headers values
func NewDcimRearPortsPathsOK() *DcimRearPortsPathsOK {
	return &DcimRearPortsPathsOK{}
}

/* DcimRearPortsPathsOK describes a response with status code 200, with default header values.

DcimRearPortsPathsOK dcim rear ports paths o k
*/
type DcimRearPortsPathsOK struct {
	Payload *models.RearPort
}

func (o *DcimRearPortsPathsOK) Error() string {
	return fmt.Sprintf("[GET /dcim/rear-ports/{id}/paths/][%d] dcimRearPortsPathsOK  %+v", 200, o.Payload)
}
func (o *DcimRearPortsPathsOK) GetPayload() *models.RearPort {
	return o.Payload
}

func (o *DcimRearPortsPathsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RearPort)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
