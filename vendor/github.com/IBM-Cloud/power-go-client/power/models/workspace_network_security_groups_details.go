// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WorkspaceNetworkSecurityGroupsDetails workspace network security groups details
//
// swagger:model WorkspaceNetworkSecurityGroupsDetails
type WorkspaceNetworkSecurityGroupsDetails struct {

	// The state of a Network Security Groups configuration
	// Required: true
	// Enum: ["active","error","configuring","removing","inactive","inaccessible"]
	State *string `json:"state"`
}

// Validate validates this workspace network security groups details
func (m *WorkspaceNetworkSecurityGroupsDetails) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateState(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var workspaceNetworkSecurityGroupsDetailsTypeStatePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","error","configuring","removing","inactive","inaccessible"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		workspaceNetworkSecurityGroupsDetailsTypeStatePropEnum = append(workspaceNetworkSecurityGroupsDetailsTypeStatePropEnum, v)
	}
}

const (

	// WorkspaceNetworkSecurityGroupsDetailsStateActive captures enum value "active"
	WorkspaceNetworkSecurityGroupsDetailsStateActive string = "active"

	// WorkspaceNetworkSecurityGroupsDetailsStateError captures enum value "error"
	WorkspaceNetworkSecurityGroupsDetailsStateError string = "error"

	// WorkspaceNetworkSecurityGroupsDetailsStateConfiguring captures enum value "configuring"
	WorkspaceNetworkSecurityGroupsDetailsStateConfiguring string = "configuring"

	// WorkspaceNetworkSecurityGroupsDetailsStateRemoving captures enum value "removing"
	WorkspaceNetworkSecurityGroupsDetailsStateRemoving string = "removing"

	// WorkspaceNetworkSecurityGroupsDetailsStateInactive captures enum value "inactive"
	WorkspaceNetworkSecurityGroupsDetailsStateInactive string = "inactive"

	// WorkspaceNetworkSecurityGroupsDetailsStateInaccessible captures enum value "inaccessible"
	WorkspaceNetworkSecurityGroupsDetailsStateInaccessible string = "inaccessible"
)

// prop value enum
func (m *WorkspaceNetworkSecurityGroupsDetails) validateStateEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, workspaceNetworkSecurityGroupsDetailsTypeStatePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WorkspaceNetworkSecurityGroupsDetails) validateState(formats strfmt.Registry) error {

	if err := validate.Required("state", "body", m.State); err != nil {
		return err
	}

	// value enum
	if err := m.validateStateEnum("state", "body", *m.State); err != nil {
		return err
	}

	return nil
}

// ContextValidate validates this workspace network security groups details based on context it is used
func (m *WorkspaceNetworkSecurityGroupsDetails) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *WorkspaceNetworkSecurityGroupsDetails) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WorkspaceNetworkSecurityGroupsDetails) UnmarshalBinary(b []byte) error {
	var res WorkspaceNetworkSecurityGroupsDetails
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
