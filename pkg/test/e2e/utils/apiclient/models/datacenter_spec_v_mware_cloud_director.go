// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DatacenterSpecVMwareCloudDirector datacenter spec v mware cloud director
//
// swagger:model DatacenterSpecVMwareCloudDirector
type DatacenterSpecVMwareCloudDirector struct {

	// If set to true, disables the TLS certificate check against the endpoint.
	AllowInsecure bool `json:"allowInsecure,omitempty"`

	// The default catalog which contains the VM templates.
	DefaultCatalog string `json:"catalog,omitempty"`

	// The name of the storage profile to use for disks attached to the VMs.
	DefaultStorageProfile string `json:"storageProfile,omitempty"`

	// Endpoint URL to use, including protocol, for example "https://vclouddirector.example.com".
	URL string `json:"url,omitempty"`

	// templates
	Templates ImageList `json:"templates,omitempty"`
}

// Validate validates this datacenter spec v mware cloud director
func (m *DatacenterSpecVMwareCloudDirector) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateTemplates(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatacenterSpecVMwareCloudDirector) validateTemplates(formats strfmt.Registry) error {
	if swag.IsZero(m.Templates) { // not required
		return nil
	}

	if m.Templates != nil {
		if err := m.Templates.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("templates")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("templates")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this datacenter spec v mware cloud director based on the context it is used
func (m *DatacenterSpecVMwareCloudDirector) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateTemplates(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *DatacenterSpecVMwareCloudDirector) contextValidateTemplates(ctx context.Context, formats strfmt.Registry) error {

	if err := m.Templates.ContextValidate(ctx, formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("templates")
		} else if ce, ok := err.(*errors.CompositeError); ok {
			return ce.ValidateName("templates")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *DatacenterSpecVMwareCloudDirector) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DatacenterSpecVMwareCloudDirector) UnmarshalBinary(b []byte) error {
	var res DatacenterSpecVMwareCloudDirector
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
