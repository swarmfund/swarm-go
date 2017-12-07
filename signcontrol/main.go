package signcontrol

import (
	"net/http"
	"strconv"
	"time"

	"fmt"

	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	"gitlab.com/swarmfund/go/hash"
	"gitlab.com/swarmfund/go/keypair"
	"gitlab.com/swarmfund/go/xdr"
)

type ctxKey int

const (
	SkipValidation ctxKey = iota
)

const (
	SignatureHeader     = "x-authsignature"
	HMACSignatureHeader = "x-hmac-signature"
	PublicKeyHeader     = "x-authpublickey"
	ValidUntilHeader    = "x-authvaliduntilltimestamp"
	ValidUntilOffset    = 60
)

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
	// check if signature is valid
	signer := request.Header.Get(PublicKeyHeader)

	if request.Context().Value(SkipValidation) != nil {
		return signer, nil
	}

	signature := request.Header.Get(SignatureHeader)
	rawValidUntil := request.Header.Get(ValidUntilHeader)

	if signature == "" || signer == "" || rawValidUntil == "" {
		return "", ErrNotSigned
	}

	validUntil, err := strconv.ParseInt(rawValidUntil, 10, 64)
	if err != nil {
		return "", ErrValidUntil
	}

	if time.Unix(validUntil, 0).Before(time.Now()) {
		return "", ErrExpired
	}

	signatureBase := "{ uri: '" + request.URL.RequestURI() + "', valid_untill: '" + rawValidUntil + "'}"
	hashBase := hash.Hash([]byte(signatureBase))
	pubKey, err := keypair.Parse(signer)
	if err != nil {
		return "", ErrSignerKey
	}

	var decoratedSign xdr.DecoratedSignature
	err = xdr.SafeUnmarshalBase64(signature, &decoratedSign)
	if err != nil {
		return "", ErrSignature
	}

	if pubKey.Verify(hashBase[:], decoratedSign.Signature) != nil {
		return "", ErrSignature
	}

	return signer, nil
}

func CheckHMACSignature(request *http.Request, key string) error {
	signature := request.Header.Get(HMACSignatureHeader)
	rawValidUntil := request.Header.Get(ValidUntilHeader)
	signer := request.Header.Get(PublicKeyHeader)

	if signature == "" || signer == "" || rawValidUntil == "" {
		return ErrNotSigned
	}

	base := fmt.Sprintf("%s:%s:%s", request.Method, request.URL.RequestURI(), rawValidUntil)
	h := hmac.New(sha256.New, []byte(key))
	h.Write([]byte(base))
	sign := h.Sum(nil)
	encodedSign := hex.EncodeToString(sign)
	if encodedSign != signature {
		return ErrNotAllowed
	}
	return nil
}
