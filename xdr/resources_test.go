package xdr_test

import (
	"testing"

	"encoding/json"

	"github.com/stretchr/testify/assert"
	"gitlab.com/swarmfund/go/xdr"
)

func TestFlag(t *testing.T) {
	expected := `{"int":6,"flags":[{"name":"kyc_update","value":2},{"name":"suspicious_behavior","value":4}]}`
	value := xdr.BlockReasonsKycUpdate ^ xdr.BlockReasonsSuspiciousBehavior

	t.Run("marshal", func(t *testing.T) {
		got, err := json.Marshal(&value)
		if err != nil {
			t.Fatal(err)
		}
		assert.JSONEq(t, string(got), expected)
	})

	t.Run("unmarshal", func(t *testing.T) {
		var got xdr.BlockReasons
		assert.NoError(t, json.Unmarshal([]byte(expected), &got))
		assert.Equal(t, got, value)
	})
}

func TestEnum(t *testing.T) {
	expected := `{"int":2,"string":"general"}`
	value := xdr.AccountTypeGeneral

	t.Run("marshal", func(t *testing.T) {
		got, err := json.Marshal(&value)
		if err != nil {
			t.Fatal(err)
		}
		assert.JSONEq(t, string(got), expected)
	})

	t.Run("unmarshal", func(t *testing.T) {
		var got xdr.AccountType
		assert.NoError(t, json.Unmarshal([]byte(expected), &got))
		assert.Equal(t, got, value)
	})
}
