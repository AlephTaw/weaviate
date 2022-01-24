//                           _       _
// __      _____  __ ___   ___  __ _| |_ ___
// \ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
//  \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
//   \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
//
//  Copyright © 2016 - 2022 SeMI Technologies B.V. All rights reserved.
//
//  CONTACT: hello@semi.technology
//

// Code generated by go-swagger; DO NOT EDIT.

package objects

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/semi-technologies/weaviate/entities/models"
)

// ObjectsReferencesUpdateReader is a Reader for the ObjectsReferencesUpdate structure.
type ObjectsReferencesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ObjectsReferencesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewObjectsReferencesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewObjectsReferencesUpdateUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewObjectsReferencesUpdateForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 422:
		result := NewObjectsReferencesUpdateUnprocessableEntity()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewObjectsReferencesUpdateInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewObjectsReferencesUpdateOK creates a ObjectsReferencesUpdateOK with default headers values
func NewObjectsReferencesUpdateOK() *ObjectsReferencesUpdateOK {
	return &ObjectsReferencesUpdateOK{}
}

/*ObjectsReferencesUpdateOK handles this case with default header values.

Successfully replaced all the references.
*/
type ObjectsReferencesUpdateOK struct {
}

func (o *ObjectsReferencesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /objects/{id}/references/{propertyName}][%d] objectsReferencesUpdateOK ", 200)
}

func (o *ObjectsReferencesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewObjectsReferencesUpdateUnauthorized creates a ObjectsReferencesUpdateUnauthorized with default headers values
func NewObjectsReferencesUpdateUnauthorized() *ObjectsReferencesUpdateUnauthorized {
	return &ObjectsReferencesUpdateUnauthorized{}
}

/*ObjectsReferencesUpdateUnauthorized handles this case with default header values.

Unauthorized or invalid credentials.
*/
type ObjectsReferencesUpdateUnauthorized struct {
}

func (o *ObjectsReferencesUpdateUnauthorized) Error() string {
	return fmt.Sprintf("[PUT /objects/{id}/references/{propertyName}][%d] objectsReferencesUpdateUnauthorized ", 401)
}

func (o *ObjectsReferencesUpdateUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewObjectsReferencesUpdateForbidden creates a ObjectsReferencesUpdateForbidden with default headers values
func NewObjectsReferencesUpdateForbidden() *ObjectsReferencesUpdateForbidden {
	return &ObjectsReferencesUpdateForbidden{}
}

/*ObjectsReferencesUpdateForbidden handles this case with default header values.

Forbidden
*/
type ObjectsReferencesUpdateForbidden struct {
	Payload *models.ErrorResponse
}

func (o *ObjectsReferencesUpdateForbidden) Error() string {
	return fmt.Sprintf("[PUT /objects/{id}/references/{propertyName}][%d] objectsReferencesUpdateForbidden  %+v", 403, o.Payload)
}

func (o *ObjectsReferencesUpdateForbidden) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ObjectsReferencesUpdateForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewObjectsReferencesUpdateUnprocessableEntity creates a ObjectsReferencesUpdateUnprocessableEntity with default headers values
func NewObjectsReferencesUpdateUnprocessableEntity() *ObjectsReferencesUpdateUnprocessableEntity {
	return &ObjectsReferencesUpdateUnprocessableEntity{}
}

/*ObjectsReferencesUpdateUnprocessableEntity handles this case with default header values.

Request body is well-formed (i.e., syntactically correct), but semantically erroneous. Are you sure the property exists or that it is a class?
*/
type ObjectsReferencesUpdateUnprocessableEntity struct {
	Payload *models.ErrorResponse
}

func (o *ObjectsReferencesUpdateUnprocessableEntity) Error() string {
	return fmt.Sprintf("[PUT /objects/{id}/references/{propertyName}][%d] objectsReferencesUpdateUnprocessableEntity  %+v", 422, o.Payload)
}

func (o *ObjectsReferencesUpdateUnprocessableEntity) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ObjectsReferencesUpdateUnprocessableEntity) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewObjectsReferencesUpdateInternalServerError creates a ObjectsReferencesUpdateInternalServerError with default headers values
func NewObjectsReferencesUpdateInternalServerError() *ObjectsReferencesUpdateInternalServerError {
	return &ObjectsReferencesUpdateInternalServerError{}
}

/*ObjectsReferencesUpdateInternalServerError handles this case with default header values.

An error has occurred while trying to fulfill the request. Most likely the ErrorResponse will contain more information about the error.
*/
type ObjectsReferencesUpdateInternalServerError struct {
	Payload *models.ErrorResponse
}

func (o *ObjectsReferencesUpdateInternalServerError) Error() string {
	return fmt.Sprintf("[PUT /objects/{id}/references/{propertyName}][%d] objectsReferencesUpdateInternalServerError  %+v", 500, o.Payload)
}

func (o *ObjectsReferencesUpdateInternalServerError) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *ObjectsReferencesUpdateInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
