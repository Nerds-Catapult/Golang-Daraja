package darajaAuth

type C2BPayload struct {
	ShortCode     string `json:"ShortCode"`
	CommandID     string `json:"CommandID"`
	Amount        string `json:"Amount"`
	Msisdn        string `json:"Msisdn"`
	BillRefNumber string `json:"BillRefNumber"`
}

type C2BResponse struct {
	OriginitatorConversationID string `json:"OriginitatorConversationID"`
	ResponseDescription        string `json:"ResponseDescription"`
	ConversationID             string `json:"ConversationID"`
}

type C2BRegistrationPayload struct {
	ValidationURL   string `json:"ValidationURL"`
	ConfirmationURL string `json:"ConfirmationURL"`
	ResponseType    string `json:"ResponseType"`
	ShortCode       string `json:"ShortCode"`
}

type C2BRegistrationResponse struct {
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ConversationID           string `json:"ConversationID"`
	ResponseDescription      string `json:"ResponseDescription"`
}

func (d *Daraja) RegisterC2BCallback(payload C2BRegistrationPayload) (*C2BRegistrationResponse, *ErrorResponse) {
	secureResponse, err := performSecurePostRequest[*C2BRegistrationResponse](payload, endpointRegisterConfirmValidation, d)
	if err != nil {
		return nil, err
	}
	return secureResponse.Body, nil
}
func (d *Daraja) MakeC2BPayment(c2b C2BPayload) (*C2BResponse, *ErrorResponse) {
	c2b.CommandID = "CustomerPayBillOnline"
	secureResponse, err := performSecurePostRequest[*C2BResponse](c2b, endpointSimulatePmtC2B, d)
	if err != nil {
		return nil, err
	}
	return secureResponse.Body, nil
}
