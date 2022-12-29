package values

import (
	"encoding/json"
	"github.com/stretchr/testify/require"
	"testing"
)

func TestAccount_SerializeToJson(t *testing.T) {
	date, err := ParseDate("2012-07-15")
	require.NoError(t, err)
	account := Account{
		Id:       "a1234",
		Name:     "Account Name",
		Type:     "Bank",
		Balance:  uint64(9999),
		Position: uint64(3),
		Date:     date,
	}
	result, err := json.Marshal(account)
	require.NoError(t, err)
	require.JSONEq(t, "{"+
		"\"id\":\"a1234\","+
		"\"name\":\"Account Name\","+
		"\"type\": \"Bank\","+
		"\"balance\": 9999,"+
		"\"position\": 3,"+
		"\"date\": \"2012-07-15\""+
		"}", string(result))
}
