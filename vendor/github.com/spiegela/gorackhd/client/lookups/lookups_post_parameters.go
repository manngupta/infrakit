package lookups

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/spiegela/gorackhd/models"
)

// NewLookupsPostParams creates a new LookupsPostParams object
// with the default values initialized.
func NewLookupsPostParams() *LookupsPostParams {
	var ()
	return &LookupsPostParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewLookupsPostParamsWithTimeout creates a new LookupsPostParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewLookupsPostParamsWithTimeout(timeout time.Duration) *LookupsPostParams {
	var ()
	return &LookupsPostParams{

		timeout: timeout,
	}
}

// NewLookupsPostParamsWithContext creates a new LookupsPostParams object
// with the default values initialized, and the ability to set a context for a request
func NewLookupsPostParamsWithContext(ctx context.Context) *LookupsPostParams {
	var ()
	return &LookupsPostParams{

		Context: ctx,
	}
}

// NewLookupsPostParamsWithHTTPClient creates a new LookupsPostParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewLookupsPostParamsWithHTTPClient(client *http.Client) *LookupsPostParams {
	var ()
	return &LookupsPostParams{
		HTTPClient: client,
	}
}

/*LookupsPostParams contains all the parameters to send to the API endpoint
for the lookups post operation typically these are written to a http.Request
*/
type LookupsPostParams struct {

	/*Body
	  The properties of the lookup

	*/
	Body *models.Lookups20LookupBase

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the lookups post params
func (o *LookupsPostParams) WithTimeout(timeout time.Duration) *LookupsPostParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the lookups post params
func (o *LookupsPostParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the lookups post params
func (o *LookupsPostParams) WithContext(ctx context.Context) *LookupsPostParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the lookups post params
func (o *LookupsPostParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the lookups post params
func (o *LookupsPostParams) WithHTTPClient(client *http.Client) *LookupsPostParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the lookups post params
func (o *LookupsPostParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithBody adds the body to the lookups post params
func (o *LookupsPostParams) WithBody(body *models.Lookups20LookupBase) *LookupsPostParams {
	o.SetBody(body)
	return o
}

// SetBody adds the body to the lookups post params
func (o *LookupsPostParams) SetBody(body *models.Lookups20LookupBase) {
	o.Body = body
}

// WriteToRequest writes these params to a swagger request
func (o *LookupsPostParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Body == nil {
		o.Body = new(models.Lookups20LookupBase)
	}

	if err := r.SetBodyParam(o.Body); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
