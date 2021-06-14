package logout

import (
	ozzo "github.com/go-ozzo/ozzo-validation/v4"
)

type Request struct {
	SessionID string `json:"session_id"`
}

func (r *Request) Validate() error {
	return ozzo.ValidateStruct(
		r,
		ozzo.Field(&r.SessionID, ozzo.Required),
	)
}
