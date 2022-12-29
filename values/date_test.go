package values

import (
	"encoding/json"
	"fmt"
	"github.com/stretchr/testify/require"
	"testing"
	"time"
)

func TestDate_MarshalJSON(t *testing.T) {
	d :=
		Date{time.Date(2016, 5, 17, 0, 0, 0, 0, time.UTC)}
	result, err := json.Marshal(d)
	if err != nil {
		t.Error(err)
	}
	fmt.Println(d, result)
	require.JSONEq(t, "\"2016-05-17\"", string(result))
}

func TestDate_UnmarshallJSON(t *testing.T) {
	result := Date{}
	err := json.Unmarshal([]byte("\"2016-05-17\""), &result)
	if err != nil {
		t.Error(err)
	}
	expected :=
		Date{time.Date(2016, 5, 17, 0, 0, 0, 0, time.UTC)}
	if result != expected {
		t.Error("Expected", expected, "but got", result)
	}
}
