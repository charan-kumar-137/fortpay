package dto

type AccountRequest struct {
	Id     string `json:"id"`
	Amount int64  `json:"amount"`
}
