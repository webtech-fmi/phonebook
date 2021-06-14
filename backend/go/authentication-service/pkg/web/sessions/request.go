package sessions

import (
	ozzo "github.com/go-ozzo/ozzo-validation/v4"
)

type SessionValidRequest struct {
	ID    string
}

func (r *SessionValidRequest) Validate() error {
	return ozzo.ValidateStruct(
		r,
		ozzo.Field(&r.ID, ozzo.Required, ozzo.Length(64, 64)),
	)
}
