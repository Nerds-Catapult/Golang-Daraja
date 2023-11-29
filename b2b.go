package darajaAuth
//b2b means business to business

type B2BPaymentPayload struct {
	InitiatorName              string  `json:"Initiator"`
	Passkey 				  string  `json:"SecurityCredential"`
	CommandID                  string  `json:"CommandID"`
	Amount					 string `json:"Amount"`
	PartyA                     string  `json:"PartyA"`
	PartyB					 string  `json:"PartyB"`
	Remarks                    string  `json:"Remarks"`
	QueueTimeOutURL            string  `json:"QueueTimeOutURL"`
	ResultURL                  string  `json:"ResultURL"`
	Occasion                   string  `json:"Occasion"`
}
