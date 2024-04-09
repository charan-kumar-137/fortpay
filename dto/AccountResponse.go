package dto

type AccountResponse struct {
	AccountId   string `json:"account_id"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

func CreateAccountResponse(accountId string, status string, description string) AccountResponse {

	var accountResponse AccountResponse

	accountResponse.AccountId = accountId
	accountResponse.Status = status
	accountResponse.Description = description

	return accountResponse
}
