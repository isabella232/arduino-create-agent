// Code generated by goa v2.0.0-wip, DO NOT EDIT.
//
// tools HTTP server types
//
// Command:
// $ goa gen github.com/arduino/arduino-create-agent/design

package server

import (
	tools "github.com/arduino/arduino-create-agent/gen/tools"
	toolsviews "github.com/arduino/arduino-create-agent/gen/tools/views"
	goa "goa.design/goa"
)

// InstallRequestBody is the type of the "tools" service "install" endpoint
// HTTP request body.
type InstallRequestBody struct {
	// The name of the tool
	Name *string `form:"name,omitempty" json:"name,omitempty" xml:"name,omitempty"`
	// The version of the tool
	Version *string `form:"version,omitempty" json:"version,omitempty" xml:"version,omitempty"`
	// The packager of the tool
	Packager *string `form:"packager,omitempty" json:"packager,omitempty" xml:"packager,omitempty"`
}

// ToolResponseCollection is the type of the "tools" service "available"
// endpoint HTTP response body.
type ToolResponseCollection []*ToolResponse

// ToolResponse is used to define fields on response body types.
type ToolResponse struct {
	// The name of the tool
	Name string `form:"name" json:"name" xml:"name"`
	// The version of the tool
	Version string `form:"version" json:"version" xml:"version"`
	// The packager of the tool
	Packager string `form:"packager" json:"packager" xml:"packager"`
}

// NewToolResponseCollection builds the HTTP response body from the result of
// the "available" endpoint of the "tools" service.
func NewToolResponseCollection(res toolsviews.ToolCollectionView) ToolResponseCollection {
	body := make([]*ToolResponse, len(res))
	for i, val := range res {
		body[i] = &ToolResponse{
			Name:     *val.Name,
			Version:  *val.Version,
			Packager: *val.Packager,
		}
	}
	return body
}

// NewInstallToolPayload builds a tools service install endpoint payload.
func NewInstallToolPayload(body *InstallRequestBody) *tools.ToolPayload {
	v := &tools.ToolPayload{
		Name:     *body.Name,
		Version:  *body.Version,
		Packager: *body.Packager,
	}
	return v
}

// ValidateInstallRequestBody runs the validations defined on InstallRequestBody
func ValidateInstallRequestBody(body *InstallRequestBody) (err error) {
	if body.Name == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("name", "body"))
	}
	if body.Version == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("version", "body"))
	}
	if body.Packager == nil {
		err = goa.MergeErrors(err, goa.MissingFieldError("packager", "body"))
	}
	return
}
