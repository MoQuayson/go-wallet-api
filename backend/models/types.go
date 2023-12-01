package models

import (
	"database/sql"
	"encoding/json"
	"time"
)

type NullString struct {
	sql.NullString
}

var nullString *NullString

// func NullString() *NullString {
// 	return nullString
// }

func ConvertToNullString(v interface{}) NullString {
	if v == nil {
		v = ""
	}
	return NullString{
		sql.NullString{
			String: v.(string),
			Valid:  true,
		},
	}
}

func (NullString) FromString(v string) NullString {
	return NullString{
		sql.NullString{
			String: v,
			Valid:  true,
		},
	}
}

func (v NullString) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.String)
	} else {
		return json.Marshal(nil)
	}
}

func (v *NullString) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *string
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.String = *x
	} else {
		v.Valid = false
	}
	return nil
}

type NullDateTime struct {
	sql.NullTime
}

func (v NullDateTime) MarshalJSON() ([]byte, error) {
	if v.Valid {
		return json.Marshal(v.Time)
	} else {
		return json.Marshal(nil)
	}
}

func (v *NullDateTime) UnmarshalJSON(data []byte) error {
	// Unmarshalling into a pointer will let us detect null
	var x *time.Time
	if err := json.Unmarshal(data, &x); err != nil {
		return err
	}
	if x != nil {
		v.Valid = true
		v.Time = *x
	} else {
		v.Valid = false
	}
	return nil
}
