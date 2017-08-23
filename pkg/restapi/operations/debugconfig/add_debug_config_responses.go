package debugconfig

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"

	"github.com/solo-io/squash/pkg/models"
)

// AddDebugConfigCreatedCode is the HTTP code returned for type AddDebugConfigCreated
const AddDebugConfigCreatedCode int = 201

/*AddDebugConfigCreated Debug config created

swagger:response addDebugConfigCreated
*/
type AddDebugConfigCreated struct {

	/*
	  In: Body
	*/
	Payload *models.DebugConfig `json:"body,omitempty"`
}

// NewAddDebugConfigCreated creates AddDebugConfigCreated with default headers values
func NewAddDebugConfigCreated() *AddDebugConfigCreated {
	return &AddDebugConfigCreated{}
}

// WithPayload adds the payload to the add debug config created response
func (o *AddDebugConfigCreated) WithPayload(payload *models.DebugConfig) *AddDebugConfigCreated {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the add debug config created response
func (o *AddDebugConfigCreated) SetPayload(payload *models.DebugConfig) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *AddDebugConfigCreated) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(201)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}

// AddDebugConfigBadRequestCode is the HTTP code returned for type AddDebugConfigBadRequest
const AddDebugConfigBadRequestCode int = 400

/*AddDebugConfigBadRequest Bad request

swagger:response addDebugConfigBadRequest
*/
type AddDebugConfigBadRequest struct {
}

// NewAddDebugConfigBadRequest creates AddDebugConfigBadRequest with default headers values
func NewAddDebugConfigBadRequest() *AddDebugConfigBadRequest {
	return &AddDebugConfigBadRequest{}
}

// WriteResponse to the client
func (o *AddDebugConfigBadRequest) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(400)
}

// AddDebugConfigNotFoundCode is the HTTP code returned for type AddDebugConfigNotFound
const AddDebugConfigNotFoundCode int = 404

/*AddDebugConfigNotFound Attachment not found

swagger:response addDebugConfigNotFound
*/
type AddDebugConfigNotFound struct {
}

// NewAddDebugConfigNotFound creates AddDebugConfigNotFound with default headers values
func NewAddDebugConfigNotFound() *AddDebugConfigNotFound {
	return &AddDebugConfigNotFound{}
}

// WriteResponse to the client
func (o *AddDebugConfigNotFound) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(404)
}

// AddDebugConfigUnprocessableEntityCode is the HTTP code returned for type AddDebugConfigUnprocessableEntity
const AddDebugConfigUnprocessableEntityCode int = 422

/*AddDebugConfigUnprocessableEntity Invalid input

swagger:response addDebugConfigUnprocessableEntity
*/
type AddDebugConfigUnprocessableEntity struct {
}

// NewAddDebugConfigUnprocessableEntity creates AddDebugConfigUnprocessableEntity with default headers values
func NewAddDebugConfigUnprocessableEntity() *AddDebugConfigUnprocessableEntity {
	return &AddDebugConfigUnprocessableEntity{}
}

// WriteResponse to the client
func (o *AddDebugConfigUnprocessableEntity) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(422)
}

// AddDebugConfigServiceUnavailableCode is the HTTP code returned for type AddDebugConfigServiceUnavailable
const AddDebugConfigServiceUnavailableCode int = 503

/*AddDebugConfigServiceUnavailable Service Unavailable

swagger:response addDebugConfigServiceUnavailable
*/
type AddDebugConfigServiceUnavailable struct {
}

// NewAddDebugConfigServiceUnavailable creates AddDebugConfigServiceUnavailable with default headers values
func NewAddDebugConfigServiceUnavailable() *AddDebugConfigServiceUnavailable {
	return &AddDebugConfigServiceUnavailable{}
}

// WriteResponse to the client
func (o *AddDebugConfigServiceUnavailable) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(503)
}