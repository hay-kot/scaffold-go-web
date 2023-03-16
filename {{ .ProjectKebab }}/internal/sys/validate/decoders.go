package validate

import (
	"encoding/json"
	"net/http"
)

// DecodeT decodes a request body into a struct of type T. This is a wrapper around
// DecodeStruct using generics.
func DecodeT[T any](r *http.Request) (T, error) {
	var v T
	return v, DecodeStruct(r, &v)
}

// DecodeStruct decodes a request body into a struct and validates the struct using
// the Check function.
func DecodeStruct(r *http.Request, v interface{}) error {
	if err := json.NewDecoder(r.Body).Decode(v); err != nil {
		return err
	}

	err := Check(v)
	if err != nil {
		return err
	}

	return nil
}
