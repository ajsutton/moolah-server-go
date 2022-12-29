package services

import "github.com/moolah-server-go/values"

type Accounts interface {
	List() []values.Account
}

func NewAccounts() Accounts {
	return &AccountsDao{}
}

type AccountsDao struct {
}

func (dao *AccountsDao) List() []values.Account {
	var accounts []values.Account
	return accounts
}
