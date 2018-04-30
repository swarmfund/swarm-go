package signcontrol

import (
	"net/http"
	"testing"

	"fmt"

	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
	"gitlab.com/tokend/go/keypair"
)

const (
	seed = "SCDMOOXVNMO6SA22AYUMZDIGLDJMBUTVEGB73FFNTLFJILBJWIU4NQ3D"
)

func TestSignRequest(t *testing.T) {
	kp, _ := keypair.Parse(seed)
	request, _ := http.NewRequest("GET", "http://example.com/users", nil)
	request.Header.Set("date", "Sun, 05 Jan 2014 21:31:40 GMT")

	err := SignRequest(request, kp)
	assert.NoError(t, err)
	assert.Equal(t,
		fmt.Sprintf(`keyId="%s",algorithm="%s",signature="gg0tjYeo/wicZVKpR/NRCUoIgJSLZdb4s5RXAQ1jNF1VYs4W3wDGg1JWQM1U3tCQscIzfYlTNRsgW1JA2ax5Ag==",headers="date (request-target)"`, kp.Address(), SignatureAlgorithm),
		request.Header.Get("signature"),
	)
}

func TestCheckSignatureValid(t *testing.T) {
	kp, _ := keypair.Parse(seed)
	request, _ := http.NewRequest("GET", "http://host/path?q=a", nil)
	SignRequest(request, kp)

	signer, err := CheckSignature(request)
	require.Nil(t, err)
	require.Equal(t, signer, kp.Address())
}
