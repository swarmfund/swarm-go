package signcontrol

import "errors"

var (
	ErrNotSigned  = errors.New("request is not signed")
	ErrValidUntil = errors.New("valid until is not valid")
	ErrExpired    = errors.New("expired signature")
	ErrSignerKey  = errors.New("signer key is not valid")
	ErrSignature  = errors.New("signature is not valid")
	ErrNotAllowed = errors.New("not allowed")
)
