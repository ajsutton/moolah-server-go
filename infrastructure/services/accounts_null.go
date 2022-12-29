package services

import "github.com/moolah-server-go/values"

func NullAccounts() Accounts {
	return &AccountsNull{}
}

type AccountsNull struct {
}

func (a *AccountsNull) List() []values.Account {
	return []values.Account{}
}
