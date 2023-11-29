package darajaAuth

//business to customer

type B2CCommandID string

const (
	B2CCommandIDSalary B2CCommandID = "SalaryPayment"
	B2CCommandIDBusinessPayment B2CCommandID = "BusinessPayment"
	B2CCommandIDPromotionPayment B2CCommandID = "PromotionPayment"
)

type B2CPaymentPayload struct {
	InitiatorName string `json:"InitiatorName"`
	PassKey string `json:"PassKey"`
	CommandID B2CCommandID `json:"CommandID"`
	Amount string `json:"Amount"`
	PartyA string `json:"PartyA"`
	PartyB string `json:"PartyB"`
	Remarks string `json:"Remarks"`
	QueueTimeOutURL string `json:"QueueTimeOutURL"`
	ResultURL string `json:"ResultURL"`
	Occasion string `json:"Occasion"`
}

type B2CPaymentResponse struct {
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ResponseCode string `json:"ResponseCode"`
	ResponseDescription string `json:"ResponseDescription"`
	ConversationID string `json:"ConversationID"`
}

func (d *Daraja)MakeB2CPaymentRequest(b2c B2CPaymentPayload, certPath string) (*B2CPaymentResponse, *ErrorResponse) {
	b2c.CommandID = "BusinessPayment"
	encryptedCredentials, err := openSSlEncrypt(b2c.PassKey, certPath)
	if err != nil {
		return nil, &ErrorResponse{error: err}
	}
	b2c.PassKey = encryptedCredentials
	secureResponse, errRes := performSecurePostRequest[*B2CPaymentResponse](b2c, endpointB2CPmtReq, d)
	if errRes != nil {
		return nil, errRes
	}
	return secureResponse.Body, nil
}