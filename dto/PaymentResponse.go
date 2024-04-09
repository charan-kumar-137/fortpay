package dto

type PaymentResponse struct {
	ReferenceId string `json:"referenceId"`
	Status      string `json:"status"`
	Description string `json:"description"`
}

func CreatePaymentResponse(referenceId string, status string, description string) PaymentResponse {
	var paymentResponse PaymentResponse
	paymentResponse.ReferenceId = referenceId
	paymentResponse.Status = status
	paymentResponse.Description = description
	return paymentResponse
}
