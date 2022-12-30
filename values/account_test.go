package values

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAccount_SerializeToJson(t *testing.T) {
	account := NullAccount()
	result, err := json.Marshal(account)
	require.NoError(t, err)
	require.JSONEq(t, "{"+
		"\"id\":\"abc123\","+
		"\"name\":\"Test Account 1\","+
		"\"type\": \"Bank\","+
		"\"balance\": 999,"+
		"\"position\": 3,"+
		"\"date\": \"2012-07-15\""+
		"}", string(result))
}
