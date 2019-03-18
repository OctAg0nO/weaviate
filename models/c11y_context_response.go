/*                          _       _
 *__      _____  __ ___   ___  __ _| |_ ___
 *\ \ /\ / / _ \/ _` \ \ / / |/ _` | __/ _ \
 * \ V  V /  __/ (_| |\ V /| | (_| | ||  __/
 *  \_/\_/ \___|\__,_| \_/ |_|\__,_|\__\___|
 *
 * Copyright © 2016 - 2019 Weaviate. All rights reserved.
 * LICENSE: https://github.com/creativesoftwarefdn/weaviate/blob/develop/LICENSE.md
 * DESIGN & CONCEPT: Bob van Luijt (@bobvanluijt)
 * CONTACT: hello@creativesoftwarefdn.org
 */ // Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"strconv"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// C11yContextResponse An array of available words and contexts.
// swagger:model C11yContextResponse
type C11yContextResponse struct {

	// results
	Results []*C11yContextResponseResultsItems0 `json:"results"`
}

// Validate validates this c11y context response
func (m *C11yContextResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateResults(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *C11yContextResponse) validateResults(formats strfmt.Registry) error {

	if swag.IsZero(m.Results) { // not required
		return nil
	}

	for i := 0; i < len(m.Results); i++ {
		if swag.IsZero(m.Results[i]) { // not required
			continue
		}

		if m.Results[i] != nil {
			if err := m.Results[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("results" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *C11yContextResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *C11yContextResponse) UnmarshalBinary(b []byte) error {
	var res C11yContextResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// C11yContextResponseResultsItems0 c11y context response results items0
// swagger:model C11yContextResponseResultsItems0
type C11yContextResponseResultsItems0 struct {

	// in c11y
	InC11y bool `json:"inC11y,omitempty"`

	// info
	Info *C11yContextResponseResultsItems0Info `json:"info,omitempty"`

	// word
	Word string `json:"word,omitempty"`
}

// Validate validates this c11y context response results items0
func (m *C11yContextResponseResultsItems0) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateInfo(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *C11yContextResponseResultsItems0) validateInfo(formats strfmt.Registry) error {

	if swag.IsZero(m.Info) { // not required
		return nil
	}

	if m.Info != nil {
		if err := m.Info.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("info")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *C11yContextResponseResultsItems0) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *C11yContextResponseResultsItems0) UnmarshalBinary(b []byte) error {
	var res C11yContextResponseResultsItems0
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}

// C11yContextResponseResultsItems0Info c11y context response results items0 info
// swagger:model C11yContextResponseResultsItems0Info
type C11yContextResponseResultsItems0Info struct {

	// nearest neighbors
	NearestNeighbors C11yNearestNeighbors `json:"nearestNeighbors,omitempty"`

	// vector
	Vector C11yVector `json:"vector,omitempty"`
}

// Validate validates this c11y context response results items0 info
func (m *C11yContextResponseResultsItems0Info) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNearestNeighbors(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateVector(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *C11yContextResponseResultsItems0Info) validateNearestNeighbors(formats strfmt.Registry) error {

	if swag.IsZero(m.NearestNeighbors) { // not required
		return nil
	}

	if err := m.NearestNeighbors.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("info" + "." + "nearestNeighbors")
		}
		return err
	}

	return nil
}

func (m *C11yContextResponseResultsItems0Info) validateVector(formats strfmt.Registry) error {

	if swag.IsZero(m.Vector) { // not required
		return nil
	}

	if err := m.Vector.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("info" + "." + "vector")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *C11yContextResponseResultsItems0Info) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *C11yContextResponseResultsItems0Info) UnmarshalBinary(b []byte) error {
	var res C11yContextResponseResultsItems0Info
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
