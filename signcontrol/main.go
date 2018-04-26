package signcontrol

import (
	"errors"
	"net/http"
	"strconv"
	"time"

	"gitlab.com/tokend/go/hash"
	"gitlab.com/tokend/go/keypair"
	"gitlab.com/tokend/go/signcontrol/internal/httpsignatures"
	"gitlab.com/tokend/go/xdr"
)

const (
	SignatureHeader  = "x-authsignature"
	PublicKeyHeader  = "x-authpublickey"
	ValidUntilHeader = "x-authvaliduntilltimestamp"
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
	signer := httpsignatures.NewSigner(httpsignatures.KeypairAlgorithm{}, "date", httpsignatures.RequestTarget)
	err := signer.SignRequest(kp.Address(), kp, request)
	if err != nil {
		return err
	}
	return nil
}

func CheckSignature(request *http.Request) (string, error) {
	// check if it v2 signature
	if request.Header.Get("signature") != "" || request.Header.Get("authorization") != "" {
		return checkV2(request)
	}

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

func checkV2(request *http.Request) (string, error) {
	signature, err := httpsignatures.FromRequest(request)
	if err != nil {
		return "", err
	}
	if ok := signature.IsValid(request); !ok {
		return "", errors.New("invalid signature")
	}
	return signature.KeyID, nil
}
