package services

import (
	"github.com/moolah-server-go/values"
	"github.com/stretchr/testify/assert"
	"testing"
)

var tests = []struct {
	name     string
	accounts Accounts
}{
	{"AccountsDao", NewAccounts()},
	{"Null", NullAccounts()},
}

func TestAccountsList(t *testing.T) {
	runTests(t, func(t *testing.T, accounts Accounts) {
		var want []values.Account
		got := accounts.List()
		assert.Equal(t, want, got)
	})
}

func runTests(t *testing.T, f func(t *testing.T, accounts Accounts)) {
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			f(t, test.accounts)
		})
	}
}
