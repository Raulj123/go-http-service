package httpjson

import (
	"encoding/json"
	"fmt"
	"net/http"
)

// My brain has a hard time getting these two right
// Decode → "Deconstruct JSON" → "JSON to Go".
//Encode → "Construct JSON" → "Go to JSON".

// This generic function takes any type and decodes json to specifed type
func Decode[T any](r *http.Request) (T, error){
	var v T
	d := json.NewDecoder(r.Body)
	// no unkown fields bitch
	d.DisallowUnknownFields()
	if err := d.Decode((&v)); err != nil {
		return v, fmt.Errorf("decode json %w", err)
	}
	
	return v, nil
}

// This function encodes shit
func Encode[T any](w http.ResponseWriter, staus int, v T) error {
	w.Header().Set("content/type", "application/json")
	w.WriteHeader(staus)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json %w", err)
	}
	return nil
}