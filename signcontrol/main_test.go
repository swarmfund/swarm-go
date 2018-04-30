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
	request, _ := http.NewRequest("GET", "http://example.com/fo?ob=ar", nil)
	request.Header.Set("date", "Sun, 05 Jan 2014 21:31:40 GMT")
	err := SignRequest(request, kp)
	assert.NoError(t, err)
	assert.Equal(t,
		fmt.Sprintf(`keyId="%s",algorithm="%s",signature="b3nmxe6bTCIlKh4+YAP11i2i5CRkACsNlXS+LKdzGTHKG3puuCHRGlvUmfF/RMgQUPep6YC/5Bu/ZmD9egy5Dg==",headers="date (request-target)"`, kp.Address(), SignatureAlgorithm),
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
