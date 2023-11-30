package darajaAuth

type ReversePayload struct {
	Initiator string `json:"Initiator"`
	PassKey string `json:"PassKey"`
	CommandID string `json:"CommandID"`
	TransactionID string `json:"TransactionID"`
	Amount string `json:"Amount"`
	ReceiverParty string `json:"ReceiverParty"`
	RecieverIdentifierType string `json:"RecieverIdentifierType"`
	ResultURL string `json:"ResultURL"`
	QueueTimeOutURL string `json:"QueueTimeOutURL"`
	Remarks string `json:"Remarks"`
	Occasion string `json:"Occasion"`
}

type ReversalResponse struct {
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ConversationID string `json:"ConversationID"`
	ResponseDescription string `json:"ResponseDescription"`
}

func (d *Daraja) ReverseTransaction(transaction ReversePayload, certPath string) (*ReversalResponse, *ErrorResponse){
	transaction.CommandID = "TransactionReversal"
	encryptedCredential, err := openSSlEncrypt(transaction.PassKey, certPath)
	if err != nil {
		return nil, &ErrorResponse{error: err}
	}
	transaction.PassKey = encryptedCredential

	secureResponse, errRes := performSecurePostRequest[*ReversalResponse](transaction, endpointReversal, d)
	if errRes != nil {
		return nil, errRes
	}
	return secureResponse.Body, nil
}