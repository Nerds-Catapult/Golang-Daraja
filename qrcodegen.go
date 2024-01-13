package darajaAuth

type TransactionType string

const (
	TransactionTypeBuyGoods       TransactionType = "BG"
	TransactionTypePayBill        TransactionType = "PB"
	TransactionTypeWithdraw       TransactionType = "WA" // Withdraw Cash at Agent Till.
	TransactionTypeSendMoney      TransactionType = "SM" // Send Money to a Phone Number.
	TransactionTypeSendtoBusiness TransactionType = "SB" // Send Money to a Business.
)

type QrPayload struct {
	MerchantName          string          `json:"MerchantName"`
	RefNo                 string          `json:"RefNo"`
	Amount                string          `json:"Amount"`
	TransactionType       TransactionType `json:"TransactionType"`
	CreditPartyIdentifier string          `json:"CreditPartyIdentifier"`
}

type QrResponse struct {
	ResponseCode        string `json:"ResponseCode"`
	RequestID           string `json:"RequestID"`
	ResponseDescription string `json:"ResponseDescription"`
	QrCode              string `json:"QrCode"`
}

func (d *Daraja) MakeQRCodeRequest(payload QrPayload) (*QrResponse, *ErrorResponse) {
	secureResponse, err := performSecurePostRequest[*QrResponse](payload, endpointQrCode, d)
	if err != nil {
		return nil, err
	}
	return secureResponse.Body, nil
}
