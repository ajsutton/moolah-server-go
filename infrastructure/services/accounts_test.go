package services

import (
	"github.com/moolah-server-go/values"
	"github.com/stretchr/testify/require"
	"testing"

	_ "github.com/proullon/ramsql/driver"
)

func TestAccountsList_Empty(t *testing.T) {
	runAccountsTests(t, func(t *testing.T, accounts Accounts) {
		want := []values.Account{}
		got, err := accounts.List("user")
		require.NoError(t, err)
		require.Equal(t, want, got)
	})
}

func TestAccountsCreate(t *testing.T) {
	runAccountsTests(t, func(t *testing.T, accounts Accounts) {
		account := values.NullAccount()
		// Don't yet support transactions which hold balance and date
		account.Balance = 0
		account.Date = values.Date{}
		require.NoError(t, accounts.Create("user", account))

		want := []values.Account{account}
		got, err := accounts.List("user")
		require.NoError(t, err)
		require.Equal(t, want, got)
	})
}

func runAccountsTests(t *testing.T, f func(t *testing.T, accounts Accounts)) {
	accountsDao, err := NewAccounts(values.Config{
		DriverName:     "ramsql",
		DataSourceName: t.Name(),
	})
	require.NoError(t, err, "Failed to create SQL accounts")

	var tests = []struct {
		name     string
		accounts Accounts
	}{
		{"AccountsDao", accountsDao},
		{"Null", NullAccounts()},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			require.NoError(t, test.accounts.Init())
			f(t, test.accounts)
		})
	}
}
