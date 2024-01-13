package darajaAuth

//b2b means business to business

type B2BPaymentPayload struct {
	InitiatorName   string `json:"Initiator"`
	Passkey         string `json:"SecurityCredential"`
	CommandID       string `json:"CommandID"`
	Amount          string `json:"Amount"`
	PartyA          string `json:"PartyA"`
	PartyB          string `json:"PartyB"`
	Remarks         string `json:"Remarks"`
	QueueTimeOutURL string `json:"QueueTimeOutURL"`
	ResultURL       string `json:"ResultURL"`
	Occasion        string `json:"Occasion"`
}

type B2BPaymentResponse struct {
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ConversationID           string `json:"ConversationID"`
	ResponseDescription      string `json:"ResponseDescription"`
}

func (d *Daraja) MakeB2BPaymentRequest(b2c B2BPaymentPayload, certPath string) (*B2CPaymentResponse, *ErrorResponse) {
	b2c.CommandID = "BusinessPayment"
	encryptedCredentials, err := openSSlEncrypt(b2c.Passkey, certPath)
	if err != nil {
		return nil, &ErrorResponse{error: err}
	}
	b2c.Passkey = encryptedCredentials
	secureResponse, errRes := performSecurePostRequest[*B2CPaymentResponse](b2c, endpointB2CPmtReq, d)
	if err != nil {
		return nil, errRes
	}
	return secureResponse.Body, nil
}
