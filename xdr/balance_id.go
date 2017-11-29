package xdr

import (
	"fmt"

	"gitlab.com/swarmfund/go/strkey"
)

func (bid *BalanceId) AsString() string {
	if bid == nil {
		return ""
	}

	switch bid.Type {
	case CryptoKeyTypeKeyTypeEd25519:
		ed := bid.MustEd25519()
		raw := make([]byte, 32)
		copy(raw, ed[:])
		return strkey.MustEncode(strkey.VersionByteBalanceID, raw)
	default:
		panic(fmt.Errorf("Unknown account id type: %v", bid.Type))
	}
}

func (bid *BalanceId) MarshalJSON() ([]byte, error) {
	return []byte("\"" + bid.AsString() + "\""), nil
}

// Equals returns true if `other` is equivalent to `aid`
func (bid *BalanceId) Equals(other BalanceId) bool {
	if bid.Type != other.Type {
		return false
	}

	switch bid.Type {
	case CryptoKeyTypeKeyTypeEd25519:
		l := bid.MustEd25519()
		r := other.MustEd25519()
		return l == r
	default:
		panic(fmt.Errorf("Unknown account id type: %v", bid.Type))
	}
}
