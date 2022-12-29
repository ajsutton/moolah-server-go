package services

import "github.com/moolah-server-go/values"

type Accounts interface {
	List() []values.Account
}

func NewAccounts() Accounts {
	return &accountsDao{}
}

type accountsDao struct {
}

func (dao *accountsDao) List() []values.Account {
	accounts := []values.Account{}
	return accounts
}
