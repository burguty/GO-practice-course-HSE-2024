package models

type Account struct {
	Name   string `json: "name"`
	Amount int    `json: "amount"`
}

// Queries
type CreateAccountRequest struct {
	Name string `json:"name"`
}

type DeleteAccountRequest struct {
	Name string `json:"name"`
}

type UpdateAmountRequest struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}

type UpdateNameRequest struct {
	Name    string `json:"name"`
	NewName string `json:"new_name"`
}

type GetAccountRequest struct {
	Name string `json:"name"`
}

// Responses
type GetAccountResponse struct {
	Name   string `json:"name"`
	Amount int    `json:"amount"`
}
