// Code generated by go-swagger; DO NOT EDIT.

package edge_application_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"
	"github.com/go-openapi/strfmt"

	"github.com/zededa/terraform-provider/models"
)

// EdgeApplicationConfigurationGetEdgeApplicationBundleByNameReader is a Reader for the EdgeApplicationConfigurationGetEdgeApplicationBundleByName structure.
type EdgeApplicationConfigurationGetEdgeApplicationBundleByNameReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewEdgeApplicationConfigurationGetEdgeApplicationBundleByNameOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewEdgeApplicationConfigurationGetEdgeApplicationBundleByNameUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewEdgeApplicationConfigurationGetEdgeApplicationBundleByNameForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewEdgeApplicationConfigurationGetEdgeApplicationBundleByNameNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 500:
		result := NewEdgeApplicationConfigurationGetEdgeApplicationBundleByNameInternalServerError()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 504:
		result := NewEdgeApplicationConfigurationGetEdgeApplicationBundleByNameGatewayTimeout()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewEdgeApplicationConfigurationGetEdgeApplicationBundleByNameDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewEdgeApplicationConfigurationGetEdgeApplicationBundleByNameOK creates a EdgeApplicationConfigurationGetEdgeApplicationBundleByNameOK with default headers values
func NewEdgeApplicationConfigurationGetEdgeApplicationBundleByNameOK() *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameOK {
	return &EdgeApplicationConfigurationGetEdgeApplicationBundleByNameOK{}
}

/*
EdgeApplicationConfigurationGetEdgeApplicationBundleByNameOK describes a response with status code 200, with default header values.

A successful response.
*/
type EdgeApplicationConfigurationGetEdgeApplicationBundleByNameOK struct {
	Payload *models.App
}

// IsSuccess returns true when this edge application configuration get edge application bundle by name o k response has a 2xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameOK) IsSuccess() bool {
	return true
}

// IsRedirect returns true when this edge application configuration get edge application bundle by name o k response has a 3xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameOK) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration get edge application bundle by name o k response has a 4xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameOK) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application configuration get edge application bundle by name o k response has a 5xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameOK) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application configuration get edge application bundle by name o k response a status code equal to that given
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameOK) IsCode(code int) bool {
	return code == 200
}

// Code gets the status code for the edge application configuration get edge application bundle by name o k response
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameOK) Code() int {
	return 200
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameOK) Error() string {
	return fmt.Sprintf("[GET /v1/apps/name/{name}][%d] edgeApplicationConfigurationGetEdgeApplicationBundleByNameOK  %+v", 200, o.Payload)
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameOK) String() string {
	return fmt.Sprintf("[GET /v1/apps/name/{name}][%d] edgeApplicationConfigurationGetEdgeApplicationBundleByNameOK  %+v", 200, o.Payload)
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameOK) GetPayload() *models.App {
	return o.Payload
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.App)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationGetEdgeApplicationBundleByNameUnauthorized creates a EdgeApplicationConfigurationGetEdgeApplicationBundleByNameUnauthorized with default headers values
func NewEdgeApplicationConfigurationGetEdgeApplicationBundleByNameUnauthorized() *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameUnauthorized {
	return &EdgeApplicationConfigurationGetEdgeApplicationBundleByNameUnauthorized{}
}

/*
EdgeApplicationConfigurationGetEdgeApplicationBundleByNameUnauthorized describes a response with status code 401, with default header values.

Unauthorized. The API gateway did not process the request because it lacks valid authentication credentials for the target resource. The request header has either no authorization details or an authorization that has been refused.
*/
type EdgeApplicationConfigurationGetEdgeApplicationBundleByNameUnauthorized struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration get edge application bundle by name unauthorized response has a 2xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameUnauthorized) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration get edge application bundle by name unauthorized response has a 3xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameUnauthorized) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration get edge application bundle by name unauthorized response has a 4xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameUnauthorized) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application configuration get edge application bundle by name unauthorized response has a 5xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameUnauthorized) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application configuration get edge application bundle by name unauthorized response a status code equal to that given
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameUnauthorized) IsCode(code int) bool {
	return code == 401
}

// Code gets the status code for the edge application configuration get edge application bundle by name unauthorized response
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameUnauthorized) Code() int {
	return 401
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameUnauthorized) Error() string {
	return fmt.Sprintf("[GET /v1/apps/name/{name}][%d] edgeApplicationConfigurationGetEdgeApplicationBundleByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameUnauthorized) String() string {
	return fmt.Sprintf("[GET /v1/apps/name/{name}][%d] edgeApplicationConfigurationGetEdgeApplicationBundleByNameUnauthorized  %+v", 401, o.Payload)
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameUnauthorized) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationGetEdgeApplicationBundleByNameForbidden creates a EdgeApplicationConfigurationGetEdgeApplicationBundleByNameForbidden with default headers values
func NewEdgeApplicationConfigurationGetEdgeApplicationBundleByNameForbidden() *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameForbidden {
	return &EdgeApplicationConfigurationGetEdgeApplicationBundleByNameForbidden{}
}

/*
EdgeApplicationConfigurationGetEdgeApplicationBundleByNameForbidden describes a response with status code 403, with default header values.

Forbidden. The API gateway did not process the request because the requestor does not have application level access permission for the operation or does not have access scope to the project.
*/
type EdgeApplicationConfigurationGetEdgeApplicationBundleByNameForbidden struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration get edge application bundle by name forbidden response has a 2xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameForbidden) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration get edge application bundle by name forbidden response has a 3xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameForbidden) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration get edge application bundle by name forbidden response has a 4xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameForbidden) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application configuration get edge application bundle by name forbidden response has a 5xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameForbidden) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application configuration get edge application bundle by name forbidden response a status code equal to that given
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameForbidden) IsCode(code int) bool {
	return code == 403
}

// Code gets the status code for the edge application configuration get edge application bundle by name forbidden response
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameForbidden) Code() int {
	return 403
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameForbidden) Error() string {
	return fmt.Sprintf("[GET /v1/apps/name/{name}][%d] edgeApplicationConfigurationGetEdgeApplicationBundleByNameForbidden  %+v", 403, o.Payload)
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameForbidden) String() string {
	return fmt.Sprintf("[GET /v1/apps/name/{name}][%d] edgeApplicationConfigurationGetEdgeApplicationBundleByNameForbidden  %+v", 403, o.Payload)
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameForbidden) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationGetEdgeApplicationBundleByNameNotFound creates a EdgeApplicationConfigurationGetEdgeApplicationBundleByNameNotFound with default headers values
func NewEdgeApplicationConfigurationGetEdgeApplicationBundleByNameNotFound() *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameNotFound {
	return &EdgeApplicationConfigurationGetEdgeApplicationBundleByNameNotFound{}
}

/*
EdgeApplicationConfigurationGetEdgeApplicationBundleByNameNotFound describes a response with status code 404, with default header values.

Not Found. The API gateway did not process the request because the requested resource could not be found.
*/
type EdgeApplicationConfigurationGetEdgeApplicationBundleByNameNotFound struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration get edge application bundle by name not found response has a 2xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameNotFound) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration get edge application bundle by name not found response has a 3xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameNotFound) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration get edge application bundle by name not found response has a 4xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameNotFound) IsClientError() bool {
	return true
}

// IsServerError returns true when this edge application configuration get edge application bundle by name not found response has a 5xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameNotFound) IsServerError() bool {
	return false
}

// IsCode returns true when this edge application configuration get edge application bundle by name not found response a status code equal to that given
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameNotFound) IsCode(code int) bool {
	return code == 404
}

// Code gets the status code for the edge application configuration get edge application bundle by name not found response
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameNotFound) Code() int {
	return 404
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameNotFound) Error() string {
	return fmt.Sprintf("[GET /v1/apps/name/{name}][%d] edgeApplicationConfigurationGetEdgeApplicationBundleByNameNotFound  %+v", 404, o.Payload)
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameNotFound) String() string {
	return fmt.Sprintf("[GET /v1/apps/name/{name}][%d] edgeApplicationConfigurationGetEdgeApplicationBundleByNameNotFound  %+v", 404, o.Payload)
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameNotFound) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationGetEdgeApplicationBundleByNameInternalServerError creates a EdgeApplicationConfigurationGetEdgeApplicationBundleByNameInternalServerError with default headers values
func NewEdgeApplicationConfigurationGetEdgeApplicationBundleByNameInternalServerError() *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameInternalServerError {
	return &EdgeApplicationConfigurationGetEdgeApplicationBundleByNameInternalServerError{}
}

/*
EdgeApplicationConfigurationGetEdgeApplicationBundleByNameInternalServerError describes a response with status code 500, with default header values.

Internal Server Error. The API gateway experienced an unexpected condition. Specific error condition is indicated in error codes.
*/
type EdgeApplicationConfigurationGetEdgeApplicationBundleByNameInternalServerError struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration get edge application bundle by name internal server error response has a 2xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameInternalServerError) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration get edge application bundle by name internal server error response has a 3xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameInternalServerError) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration get edge application bundle by name internal server error response has a 4xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameInternalServerError) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application configuration get edge application bundle by name internal server error response has a 5xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameInternalServerError) IsServerError() bool {
	return true
}

// IsCode returns true when this edge application configuration get edge application bundle by name internal server error response a status code equal to that given
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameInternalServerError) IsCode(code int) bool {
	return code == 500
}

// Code gets the status code for the edge application configuration get edge application bundle by name internal server error response
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameInternalServerError) Code() int {
	return 500
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameInternalServerError) Error() string {
	return fmt.Sprintf("[GET /v1/apps/name/{name}][%d] edgeApplicationConfigurationGetEdgeApplicationBundleByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameInternalServerError) String() string {
	return fmt.Sprintf("[GET /v1/apps/name/{name}][%d] edgeApplicationConfigurationGetEdgeApplicationBundleByNameInternalServerError  %+v", 500, o.Payload)
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameInternalServerError) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameInternalServerError) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationGetEdgeApplicationBundleByNameGatewayTimeout creates a EdgeApplicationConfigurationGetEdgeApplicationBundleByNameGatewayTimeout with default headers values
func NewEdgeApplicationConfigurationGetEdgeApplicationBundleByNameGatewayTimeout() *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameGatewayTimeout {
	return &EdgeApplicationConfigurationGetEdgeApplicationBundleByNameGatewayTimeout{}
}

/*
EdgeApplicationConfigurationGetEdgeApplicationBundleByNameGatewayTimeout describes a response with status code 504, with default header values.

Gateway Timeout. The API gateway did not receive a timely response from an upstream microservice it needed to communicate with in order to complete the request.
*/
type EdgeApplicationConfigurationGetEdgeApplicationBundleByNameGatewayTimeout struct {
	Payload *models.ZsrvResponse
}

// IsSuccess returns true when this edge application configuration get edge application bundle by name gateway timeout response has a 2xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameGatewayTimeout) IsSuccess() bool {
	return false
}

// IsRedirect returns true when this edge application configuration get edge application bundle by name gateway timeout response has a 3xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameGatewayTimeout) IsRedirect() bool {
	return false
}

// IsClientError returns true when this edge application configuration get edge application bundle by name gateway timeout response has a 4xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameGatewayTimeout) IsClientError() bool {
	return false
}

// IsServerError returns true when this edge application configuration get edge application bundle by name gateway timeout response has a 5xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameGatewayTimeout) IsServerError() bool {
	return true
}

// IsCode returns true when this edge application configuration get edge application bundle by name gateway timeout response a status code equal to that given
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameGatewayTimeout) IsCode(code int) bool {
	return code == 504
}

// Code gets the status code for the edge application configuration get edge application bundle by name gateway timeout response
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameGatewayTimeout) Code() int {
	return 504
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameGatewayTimeout) Error() string {
	return fmt.Sprintf("[GET /v1/apps/name/{name}][%d] edgeApplicationConfigurationGetEdgeApplicationBundleByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameGatewayTimeout) String() string {
	return fmt.Sprintf("[GET /v1/apps/name/{name}][%d] edgeApplicationConfigurationGetEdgeApplicationBundleByNameGatewayTimeout  %+v", 504, o.Payload)
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameGatewayTimeout) GetPayload() *models.ZsrvResponse {
	return o.Payload
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameGatewayTimeout) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ZsrvResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEdgeApplicationConfigurationGetEdgeApplicationBundleByNameDefault creates a EdgeApplicationConfigurationGetEdgeApplicationBundleByNameDefault with default headers values
func NewEdgeApplicationConfigurationGetEdgeApplicationBundleByNameDefault(code int) *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameDefault {
	return &EdgeApplicationConfigurationGetEdgeApplicationBundleByNameDefault{
		_statusCode: code,
	}
}

/*
EdgeApplicationConfigurationGetEdgeApplicationBundleByNameDefault describes a response with status code -1, with default header values.

An unexpected error response.
*/
type EdgeApplicationConfigurationGetEdgeApplicationBundleByNameDefault struct {
	_statusCode int

	Payload *models.GooglerpcStatus
}

// IsSuccess returns true when this edge application configuration get edge application bundle by name default response has a 2xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameDefault) IsSuccess() bool {
	return o._statusCode/100 == 2
}

// IsRedirect returns true when this edge application configuration get edge application bundle by name default response has a 3xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameDefault) IsRedirect() bool {
	return o._statusCode/100 == 3
}

// IsClientError returns true when this edge application configuration get edge application bundle by name default response has a 4xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameDefault) IsClientError() bool {
	return o._statusCode/100 == 4
}

// IsServerError returns true when this edge application configuration get edge application bundle by name default response has a 5xx status code
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameDefault) IsServerError() bool {
	return o._statusCode/100 == 5
}

// IsCode returns true when this edge application configuration get edge application bundle by name default response a status code equal to that given
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameDefault) IsCode(code int) bool {
	return o._statusCode == code
}

// Code gets the status code for the edge application configuration get edge application bundle by name default response
func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameDefault) Code() int {
	return o._statusCode
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameDefault) Error() string {
	return fmt.Sprintf("[GET /v1/apps/name/{name}][%d] EdgeApplicationConfiguration_GetEdgeApplicationBundleByName default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameDefault) String() string {
	return fmt.Sprintf("[GET /v1/apps/name/{name}][%d] EdgeApplicationConfiguration_GetEdgeApplicationBundleByName default  %+v", o._statusCode, o.Payload)
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameDefault) GetPayload() *models.GooglerpcStatus {
	return o.Payload
}

func (o *EdgeApplicationConfigurationGetEdgeApplicationBundleByNameDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GooglerpcStatus)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
