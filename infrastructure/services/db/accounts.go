package db

import (
	"database/sql"
	"github.com/moolah-server-go/values"
)

type Accounts interface {
	Init() error
	List(userId string) ([]values.Account, error)
	Create(userId string, account values.Account) error
}

func NewAccounts(config values.Config) (Accounts, error) {
	db, err := sql.Open(config.DriverName, config.DataSourceName)
	if err != nil {
		return nil, err
	}
	return &accountsDao{db: db}, nil
}

type accountsDao struct {
	db *sql.DB
}

func (dao *accountsDao) Init() error {
	_, err := dao.db.Exec(
		"CREATE TABLE IF NOT EXISTS account (" +
			"    user_id VARCHAR(255) NOT NULL," +
			"    id VARCHAR(100) NOT NULL," +
			"    name VARCHAR(255) NOT NULL," +
			"    type VARCHAR(25) NOT NULL," +
			"    balance INT NOT NULL," +
			"    position int(16) NOT NULL default 0," +
			"    PRIMARY KEY (user_id, id)" +
			")")
	return err
}

func (dao *accountsDao) Create(userId string, account values.Account) error {
	_, err := dao.db.Exec("INSERT INTO account (user_id, id, name, type, position) VALUES (?, ?, ?, ?, ?)",
		userId, account.Id, account.Name, account.Type, account.Position)
	return err
}

func (dao *accountsDao) List(userId string) ([]values.Account, error) {
	rows, err := dao.db.Query(
		"SELECT id, name, type, position"+
			" FROM account"+
			" WHERE user_id = ?"+
			" ORDER BY position, name", userId)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	accounts := []values.Account{}
	for rows.Next() {
		var account values.Account
		if err := rows.Scan(&account.Id, &account.Name, &account.Type, &account.Position); err != nil {
			return accounts, err
		}
		accounts = append(accounts, account)
	}
	return accounts, nil
}
