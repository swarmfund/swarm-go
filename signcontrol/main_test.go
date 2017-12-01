package signcontrol

import (
	"testing"

	"net/http"

	"crypto/hmac"

	"encoding/hex"

	"crypto/sha256"

	"gitlab.com/swarmfund/go/keypair"
	"github.com/stretchr/testify/require"
)

const (
	seed = "SCIMXOIWIB32R7JI6ISNYSCO2BFST5E6P3TLBM4TTLHH57IK6SJPGZT2"
)

func TestSignRequest(t *testing.T) {
	kp, err := keypair.Parse(seed)
	require.Nil(t, err)

	request, err := http.NewRequest("GET", "http://host/path?q=a", nil)
	require.Nil(t, err)

	err = SignRequest(request, kp)
	require.Nil(t, err)

	require.Equal(t, request.Host, "host")
	require.Equal(t, request.URL.Path, "/path")
	require.Equal(t, request.URL.RawQuery, "q=a")
	// TODO mock time.Now() to assert actual values
	require.NotEmpty(t, request.Header.Get(SignatureHeader))
	require.NotEmpty(t, request.Header.Get(ValidUntilHeader))
	require.Equal(t, request.Header.Get(PublicKeyHeader), kp.Address())
}

func TestCheckSignatureValid(t *testing.T) {
	kp, _ := keypair.Parse(seed)
	request, _ := http.NewRequest("GET", "http://host/path?q=a", nil)
	SignRequest(request, kp)

	signer, err := CheckSignature(request)
	require.Nil(t, err)
	require.Equal(t, signer, kp.Address())
}

func TestPHPHMAC(t *testing.T) {
	// > hash_hmac("sha256", "foo&bar", "secret");
	// > 92b4d2c0dad8be0b4bc747f75d8bb4810543b7bec314a9f8abd1b64058f77d8a
	h := hmac.New(sha256.New, []byte("secret"))
	h.Write([]byte("foo&bar"))
	sign := h.Sum(nil)
	if hex.EncodeToString(sign) != "92b4d2c0dad8be0b4bc747f75d8bb4810543b7bec314a9f8abd1b64058f77d8a" {
		t.Fatal("does not match")
	}
}
