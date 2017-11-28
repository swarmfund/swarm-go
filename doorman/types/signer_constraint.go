package types

import "net/http"

type SignerConstraint func(r *http.Request) error
