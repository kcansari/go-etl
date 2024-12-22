package models

import (
	"database/sql"
	"database/sql/driver"
	"encoding/json"
)

// NullYear represents a nullable year
type NullYear sql.NullInt16

// Tsvector represents a PostgreSQL tsvector
type Tsvector string

// Trigger represents a PostgreSQL trigger
type Trigger string

// Value implements the driver.Valuer interface
func (ny NullYear) Value() (driver.Value, error) {
	if !ny.Valid {
		return nil, nil
	}
	return ny.Int16, nil
}

// Scan implements the sql.Scanner interface
func (ny *NullYear) Scan(value interface{}) error {
	var n sql.NullInt16
	if err := n.Scan(value); err != nil {
		return err
	}
	*ny = NullYear(n)
	return nil
}

// MarshalJSON implements the json.Marshaler interface
func (ny NullYear) MarshalJSON() ([]byte, error) {
	if !ny.Valid {
		return []byte("null"), nil
	}
	return json.Marshal(ny.Int16)
}

// UnmarshalJSON implements the json.Unmarshaler interface
func (ny *NullYear) UnmarshalJSON(data []byte) error {
	var v *int16
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	if v == nil {
		ny.Valid = false
		return nil
	}
	ny.Int16 = *v
	ny.Valid = true
	return nil
}
