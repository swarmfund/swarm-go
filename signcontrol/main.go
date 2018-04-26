package signcontrol

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"gitlab.com/tokend/go/hash"
	"gitlab.com/tokend/go/keypair"
	"gitlab.com/tokend/go/xdr"
)

const (
	SignatureHeader  = "x-authsignature"
	PublicKeyHeader  = "x-authpublickey"
	ValidUntilHeader = "x-authvaliduntilltimestamp"
	ValidUntilOffset = 60
)

type Signature struct {
	Address       string
	Signer        string
	RawValidUntil string
	Signature     string
}

func IsSigned(request *http.Request) (*Signature, bool) {
	if request == nil {
		panic(ErrNilRequest)
	}

	s := Signature{
		Signer:        request.Header.Get(PublicKeyHeader),
		Signature:     request.Header.Get(SignatureHeader),
		RawValidUntil: request.Header.Get(ValidUntilHeader),
	}
	return &s, s.Signature != "" && s.Signer != "" && s.RawValidUntil != ""
}

func SignRequest(request *http.Request, kp keypair.KP) error {
	// TODO check if request is nil
	// TODO check if kp is nil

	validUntil := fmt.Sprintf("%d", time.Now().Unix()+ValidUntilOffset)

	request.Header.Set(ValidUntilHeader, validUntil)
	request.Header.Set(PublicKeyHeader, kp.Address())

	base := fmt.Sprintf("{ uri: '%s', valid_untill: '%s'}", request.URL.RequestURI(), validUntil)
	hashBase := hash.Hash([]byte(base))

	decorated, err := kp.SignDecorated(hashBase[:])
	if err != nil {
		return err
	}

	encodedSign, err := xdr.MarshalBase64(decorated)
	if err != nil {
		return err
	}

	request.Header.Set(SignatureHeader, encodedSign)
	return nil
}

func CheckSignature(request *http.Request) (string, error) {
	// TODO cache results to request.Context()
	sig, ok := IsSigned(request)
	if !ok {
		return "", ErrNotSigned
	}

	validUntil, err := strconv.ParseInt(sig.RawValidUntil, 10, 64)
	if err != nil {
		return "", ErrValidUntil
	}

	if time.Unix(validUntil, 0).Before(time.Now()) {
		return "", ErrExpired
	}

	signatureBase := "{ uri: '" + request.URL.RequestURI() + "', valid_untill: '" + sig.RawValidUntil + "'}"
	hashBase := hash.Hash([]byte(signatureBase))
	pubKey, err := keypair.Parse(sig.Signer)
	if err != nil {
		return "", ErrSignerKey
	}

	var decoratedSign xdr.DecoratedSignature
	err = xdr.SafeUnmarshalBase64(sig.Signature, &decoratedSign)
	if err != nil {
		return "", ErrSignature
	}

	if pubKey.Verify(hashBase[:], decoratedSign.Signature) != nil {
		return "", ErrSignature
	}

	return sig.Signer, nil
}
