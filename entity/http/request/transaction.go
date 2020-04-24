package request

type TopUpRequest struct {
	Amount int `json:"amount"`
}

type TransferRequest struct {
	BeneficiaryId int `json:"beneficiary_id"`
	Amount        int `json:"amount"`
}
