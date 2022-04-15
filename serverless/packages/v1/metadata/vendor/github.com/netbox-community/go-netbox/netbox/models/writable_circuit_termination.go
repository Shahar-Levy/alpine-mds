// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

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

// WritableCircuitTermination writable circuit termination
//
// swagger:model WritableCircuitTermination
type WritableCircuitTermination struct {

	// occupied
	// Read Only: true
	Occupied *bool `json:"_occupied,omitempty"`

	// cable
	Cable *NestedCable `json:"cable,omitempty"`

	// Cable peer
	//
	//
	// Return the appropriate serializer for the cable termination model.
	//
	// Read Only: true
	CablePeer interface{} `json:"cable_peer,omitempty"`

	// Cable peer type
	// Read Only: true
	CablePeerType string `json:"cable_peer_type,omitempty"`

	// Circuit
	// Required: true
	Circuit *int64 `json:"circuit"`

	// Description
	// Max Length: 200
	Description string `json:"description,omitempty"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// Id
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Mark connected
	//
	// Treat as if a cable is connected
	MarkConnected bool `json:"mark_connected,omitempty"`

	// Port speed (Kbps)
	// Maximum: 2.147483647e+09
	// Minimum: 0
	PortSpeed *int64 `json:"port_speed,omitempty"`

	// Patch panel/port(s)
	// Max Length: 100
	PpInfo string `json:"pp_info,omitempty"`

	// Provider network
	ProviderNetwork *int64 `json:"provider_network,omitempty"`

	// Site
	Site *int64 `json:"site,omitempty"`

	// Termination
	// Required: true
	// Enum: [A Z]
	TermSide *string `json:"term_side"`

	// Upstream speed (Kbps)
	//
	// Upstream speed, if different from port speed
	// Maximum: 2.147483647e+09
	// Minimum: 0
	UpstreamSpeed *int64 `json:"upstream_speed,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`

	// Cross-connect ID
	// Max Length: 50
	XconnectID string `json:"xconnect_id,omitempty"`
}

// Validate validates this writable circuit termination
func (m *WritableCircuitTermination) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCable(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCircuit(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateDescription(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePortSpeed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePpInfo(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTermSide(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateUpstreamSpeed(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateXconnectID(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableCircuitTermination) validateCable(formats strfmt.Registry) error {
	if swag.IsZero(m.Cable) { // not required
		return nil
	}

	if m.Cable != nil {
		if err := m.Cable.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *WritableCircuitTermination) validateCircuit(formats strfmt.Registry) error {

	if err := validate.Required("circuit", "body", m.Circuit); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuitTermination) validateDescription(formats strfmt.Registry) error {
	if swag.IsZero(m.Description) { // not required
		return nil
	}

	if err := validate.MaxLength("description", "body", m.Description, 200); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuitTermination) validatePortSpeed(formats strfmt.Registry) error {
	if swag.IsZero(m.PortSpeed) { // not required
		return nil
	}

	if err := validate.MinimumInt("port_speed", "body", *m.PortSpeed, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("port_speed", "body", *m.PortSpeed, 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuitTermination) validatePpInfo(formats strfmt.Registry) error {
	if swag.IsZero(m.PpInfo) { // not required
		return nil
	}

	if err := validate.MaxLength("pp_info", "body", m.PpInfo, 100); err != nil {
		return err
	}

	return nil
}

var writableCircuitTerminationTypeTermSidePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["A","Z"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableCircuitTerminationTypeTermSidePropEnum = append(writableCircuitTerminationTypeTermSidePropEnum, v)
	}
}

const (

	// WritableCircuitTerminationTermSideA captures enum value "A"
	WritableCircuitTerminationTermSideA string = "A"

	// WritableCircuitTerminationTermSideZ captures enum value "Z"
	WritableCircuitTerminationTermSideZ string = "Z"
)

// prop value enum
func (m *WritableCircuitTermination) validateTermSideEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableCircuitTerminationTypeTermSidePropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableCircuitTermination) validateTermSide(formats strfmt.Registry) error {

	if err := validate.Required("term_side", "body", m.TermSide); err != nil {
		return err
	}

	// value enum
	if err := m.validateTermSideEnum("term_side", "body", *m.TermSide); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuitTermination) validateUpstreamSpeed(formats strfmt.Registry) error {
	if swag.IsZero(m.UpstreamSpeed) { // not required
		return nil
	}

	if err := validate.MinimumInt("upstream_speed", "body", *m.UpstreamSpeed, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("upstream_speed", "body", *m.UpstreamSpeed, 2.147483647e+09, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuitTermination) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuitTermination) validateXconnectID(formats strfmt.Registry) error {
	if swag.IsZero(m.XconnectID) { // not required
		return nil
	}

	if err := validate.MaxLength("xconnect_id", "body", m.XconnectID, 50); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable circuit termination based on the context it is used
func (m *WritableCircuitTermination) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateOccupied(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCable(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateCablePeerType(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableCircuitTermination) contextValidateOccupied(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "_occupied", "body", m.Occupied); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuitTermination) contextValidateCable(ctx context.Context, formats strfmt.Registry) error {

	if m.Cable != nil {
		if err := m.Cable.ContextValidate(ctx, formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("cable")
			}
			return err
		}
	}

	return nil
}

func (m *WritableCircuitTermination) contextValidateCablePeerType(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "cable_peer_type", "body", string(m.CablePeerType)); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuitTermination) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuitTermination) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WritableCircuitTermination) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableCircuitTermination) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableCircuitTermination) UnmarshalBinary(b []byte) error {
	var res WritableCircuitTermination
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
