package values

import "time"

type Account struct {
	Id       string `json:"id"`
	Name     string `json:"name" binding:"required"`
	Type     string `json:"type" binding:"required"`
	Balance  uint64 `json:"balance" binding:"required"`
	Position uint64 `json:"position"`
	Date     Date   `json:"date"`
}

func NullAccount() Account {
	return Account{
		Id:       "abc123",
		Name:     "Test Account 1",
		Type:     "Bank",
		Balance:  999,
		Position: 3,
		Date:     MakeDate(2012, time.July, 15),
	}
}
