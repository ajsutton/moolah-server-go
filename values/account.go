package values

type Account struct {
	Id       string `json:"id"`
	Name     string `json:"name" binding:"required"`
	Type     string `json:"type" binding:"required"`
	Balance  uint64 `json:"balance" binding:"required"`
	Position uint64 `json:"position"`
	Date     Date   `json:"date"`
}
