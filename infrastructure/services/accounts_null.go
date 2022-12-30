package services

import "github.com/moolah-server-go/values"

func NullAccounts() Accounts {
	return &AccountsNull{accounts: []values.Account{}}
}

type AccountsNull struct {
	accounts []values.Account
}

func (a *AccountsNull) Init() error {
	return nil
}

func (a *AccountsNull) Create(userId string, account values.Account) error {
	a.accounts = append(a.accounts, account)
	return nil
}

func (a *AccountsNull) List(userId string) ([]values.Account, error) {
	return a.accounts, nil
}
