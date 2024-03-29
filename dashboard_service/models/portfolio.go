// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/swag"
)

// Portfolio portfolio
// swagger:model Portfolio
type Portfolio struct {

	// loss
	Loss float64 `json:"loss,omitempty"`

	// name
	Name string `json:"name,omitempty"`

	// price per share bought
	PricePerShareBought float64 `json:"pricePerShareBought,omitempty"`

	// price per share now
	PricePerShareNow float64 `json:"pricePerShareNow,omitempty"`

	// profit
	Profit float64 `json:"profit,omitempty"`

	// shares held
	SharesHeld float64 `json:"sharesHeld,omitempty"`

	// symbol
	Symbol string `json:"symbol,omitempty"`

	// total value bought
	TotalValueBought float64 `json:"totalValueBought,omitempty"`

	// total value now
	TotalValueNow float64 `json:"totalValueNow,omitempty"`
}

// Validate validates this portfolio
func (m *Portfolio) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *Portfolio) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *Portfolio) UnmarshalBinary(b []byte) error {
	var res Portfolio
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
