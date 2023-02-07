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

// ApplicationSource application source
//
// swagger:model ApplicationSource
type ApplicationSource struct {

	// git
	Git *GitSource `json:"git,omitempty"`

	// helm
	Helm *HelmSource `json:"helm,omitempty"`
}

// Validate validates this application source
func (m *ApplicationSource) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateHelm(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationSource) validateGit(formats strfmt.Registry) error {
	if swag.IsZero(m.Git) { // not required
		return nil
	}

	if m.Git != nil {
		if err := m.Git.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("git")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("git")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationSource) validateHelm(formats strfmt.Registry) error {
	if swag.IsZero(m.Helm) { // not required
		return nil
	}

	if m.Helm != nil {
		if err := m.Helm.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("helm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("helm")
			}
			return err
		}
	}

	return nil
}

// ContextValidate validate this application source based on the context it is used
func (m *ApplicationSource) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateGit(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateHelm(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *ApplicationSource) contextValidateGit(ctx context.Context, formats strfmt.Registry) error {

	if m.Git != nil {
		if err := m.Git.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("git")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("git")
			}
			return err
		}
	}

	return nil
}

func (m *ApplicationSource) contextValidateHelm(ctx context.Context, formats strfmt.Registry) error {

	if m.Helm != nil {
		if err := m.Helm.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("helm")
			} else if ce, ok := err.(*errors.CompositeError); ok {
				return ce.ValidateName("helm")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *ApplicationSource) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *ApplicationSource) UnmarshalBinary(b []byte) error {
	var res ApplicationSource
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}