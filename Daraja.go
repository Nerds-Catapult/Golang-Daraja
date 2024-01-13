package darajaAuth

import (
	"time"
)

// darajaAuthorizationError is a custom error type that extends the default error.
type darajaAuthorizationError struct {
	Context string
	Err     error
}

func (e *darajaAuthorizationError) Error() string {
	return e.Context + ": " + e.Err.Error()
}

const (
	ENVIROMENT_SANDBOX    = "sandbox"
	ENVIROMENT_PRODUCTION = "production"
)

type Environment string

type Daraja struct {
	authorization  Authorization
	environment    Environment
	nextAuthTime   time.Time
	ConsumerKey    string
	ConsumerSecret string
}

type darajaApiImpl interface {
	Authorize() (*Authorization, error)
	ReverseTransaction(transaction ReversePayload) (*ReversalResponse, *ErrorResponse)
	MakeSTKPushRequest(mpesaConfig LipaNaMpesaPayload) (*LipaNaMpesaResponse, *ErrorResponse)
	MakeB2BPaymentRequest(b2bPayment B2BPaymentPayload) (*B2BPaymentResponse, *ErrorResponse)
	MakeB2CPaymentRequest(b2CPayment B2CPaymentPayload) (*B2CPaymentResponse, *ErrorResponse)
	MakeQRCodeRequest(payload QrPayload) (*QrResponse, *ErrorResponse)
	MakeC2BPayload(c2b C2BPayload) (*C2BResponse, *ErrorResponse)
	MakeC2BPaymentV2(c2b C2BPayload) (*C2BResponse, *ErrorResponse)
}

var darajaAPI *Daraja

func NewDaraja(consumerKey, consumerSecret string, env Environment) *Daraja {
	if darajaAPI == nil {
		darajaAPI = &Daraja{
			ConsumerKey:    consumerKey,
			ConsumerSecret: consumerSecret,
			environment:    env,
		}
	}
	return darajaAPI
}

func (d *Daraja) Authorize() (*Authorization, error) {
	authTimeStart := time.Now()
	auth, err := NewAuthorization(d.ConsumerKey, d.ConsumerSecret, d.environment)
	if err != nil {
		return nil, &darajaAuthorizationError{Context: "authorization failed", Err: err}
	}

	expiry, err := time.ParseDuration(auth.ExpiresIn + "s")
	if err != nil {
		return nil, &darajaAuthorizationError{Context: "failed to parse authorization expiry", Err: err}
	}

	d.nextAuthTime = authTimeStart.Add(expiry)
	d.authorization = *auth

	return auth, nil
}
