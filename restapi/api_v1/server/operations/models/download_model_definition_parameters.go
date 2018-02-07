// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	"github.com/go-openapi/runtime/middleware"
	"github.com/go-openapi/validate"

	strfmt "github.com/go-openapi/strfmt"
)

// NewDownloadModelDefinitionParams creates a new DownloadModelDefinitionParams object
// with the default values initialized.
func NewDownloadModelDefinitionParams() DownloadModelDefinitionParams {
	var (
		versionDefault = string("2017-02-13")
	)
	return DownloadModelDefinitionParams{
		Version: versionDefault,
	}
}

// DownloadModelDefinitionParams contains all the bound params for the download model definition operation
// typically these are obtained from a http.Request
//
// swagger:parameters downloadModelDefinition
type DownloadModelDefinitionParams struct {

	// HTTP Request Object
	HTTPRequest *http.Request

	/*The id of the model.
	  Required: true
	  In: path
	*/
	ModelID string
	/*The release date of the version of the API you want to use. Specify dates in YYYY-MM-DD format.
	  Required: true
	  In: query
	  Default: "2017-02-13"
	*/
	Version string
}

// BindRequest both binds and validates a request, it assumes that complex things implement a Validatable(strfmt.Registry) error interface
// for simple values it will use straight method calls
func (o *DownloadModelDefinitionParams) BindRequest(r *http.Request, route *middleware.MatchedRoute) error {
	var res []error
	o.HTTPRequest = r

	qs := runtime.Values(r.URL.Query())

	rModelID, rhkModelID, _ := route.Params.GetOK("model_id")
	if err := o.bindModelID(rModelID, rhkModelID, route.Formats); err != nil {
		res = append(res, err)
	}

	qVersion, qhkVersion, _ := qs.GetOK("version")
	if err := o.bindVersion(qVersion, qhkVersion, route.Formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (o *DownloadModelDefinitionParams) bindModelID(rawData []string, hasKey bool, formats strfmt.Registry) error {
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}

	o.ModelID = raw

	return nil
}

func (o *DownloadModelDefinitionParams) bindVersion(rawData []string, hasKey bool, formats strfmt.Registry) error {
	if !hasKey {
		return errors.Required("version", "query")
	}
	var raw string
	if len(rawData) > 0 {
		raw = rawData[len(rawData)-1]
	}
	if err := validate.RequiredString("version", "query", raw); err != nil {
		return err
	}

	o.Version = raw

	return nil
}