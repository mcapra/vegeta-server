// Code generated by go-swagger; DO NOT EDIT.

package report

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	models "vegeta-server/models"
)

// GetReportByIDOKCode is the HTTP code returned for type GetReportByIDOK
const GetReportByIDOKCode int = 200

/*GetReportByIDOK Success

swagger:response getReportByIdOK
*/
type GetReportByIDOK struct {

	/*
	  In: Body
	*/
	Payload *models.ReportResponse `json:"body,omitempty"`
}

// NewGetReportByIDOK creates GetReportByIDOK with default headers values
func NewGetReportByIDOK() *GetReportByIDOK {

	return &GetReportByIDOK{}
}

// WithPayload adds the payload to the get report by Id o k response
func (o *GetReportByIDOK) WithPayload(payload *models.ReportResponse) *GetReportByIDOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get report by Id o k response
func (o *GetReportByIDOK) SetPayload(payload *models.ReportResponse) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReportByIDOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetReportByIDNotFoundCode is the HTTP code returned for type GetReportByIDNotFound
const GetReportByIDNotFoundCode int = 404

/*GetReportByIDNotFound Not Found

swagger:response getReportByIdNotFound
*/
type GetReportByIDNotFound struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetReportByIDNotFound creates GetReportByIDNotFound with default headers values
func NewGetReportByIDNotFound() *GetReportByIDNotFound {

	return &GetReportByIDNotFound{}
}

// WithPayload adds the payload to the get report by Id not found response
func (o *GetReportByIDNotFound) WithPayload(payload *models.Error) *GetReportByIDNotFound {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get report by Id not found response
func (o *GetReportByIDNotFound) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReportByIDNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// GetReportByIDInternalServerErrorCode is the HTTP code returned for type GetReportByIDInternalServerError
const GetReportByIDInternalServerErrorCode int = 500

/*GetReportByIDInternalServerError Internal Server Error

swagger:response getReportByIdInternalServerError
*/
type GetReportByIDInternalServerError struct {

	/*
	  In: Body
	*/
	Payload *models.Error `json:"body,omitempty"`
}

// NewGetReportByIDInternalServerError creates GetReportByIDInternalServerError with default headers values
func NewGetReportByIDInternalServerError() *GetReportByIDInternalServerError {

	return &GetReportByIDInternalServerError{}
}

// WithPayload adds the payload to the get report by Id internal server error response
func (o *GetReportByIDInternalServerError) WithPayload(payload *models.Error) *GetReportByIDInternalServerError {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the get report by Id internal server error response
func (o *GetReportByIDInternalServerError) SetPayload(payload *models.Error) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *GetReportByIDInternalServerError) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(500)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}
