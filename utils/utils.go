package utils

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

func GetCurrentDate() string {
	today := time.Now().Format("2006-01-02T15:04:05.000")
	fmt.Println(today)
	return today
}

func GetFutureDate() string {
	future := time.Now().AddDate(0,0,14).Format("2006-01-02T15:04:05.000")
	return future
}

// My brain has a hard time getting these two right
// Decode → "Deconstruct JSON" → "JSON to Go".
//Encode → "Construct JSON" → "Go to JSON".

// This generic function takes any type and decodes json to specifed type
func DecodeJson[T any](r *http.Request) (T, error){
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
func EncodeJson[T any](w http.ResponseWriter, staus int, v T) error {
	w.Header().Set("content/type", "application/json")
	w.WriteHeader(staus)
	if err := json.NewEncoder(w).Encode(v); err != nil {
		return fmt.Errorf("encode json %w", err)
	}
	fmt.Fprintf(w,"Recored! %v", v)
	return nil
}