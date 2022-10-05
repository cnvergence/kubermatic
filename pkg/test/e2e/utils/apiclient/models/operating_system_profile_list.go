// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"

	"github.com/go-openapi/strfmt"
)

// OperatingSystemProfileList OperatingSystemProfileList defines a map of operating system and the OperatingSystemProfile to use.
//
// swagger:model OperatingSystemProfileList
type OperatingSystemProfileList map[string]string

// Validate validates this operating system profile list
func (m OperatingSystemProfileList) Validate(formats strfmt.Registry) error {
	return nil
}

// ContextValidate validates this operating system profile list based on context it is used
func (m OperatingSystemProfileList) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}