// Code generated by go-swagger; DO NOT EDIT.

package patch_envelope_configuration

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// NewPatchEnvelopeConfigurationGetPatchEnvelopeParams creates a new PatchEnvelopeConfigurationGetPatchEnvelopeParams object,
// with the default timeout for this client.
//
// Default values are not hydrated, since defaults are normally applied by the API server side.
//
// To enforce default values in parameter, use SetDefaults or WithDefaults.
func NewPatchEnvelopeConfigurationGetPatchEnvelopeParams() *PatchEnvelopeConfigurationGetPatchEnvelopeParams {
	return &PatchEnvelopeConfigurationGetPatchEnvelopeParams{
		timeout: cr.DefaultTimeout,
	}
}

// NewPatchEnvelopeConfigurationGetPatchEnvelopeParamsWithTimeout creates a new PatchEnvelopeConfigurationGetPatchEnvelopeParams object
// with the ability to set a timeout on a request.
func NewPatchEnvelopeConfigurationGetPatchEnvelopeParamsWithTimeout(timeout time.Duration) *PatchEnvelopeConfigurationGetPatchEnvelopeParams {
	return &PatchEnvelopeConfigurationGetPatchEnvelopeParams{
		timeout: timeout,
	}
}

// NewPatchEnvelopeConfigurationGetPatchEnvelopeParamsWithContext creates a new PatchEnvelopeConfigurationGetPatchEnvelopeParams object
// with the ability to set a context for a request.
func NewPatchEnvelopeConfigurationGetPatchEnvelopeParamsWithContext(ctx context.Context) *PatchEnvelopeConfigurationGetPatchEnvelopeParams {
	return &PatchEnvelopeConfigurationGetPatchEnvelopeParams{
		Context: ctx,
	}
}

// NewPatchEnvelopeConfigurationGetPatchEnvelopeParamsWithHTTPClient creates a new PatchEnvelopeConfigurationGetPatchEnvelopeParams object
// with the ability to set a custom HTTPClient for a request.
func NewPatchEnvelopeConfigurationGetPatchEnvelopeParamsWithHTTPClient(client *http.Client) *PatchEnvelopeConfigurationGetPatchEnvelopeParams {
	return &PatchEnvelopeConfigurationGetPatchEnvelopeParams{
		HTTPClient: client,
	}
}

/*
PatchEnvelopeConfigurationGetPatchEnvelopeParams contains all the parameters to send to the API endpoint

	for the patch envelope configuration get patch envelope operation.

	Typically these are written to a http.Request.
*/
type PatchEnvelopeConfigurationGetPatchEnvelopeParams struct {

	/* XRequestID.

	   User-Agent specified id to track a request
	*/
	XRequestID *string

	/* NamePattern.

	   patch envelope name pattern
	*/
	NamePattern *string

	/* NextOrderBy.

	   OrderBy helps in sorting the list response
	*/
	NextOrderBy []string

	/* NextPageNum.

	   Page Number

	   Format: int64
	*/
	NextPageNum *int64

	/* NextPageSize.

	   Defines the page size

	   Format: int64
	*/
	NextPageSize *int64

	/* NextPageToken.

	   Page Token
	*/
	NextPageToken *string

	/* NextTotalPages.

	   Total number of pages to be fetched.

	   Format: int64
	*/
	NextTotalPages *int64

	/* ProjectNamePattern.

	   project name pattern
	*/
	ProjectNamePattern *string

	/* Summary.

	   Only summary of the records required
	*/
	Summary *bool

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithDefaults hydrates default values in the patch envelope configuration get patch envelope params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeParams) WithDefaults() *PatchEnvelopeConfigurationGetPatchEnvelopeParams {
	o.SetDefaults()
	return o
}

// SetDefaults hydrates default values in the patch envelope configuration get patch envelope params (not the query body).
//
// All values with no default are reset to their zero value.
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeParams) SetDefaults() {
	// no default values defined for this parameter
}

// WithTimeout adds the timeout to the patch envelope configuration get patch envelope params
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeParams) WithTimeout(timeout time.Duration) *PatchEnvelopeConfigurationGetPatchEnvelopeParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the patch envelope configuration get patch envelope params
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the patch envelope configuration get patch envelope params
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeParams) WithContext(ctx context.Context) *PatchEnvelopeConfigurationGetPatchEnvelopeParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the patch envelope configuration get patch envelope params
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the patch envelope configuration get patch envelope params
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeParams) WithHTTPClient(client *http.Client) *PatchEnvelopeConfigurationGetPatchEnvelopeParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the patch envelope configuration get patch envelope params
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithXRequestID adds the xRequestID to the patch envelope configuration get patch envelope params
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeParams) WithXRequestID(xRequestID *string) *PatchEnvelopeConfigurationGetPatchEnvelopeParams {
	o.SetXRequestID(xRequestID)
	return o
}

// SetXRequestID adds the xRequestId to the patch envelope configuration get patch envelope params
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeParams) SetXRequestID(xRequestID *string) {
	o.XRequestID = xRequestID
}

// WithNamePattern adds the namePattern to the patch envelope configuration get patch envelope params
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeParams) WithNamePattern(namePattern *string) *PatchEnvelopeConfigurationGetPatchEnvelopeParams {
	o.SetNamePattern(namePattern)
	return o
}

// SetNamePattern adds the namePattern to the patch envelope configuration get patch envelope params
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeParams) SetNamePattern(namePattern *string) {
	o.NamePattern = namePattern
}

// WithNextOrderBy adds the nextOrderBy to the patch envelope configuration get patch envelope params
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeParams) WithNextOrderBy(nextOrderBy []string) *PatchEnvelopeConfigurationGetPatchEnvelopeParams {
	o.SetNextOrderBy(nextOrderBy)
	return o
}

// SetNextOrderBy adds the nextOrderBy to the patch envelope configuration get patch envelope params
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeParams) SetNextOrderBy(nextOrderBy []string) {
	o.NextOrderBy = nextOrderBy
}

// WithNextPageNum adds the nextPageNum to the patch envelope configuration get patch envelope params
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeParams) WithNextPageNum(nextPageNum *int64) *PatchEnvelopeConfigurationGetPatchEnvelopeParams {
	o.SetNextPageNum(nextPageNum)
	return o
}

// SetNextPageNum adds the nextPageNum to the patch envelope configuration get patch envelope params
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeParams) SetNextPageNum(nextPageNum *int64) {
	o.NextPageNum = nextPageNum
}

// WithNextPageSize adds the nextPageSize to the patch envelope configuration get patch envelope params
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeParams) WithNextPageSize(nextPageSize *int64) *PatchEnvelopeConfigurationGetPatchEnvelopeParams {
	o.SetNextPageSize(nextPageSize)
	return o
}

// SetNextPageSize adds the nextPageSize to the patch envelope configuration get patch envelope params
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeParams) SetNextPageSize(nextPageSize *int64) {
	o.NextPageSize = nextPageSize
}

// WithNextPageToken adds the nextPageToken to the patch envelope configuration get patch envelope params
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeParams) WithNextPageToken(nextPageToken *string) *PatchEnvelopeConfigurationGetPatchEnvelopeParams {
	o.SetNextPageToken(nextPageToken)
	return o
}

// SetNextPageToken adds the nextPageToken to the patch envelope configuration get patch envelope params
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeParams) SetNextPageToken(nextPageToken *string) {
	o.NextPageToken = nextPageToken
}

// WithNextTotalPages adds the nextTotalPages to the patch envelope configuration get patch envelope params
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeParams) WithNextTotalPages(nextTotalPages *int64) *PatchEnvelopeConfigurationGetPatchEnvelopeParams {
	o.SetNextTotalPages(nextTotalPages)
	return o
}

// SetNextTotalPages adds the nextTotalPages to the patch envelope configuration get patch envelope params
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeParams) SetNextTotalPages(nextTotalPages *int64) {
	o.NextTotalPages = nextTotalPages
}

// WithProjectNamePattern adds the projectNamePattern to the patch envelope configuration get patch envelope params
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeParams) WithProjectNamePattern(projectNamePattern *string) *PatchEnvelopeConfigurationGetPatchEnvelopeParams {
	o.SetProjectNamePattern(projectNamePattern)
	return o
}

// SetProjectNamePattern adds the projectNamePattern to the patch envelope configuration get patch envelope params
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeParams) SetProjectNamePattern(projectNamePattern *string) {
	o.ProjectNamePattern = projectNamePattern
}

// WithSummary adds the summary to the patch envelope configuration get patch envelope params
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeParams) WithSummary(summary *bool) *PatchEnvelopeConfigurationGetPatchEnvelopeParams {
	o.SetSummary(summary)
	return o
}

// SetSummary adds the summary to the patch envelope configuration get patch envelope params
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeParams) SetSummary(summary *bool) {
	o.Summary = summary
}

// WriteToRequest writes these params to a swagger request
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.XRequestID != nil {

		// header param X-Request-Id
		if err := r.SetHeaderParam("X-Request-Id", *o.XRequestID); err != nil {
			return err
		}
	}

	if o.NamePattern != nil {

		// query param namePattern
		var qrNamePattern string

		if o.NamePattern != nil {
			qrNamePattern = *o.NamePattern
		}
		qNamePattern := qrNamePattern
		if qNamePattern != "" {

			if err := r.SetQueryParam("namePattern", qNamePattern); err != nil {
				return err
			}
		}
	}

	if o.NextOrderBy != nil {

		// binding items for next.orderBy
		joinedNextOrderBy := o.bindParamNextOrderBy(reg)

		// query array param next.orderBy
		if err := r.SetQueryParam("next.orderBy", joinedNextOrderBy...); err != nil {
			return err
		}
	}

	if o.NextPageNum != nil {

		// query param next.pageNum
		var qrNextPageNum int64

		if o.NextPageNum != nil {
			qrNextPageNum = *o.NextPageNum
		}
		qNextPageNum := swag.FormatInt64(qrNextPageNum)
		if qNextPageNum != "" {

			if err := r.SetQueryParam("next.pageNum", qNextPageNum); err != nil {
				return err
			}
		}
	}

	if o.NextPageSize != nil {

		// query param next.pageSize
		var qrNextPageSize int64

		if o.NextPageSize != nil {
			qrNextPageSize = *o.NextPageSize
		}
		qNextPageSize := swag.FormatInt64(qrNextPageSize)
		if qNextPageSize != "" {

			if err := r.SetQueryParam("next.pageSize", qNextPageSize); err != nil {
				return err
			}
		}
	}

	if o.NextPageToken != nil {

		// query param next.pageToken
		var qrNextPageToken string

		if o.NextPageToken != nil {
			qrNextPageToken = *o.NextPageToken
		}
		qNextPageToken := qrNextPageToken
		if qNextPageToken != "" {

			if err := r.SetQueryParam("next.pageToken", qNextPageToken); err != nil {
				return err
			}
		}
	}

	if o.NextTotalPages != nil {

		// query param next.totalPages
		var qrNextTotalPages int64

		if o.NextTotalPages != nil {
			qrNextTotalPages = *o.NextTotalPages
		}
		qNextTotalPages := swag.FormatInt64(qrNextTotalPages)
		if qNextTotalPages != "" {

			if err := r.SetQueryParam("next.totalPages", qNextTotalPages); err != nil {
				return err
			}
		}
	}

	if o.ProjectNamePattern != nil {

		// query param projectNamePattern
		var qrProjectNamePattern string

		if o.ProjectNamePattern != nil {
			qrProjectNamePattern = *o.ProjectNamePattern
		}
		qProjectNamePattern := qrProjectNamePattern
		if qProjectNamePattern != "" {

			if err := r.SetQueryParam("projectNamePattern", qProjectNamePattern); err != nil {
				return err
			}
		}
	}

	if o.Summary != nil {

		// query param summary
		var qrSummary bool

		if o.Summary != nil {
			qrSummary = *o.Summary
		}
		qSummary := swag.FormatBool(qrSummary)
		if qSummary != "" {

			if err := r.SetQueryParam("summary", qSummary); err != nil {
				return err
			}
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

// bindParamPatchEnvelopeConfigurationGetPatchEnvelope binds the parameter next.orderBy
func (o *PatchEnvelopeConfigurationGetPatchEnvelopeParams) bindParamNextOrderBy(formats strfmt.Registry) []string {
	nextOrderByIR := o.NextOrderBy

	var nextOrderByIC []string
	for _, nextOrderByIIR := range nextOrderByIR { // explode []string

		nextOrderByIIV := nextOrderByIIR // string as string
		nextOrderByIC = append(nextOrderByIC, nextOrderByIIV)
	}

	// items.CollectionFormat: "multi"
	nextOrderByIS := swag.JoinByFormat(nextOrderByIC, "multi")

	return nextOrderByIS
}
