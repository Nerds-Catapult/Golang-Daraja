![Mpesa](https://github.com/shadmeoli/Golang-Daraja/assets/85517013/68faf2ac-976b-4af8-8658-d0285c17a3ad)

# GOLANG-DARAJA.

> An adapter to help connect to Daraja API for Transactions.

This repository contains a Go implementation for interacting with the Daraja API, a mobile payments platform. The implementation covers various aspects of the API, including authorization, transaction processing, and error handling.

## Table of Contents

1. [Introduction](#introduction)
2. [Files Overview](#files-overview)
3. [Usage](#usage)

## Introduction

The Daraja API Client is a Go library that simplifies integration with the Daraja API for mobile payments. It provides functionalities for handling authentication, making various types of transactions, and managing errors gracefully.

## Files Overview

### Auth.go

This file contains the implementation for the Daraja authorization process. It includes functions for obtaining and refreshing access tokens.

### b2b.go

The `b2b.go` file implements functions related to making Business-to-Business (B2B) payment requests using the Daraja API.

### b2c.go

In `b2c.go`, you'll find functions for initiating Business-to-Customer (B2C) payments through the Daraja API.

### C2bExpress.go

This file handles Express C2B (Customer-to-Business) transactions. It includes functions for processing C2B payments efficiently.

### c2b.go

The `c2b.go` file contains functions for initiating Customer-to-Business (C2B) payments using the Daraja API.

### Certencrypt.go

`Certencrypt.go` includes functions for encrypting data using SSL certificates. It supports loading certificates and performing encryption.

### Constants.go

The `Constants.go` file contains constant values used throughout the Daraja API Client, such as environment types.

### Daraja.go

This file, `Daraja.go`, defines the main Daraja struct, which acts as the central component for making API requests. It includes functions for authorization and managing the Daraja API instance.

### Model.go

`Model.go` contains structures that represent various data models used in the Daraja API, such as authorization details and transaction payloads.

### Network.go

The `Network.go` file handles making HTTP requests to the Daraja API, including error handling and response parsing.

### qrcodegen.go

In `qrcodegen.go`, you'll find functions for generating QR codes for Daraja API transactions.

### README.md

This file provides documentation for the Daraja API Client, including an overview, file descriptions, installation instructions, and usage guidelines.

### Reversal.go

The `Reversal.go` file implements functionality for reversing transactions made through the Daraja API.

### Struct2Map.go

`Struct2Map.go` contains a utility function for converting Go structs to maps.

# Usage

### Auth.go

The `Auth.go` file in the `darajaAuth` package contains the implementation for Daraja API authorization. It includes functions for obtaining access tokens required for subsequent API calls.

## Auth internals

1. [Structs](#structs)
   - [authResponse](#authresponse)
   - [Authorization](#authorization)
2. [Functions](#functions)
   - [NewAuthorization](#newauthorization)
3. [Usage Example](#usage-example)

## Structs

### authResponse

```go
type authResponse struct {
	AccessToken string `json:"access_token"` // The access token to be used in subsequent API calls
	ExpiresIn   string `json:"expires_in"`   // The number of seconds before the access token expires
}
```

Represents the JSON structure of the authorization response received from the Daraja API.

### Authorization

```go
type Authorization struct {
	authResponse
}
```

Represents the authorization details, including the access token and its expiration time.

## Functions

### NewAuthorization

```go
func NewAuthorization(consumerKey, consumerSecret string, env Environment) (*Authorization, error)
```

Creates a new `Authorization` instance by obtaining an access token from the Daraja API.

- Parameters:

  - `consumerKey`: The consumer key for authenticating with the Daraja API.
  - `consumerSecret`: The consumer secret for authenticating with the Daraja API.
  - `env`: The environment type (`ENVIROMENT_SANDBOX` or `ENVIROMENT_PRODUCTION`).

- Returns:
  - `*Authorization`: The authorized instance with the access token.
  - `error`: An error if the authorization process fails.

## Usage Example

```go
consumerKey := "your-consumer-key"
consumerSecret := "your-consumer-secret"
env := daraja.ENVIRONMENT_SANDBOX

auth, err := darajaAuth.NewAuthorization(consumerKey, consumerSecret, env)
if err != nil {
    // Handle authorization error
    fmt.Println("Authorization error:", err)
    return
}

// Access the access token and expiration time
accessToken := auth.AccessToken
expiresIn := auth.ExpiresIn
fmt.Printf("Access Token: %s\nExpires In: %s seconds\n", accessToken, expiresIn)
```

This example demonstrates how to create a new `Authorization` instance by providing the consumer key, consumer secret, and environment type. The access token and expiration time can then be accessed for use in subsequent Daraja API calls.

---

# b2b.go

The `b2b.go` file in the `darajaAuth` package contains the implementation for making Business-to-Business (B2B) payment requests using the Daraja API.

## b2b internals

1. [Structs](#structs)
   - [B2BPaymentPayload](#b2bpaymentpayload)
   - [B2BPaymentResponse](#b2bpaymentresponse)
2. [Functions](#functions)
   - [MakeB2BPaymentRequest](#makeb2bpaymentrequest)
3. [Usage Example](#usage-example)

## Structs

### B2BPaymentPayload

```go
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
```

Represents the payload structure required for initiating a B2B payment through the Daraja API.

### B2BPaymentResponse

```go
type B2BPaymentResponse struct {
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ConversationID           string `json:"ConversationID"`
	ResponseDescription      string `json:"ResponseDescription"`
}
```

Represents the response structure received after making a B2B payment request.

## Functions

### MakeB2BPaymentRequest

```go
func (d *Daraja) MakeB2BPaymentRequest(b2c B2BPaymentPayload, certPath string) (*B2CPaymentResponse, *ErrorResponse)
```

Initiates a Business-to-Business (B2B) payment request using the Daraja API.

- Parameters:

  - `b2c`: The B2B payment payload containing transaction details.
  - `certPath`: The file path to the SSL certificate used for encryption.

- Returns:
  - `*B2CPaymentResponse`: The B2B payment response containing transaction details.
  - `*ErrorResponse`: An error response if the B2B payment request fails.

## Usage Example

```go
// Initialize Daraja API instance
daraja := daraja.NewDaraja("your-consumer-key", "your-consumer-secret", daraja.ENVIRONMENT_SANDBOX)

// Create B2B payment payloadCustomize the documentation based on your specific implementation details and use cases. This template provides a starting point for explaining the purpose and usage of the `b2b.go` file.

b2bPayload := daraja.B2BPaymentPayload{
	InitiatorName:   "InitiatorName",
	Passkey:         "YourSecurityCredential",
	CommandID:       "BusinessPayment",
	Amount:          "1000",
	PartyA:          "PartyA",
	PartyB:          "PartyB",
	Remarks:         "Payment remarks",
	QueueTimeOutURL: "https://your-queue-timeout-url.com",
	ResultURL:       "https://your-result-url.com",
	Occasion:        "Payment occasion",
}

// Make B2B payment request
response, err := daraja.MakeB2BPaymentRequest(b2bPayload, "path/to/ssl/certificate.pem")
if err != nil {
	// Handle B2B payment request error
	fmt.Println("B2B payment request error:", err)
	return
}

// Process the B2B payment response
fmt.Println("B2B Payment Response:")
fmt.Println("Originator Conversation ID:", response.OriginatorConversationID)
fmt.Println("Conversation ID:", response.ConversationID)
fmt.Println("Response Description:", response.ResponseDescription)
```

This example demonstrates how to make a B2B payment request using the Daraja API. Customize the payload details and handle the response accordingly in your application.

---

### b2c.go

The `b2c.go` file in the `darajaAuth` package contains the implementation for making Business-to-Customer (B2C) payment requests using the Daraja API.

## b2c internals

1. [Enums](#enums)
   - [B2CCommandID](#b2ccommandid)
2. [Structs](#structs)
   - [B2CPaymentPayload](#b2cpaymentpayload)
   - [B2CPaymentResponse](#b2cpaymentresponse)
3. [Functions](#functions)
   - [MakeB2CPaymentRequest](#makeb2cpaymentrequest)
4. [Usage Example](#usage-example)

## Enums

### B2CCommandID

```go
type B2CCommandID string

const (
	B2CCommandIDSalary           B2CCommandID = "SalaryPayment"
	B2CCommandIDBusinessPayment  B2CCommandID = "BusinessPayment"
	B2CCommandIDPromotionPayment B2CCommandID = "PromotionPayment"
)
```

Enumerates the possible command IDs for Business-to-Customer (B2C) payments.

## Structs

### B2CPaymentPayload

```go
type B2CPaymentPayload struct {
	InitiatorName   string       `json:"InitiatorName"`
	PassKey         string       `json:"PassKey"`
	CommandID       B2CCommandID `json:"CommandID"`
	Amount          string       `json:"Amount"`
	PartyA          string       `json:"PartyA"`
	PartyB          string       `json:"PartyB"`
	Remarks         string       `json:"Remarks"`
	QueueTimeOutURL string       `json:"QueueTimeOutURL"`
	ResultURL       string       `json:"ResultURL"`
	Occasion        string       `json:"Occasion"`
}
```

Represents the payload structure required for initiating a Business-to-Customer (B2C) payment through the Daraja API.

### B2CPaymentResponse

```go
type B2CPaymentResponse struct {
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ResponseCode             string `json:"ResponseCode"`
	ResponseDescription      string `json:"ResponseDescription"`
	ConversationID           string `json:"ConversationID"`
}
```

Represents the response structure received after making a B2C payment request.

## Functions

### MakeB2CPaymentRequest

```go
func (d *Daraja) MakeB2CPaymentRequest(b2c B2CPaymentPayload, certPath string) (*B2CPaymentResponse, *ErrorResponse)
```

Initiates a Business-to-Customer (B2C) payment request using the Daraja API.

- Parameters:

  - `b2c`: The B2C payment payload containing transaction details.
  - `certPath`: The file path to the SSL certificate used for encryption.

- Returns:
  - `*B2CPaymentResponse`: The B2C payment response containing transaction details.
  - `*ErrorResponse`: An error response if the B2C payment request fails.

## Usage Example

```go
// Initialize Daraja API instance
daraja := daraja.NewDaraja("your-consumer-key", "your-consumer-secret", daraja.ENVIRONMENT_SANDBOX)

// Create B2C payment payload
b2cPayload := daraja.B2CPaymentPayload{
	InitiatorName:   "InitiatorName",
	PassKey:         "YourPassKey",
	CommandID:       daraja.B2CCommandIDBusinessPayment,
	Amount:          "1000",
	PartyA:          "PartyA",
	PartyB:          "PartyB",
	Remarks:         "Payment remarks",
	QueueTimeOutURL: "https://your-queue-timeout-url.com",
	ResultURL:       "https://your-result-url.com",
	Occasion:        "Payment occasion",
}

// Make B2C payment request
response, err := daraja.MakeB2CPaymentRequest(b2cPayload, "path/to/ssl/certificate.pem")
if err != nil {
	// Handle B2C payment request error
	fmt.Println("B2C payment request error:", err)
	return
}

// Process the B2C payment response
fmt.Println("B2C Payment Response:")
fmt.Println("Originator Conversation ID:", response.OriginatorConversationID)
fmt.Println("Response Code:", response.ResponseCode)
fmt.Println("Response Description:", response.ResponseDescription)
fmt.Println("Conversation ID:", response.ConversationID)
```

This example demonstrates how to make a B2C payment request using the Daraja API. Customize the payload details and handle the response accordingly in your application.

# c2b.go

The `c2b.go` file in the `darajaAuth` package contains the implementation for making Consumer-to-Business (C2B) payment requests using the Daraja API.

## c2b internals

1. [Structs](#structs)
   - [C2BPayload](#c2bpayload)
   - [C2BResponse](#c2bresponse)
   - [C2BRegistrationPayload](#c2bregistrationpayload)
   - [C2BRegistrationResponse](#c2bregistrationresponse)
2. [Functions](#functions)
   - [RegisterC2BCallback](#registerc2bcallback)
   - [MakeC2BPayment](#makec2bpayment)
3. [Usage Example](#usage-example)

## Structs

### C2BPayload

```go
type C2BPayload struct {
	ShortCode     string `json:"ShortCode"`
	CommandID     string `json:"CommandID"`
	Amount        string `json:"Amount"`
	Msisdn        string `json:"Msisdn"`
	BillRefNumber string `json:"BillRefNumber"`
}
```

Represents the payload structure required for initiating a Consumer-to-Business (C2B) payment through the Daraja API.

### C2BResponse

```go
type C2BResponse struct {
	OriginitatorConversationID string `json:"OriginitatorConversationID"`
	ResponseDescription        string `json:"ResponseDescription"`
	ConversationID             string `json:"ConversationID"`
}
```

Represents the response structure received after making a C2B payment request.

### C2BRegistrationPayload

```go
type C2BRegistrationPayload struct {
	ValidationURL   string `json:"ValidationURL"`
	ConfirmationURL string `json:"ConfirmationURL"`
	ResponseType    string `json:"ResponseType"`
	ShortCode       string `json:"ShortCode"`
}
```

Represents the payload structure required for registering C2B callback URLs.

### C2BRegistrationResponse

```go
type C2BRegistrationResponse struct {
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ConversationID           string `json:"ConversationID"`
	ResponseDescription      string `json:"ResponseDescription"`
}
```

Represents the response structure received after registering C2B callback URLs.

## Functions

### RegisterC2BCallback

```go
func (d *Daraja) RegisterC2BCallback(payload C2BRegistrationPayload) (*C2BRegistrationResponse, *ErrorResponse)
```

Registers C2B callback URLs with the Daraja API.

- Parameters:

  - `payload`: The C2B registration payload containing callback URL details.

- Returns:
  - `*C2BRegistrationResponse`: The C2B registration response containing registration details.
  - `*ErrorResponse`: An error response if the registration process fails.

### MakeC2BPayment

```go
func (d *Daraja) MakeC2BPayment(c2b C2BPayload) (*C2BResponse, *ErrorResponse)
```

Initiates a Consumer-to-Business (C2B) payment request using the Daraja API.

- Parameters:

  - `c2b`: The C2B payment payload containing transaction details.

- Returns:
  - `*C2BResponse`: The C2B payment response containing transaction details.
  - `*ErrorResponse`: An error response if the C2B payment request fails.

## Usage Example

```go
// Initialize Daraja API instance
daraja := daraja.NewDaraja("your-consumer-key", "your-consumer-secret", daraja.ENVIRONMENT_SANDBOX)

// Register C2B callback URLs
c2bRegistrationPayload := daraja.C2BRegistrationPayload{
	ValidationURL:   "https://your-validation-url.com",
	ConfirmationURL: "https://your-confirmation-url.com",
	ResponseType:    "Complete",
	ShortCode:       "YourShortCode",
}

registrationResponse, err := daraja.RegisterC2BCallback(c2bRegistrationPayload)
if err != nil {
	// Handle C2B registration error
	fmt.Println("C2B registration error:", err)
	return
}

// Make C2B payment request
c2bPayload := daraja.C2BPayload{
	ShortCode:     "YourShortCode",
	CommandID:     daraja.C2BCommandIDCustomerPayBillOnline,
	Amount:        "1000",
	Msisdn:        "CustomerMsisdn",
	BillRefNumber: "BillReferenceNumber",
}

paymentResponse, err := daraja.MakeC2BPayment(c2bPayload)
if err != nil {
	// Handle C2B payment request error
	fmt.Println("C2B payment request error:", err)
	return
}

// Process the C2B payment response
fmt.Println("C2B Payment Response:")
fmt.Println("Originator Conversation ID:", paymentResponse.OriginitatorConversationID)
fmt.Println("Response Description:", paymentResponse.ResponseDescription)
fmt.Println("Conversation ID:", paymentResponse.ConversationID)
```

This example demonstrates how to register C2B callback URLs and make a C2B payment request using the Daraja API. Customize the payload details and handle the response accordingly in your application.

# C2bExpress.go

The `C2bExpress.go` file in the `darajaAuth` package contains the implementation for making Express (STK) Push requests and handling Express callback functionality using the Daraja API.

## C2bExpress internals

1. [Constants](#constants)
2. [Types](#types)
   - [ExpressCallbackFunc](#expresscallbackfunc)
   - [LipaNaMpesaPayload](#lipanampesapayload)
   - [LipaNaMpesaResponse](#lipanampesaresponse)
   - [STKPushStatusPayload](#stkpushstatuspayload)
   - [STKPushStatusResponse](#stkpushstatusresponse)
   - [CallBackResponse](#callbackresponse)
3. [Functions](#functions)
   - [MakeSTKPushRequest](#makestkpushrequest)
   - [MapExpressGinCallBack](#mapexpressgincallback)
4. [Usage Example](#usage-example)

## Constants

### ExpressDefaultCallBackURL

```go
const ExpressDefaultCallBackURL = "daraja-payments/mpesa"
```

The default callback URL for Express (STK) Push transactions.

## Types

### ExpressCallbackFunc

```go
type ExpressCallbackFunc func(response *CallBackResponse, request http.Request, err error)
```

Defines a callback function for handling Express (STK) Push transaction responses.

### LipaNaMpesaPayload

```go
type LipaNaMpesaPayload struct {
	BusinessShortCode string `json:"BusinessShortCode"`
	Password          string `json:"Password"`
	Timestamp         string `json:"Timestamp"`
	TransactionType   string `json:"TransactionType"`
	Amount            string `json:"Amount"`
	PartyA            string `json:"PartyA"`
	PartyB            string `json:"PartyB"`
	PhoneNumber       string `json:"PhoneNumber"`
	CallBackURL       string `json:"CallBackURL"`
	AccountReference  string `json:"AccountReference"`
	TransactionDesc   string `json:"TransactionDesc"`
}
```

Represents the payload structure required for initiating an Express (STK) Push transaction through the Daraja API.

### LipaNaMpesaResponse

```go
type LipaNaMpesaResponse struct {
	MerchantRequestID   string `json:"MerchantRequestID"`
	CheckoutRequestID   string `json:"CheckoutRequestID"`
	ResponseCode        string `json:"ResponseCode"`
	ResponseDescription string `json:"ResponseDescription"`
	CustomerMessage     string `json:"CustomerMessage"`
}
```

Represents the response structure received after making an Express (STK) Push transaction request.

### STKPushStatusPayload

```go
type STKPushStatusPayload struct {
	BusinessShortCode string `json:"BusinessShortCode"`
	Password          string `json:"Password"`
	Timestamp         string `json:"Timestamp"`
	CheckoutRequestID string `json:"CheckoutRequestID"`
}
```

Represents the payload structure required for checking the status of an Express (STK) Push transaction through the Daraja API.

### STKPushStatusResponse

```go
type STKPushStatusResponse struct {
	MerchantRequestID   string `json:"MerchantRequestID"`
	CheckoutRequestID   string `json:"CheckoutRequestID"`
	ResponseCode        string `json:"ResponseCode"`
	ResponseDescription string `json:"ResponseDescription"`
	ResultDesc          string `json:"ResultDesc"`
	ResultCode          string `json:"ResultCode"`
}
```

Represents the response structure received after checking the status of an Express (STK) Push transaction.

### CallBackResponse

```go
type CallBackResponse struct {
	Body struct {
		StkCallBack struct {
			MerchantRequestID string `json:"MerchantRequestID"`
			CheckoutRequestID string `json:"CheckoutRequestID"`
			ResultCode        int    `json:"ResultCode"`
			CallbackMetaData  struct {
				Item []struct {
					Name  string      `json:"Name"`
					Value interface{} `json:"value"`
				} `json:"Item"`
			} `json:"CallbackMetaData"`
		} `json:"StkCallBack"`
	} `json:"Body"`
}
```

Represents the response structure received in the callback after an Express (STK) Push transaction.

## Functions

### MakeSTKPushRequest

```go
func (d *Daraja) MakeSTKPushRequest(mpesaConfig LipaNaMpesaPayload) (*LipaNaMpesaResponse, *ErrorResponse)
```

Initiates an Express (STK) Push transaction request using the Daraja API.

- Parameters:

  - `mpesaConfig`: The Express (STK) Push payload containing transaction details.

- Returns:
  - `*LipaNaMpesaResponse`: The Express (STK) Push response containing transaction details.
  - `*ErrorResponse`: An error response if the Express (STK) Push request fails.

### MapExpressGinCallBack

```go
func MapExpressGinCallBack(gingroup *gin.RouterGroup, callBackUrl string, callback ExpressCallbackFunc)
```

Maps an Express (STK) Push callback endpoint in a Gin router group.

- Parameters:
  - `gingroup`: The Gin router group.
  - `callBackUrl`: The callback URL for Express (STK) Push transactions.
  - `callback`: The callback function to handle the Express (STK) Push transaction response.

## Usage Example

```go
// Initialize Daraja API instance
daraja := daraja.NewDaraja("your-consumer-key", "your-consumer-secret", daraja.ENVIRONMENT_SANDBOX)

// Create Express (STK) Push payload
mpesaConfig := daraja.LipaNaMpesaPayload{
	BusinessShortCode: "YourBusinessShortCode",
	Password:          "YourPassword",
	TransactionType:   "CustomerPayBillOnline",
	Amount:            "1000",
	PartyA:            "YourPartyA",
	PartyB:            "YourPartyB",
	PhoneNumber:       "CustomerPhoneNumber",
	CallBackURL:       "https://your-callback-url.com",
	AccountReference:  "YourAccountReference",
	TransactionDesc:   "YourTransactionDescription",
}

// Make Express (STK) Push request
response, err := daraja.MakeSTKPushRequest(mpesaConfig)
if err != nil {
	// Handle Express (STK) Push request error
	fmt.Println("Express (STK) Push request error:", err)
	return
}

// Process the Express (STK) Push response
fmt.Println("Express (STK) Push Response:")
fmt.Println("Merchant Request ID:", response.MerchantRequestID)
fmt.Println("Checkout Request ID:", response.CheckoutRequestID)
fmt.Println("Response Code:", response.ResponseCode)
fmt.Println("Response Description:", response.ResponseDescription)
fmt.Println("Customer Message:", response.CustomerMessage)

// Map Express (STK) Push callback endpoint
daraja.MapExpressGinCallBack(yourGinRouterGroup, "/your-callback-endpoint", yourCallbackFunction)
```

This example demonstrates how to make an Express (STK) Push request, handle the response, and map the callback endpoint using the Daraja API. Customize the

payload details and handle the response accordingly in your application.

# Certencrypt.go

The `Certencrypt.go` file in the `darajaAuth` package contains functions for encrypting data using an X.509 certificate.

## Certencrypt Internals

1. [Types](#types)
   - [certificationError](#certificationerror)
2. [Functions](#functions)
   - [openSSlEncrypt](#opensslencrypt)
   - [loadCertificate](#loadcertificate)
3. [Usage Example](#usage-example)

## Types

### certificationError

```go
type certificationError struct {
	Context string
	Err     error
}
```

`certificationError` represents a custom error type used for handling certification-related errors. It provides additional context to the standard Go error.

#### Methods

- `Error() string`: Implements the error interface for `certificationError` and returns a formatted error message.

## Functions

### openSSlEncrypt

```go
func openSSlEncrypt(data, certPath string) (string, error)
```

`openSSlEncrypt` encrypts the provided data using the public key from the X.509 certificate located at the specified file path.

#### Parameters

- `data`: The data to be encrypted.
- `certPath`: The file path to the X.509 certificate.

#### Returns

- `string`: The base64-encoded encrypted data.
- `error`: An error indicating if the encryption process failed.

### loadCertificate

```go
func loadCertificate(certPath string) (*x509.Certificate, error)
```

`loadCertificate` loads and parses the X.509 certificate from the provided file path.

#### Parameters

- `certPath`: The file path to the X.509 certificate.

#### Returns

- `*x509.Certificate`: The parsed X.509 certificate.
- `error`: An error indicating if the loading or parsing of the certificate failed.

## Usage Example

```go
// Example usage of Certencrypt.go

// Load Daraja API and other necessary packages

// Define the data to be encrypted
dataToEncrypt := "SensitiveData123"

// Specify the path to the X.509 certificate
certPath := "/path/to/certificate.pem"

// Encrypt the data using the X.509 certificate
encryptedData, err := darajaAuth.openSSlEncrypt(dataToEncrypt, certPath)
if err != nil {
	// Handle encryption error
	fmt.Println("Encryption failed:", err)
	return
}

// Process the encrypted data
fmt.Println("Encrypted Data:", encryptedData)
```

This example demonstrates how to use the encryption functions provided in `Certencrypt.go`. Customize the data and certificate path according to your application needs.

# Daraja.go

The `Daraja.go` file in the `darajaAuth` package contains the implementation of the Daraja struct, which represents a mobile money API client. It includes functions for authorization and handling various payment requests.

## Daraja.go internals

1. [Types](#types)
   - [darajaAuthorizationError](#darajaauthorizationerror)
2. [Constants](#constants)
   - [Environment](#environment)
3. [Structs](#structs)
   - [Daraja](#daraja)
4. [Interfaces](#interfaces)
   - [darajaApiImpl](#darajaapiimpl)
5. [Functions](#functions)
   - [NewDaraja](#newdaraja)
   - [Authorize](#authorize)
6. [Usage Example](#usage-example)

## Types

### darajaAuthorizationError

```go
type darajaAuthorizationError struct {
	Context string
	Err     error
}
```

`darajaAuthorizationError` is a custom error type that extends the default error. It provides additional context for errors related to Daraja authorization.

#### Methods

- `Error() string`: Implements the error interface for `darajaAuthorizationError` and returns a formatted error message.

## Constants

### Environment

```go
const (
	ENVIROMENT_SANDBOX    = "sandbox"
	ENVIROMENT_PRODUCTION = "production"
)
```

`Environment` is an enumeration representing the Daraja environment, with possible values of "sandbox" and "production."

## Structs

### Daraja

```go
type Daraja struct {
	authorization  Authorization
	environment    Environment
	nextAuthTime   time.Time
	ConsumerKey    string
	ConsumerSecret string
}
```

`Daraja` is a struct representing the Daraja mobile money API client. It contains fields for authorization details, environment, and API credentials.

## Interfaces

### darajaApiImpl

```go
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
```

`darajaApiImpl` is an interface defining the methods expected in the Daraja API implementation.

## Functions

### NewDaraja

```go
func NewDaraja(consumerKey, consumerSecret string, env Environment) *Daraja
```

`NewDaraja` creates a new instance of the Daraja API client with the specified consumer key, consumer secret, and environment.

#### Parameters

- `consumerKey`: The consumer key for API authentication.
- `consumerSecret`: The consumer secret for API authentication.
- `env`: The environment (sandbox or production) in which the API client operates.

#### Returns

- `*Daraja`: The new instance of the Daraja API client.

### Authorize

```go
func (d *Daraja) Authorize() (*Authorization, error)
```

`Authorize` initiates the authorization process for the Daraja API client. It returns an authorization token and updates the next authorization time.

#### Returns

- `*Authorization`: The authorization details.
- `error`: An error indicating if the authorization process failed.

## Usage Example

```go
// Example usage of Daraja.go

// Load Daraja API and other necessary packages

// Initialize Daraja API client
daraja := darajaAuth.NewDaraja("your_consumer_key", "your_consumer_secret", darajaAuth.ENVIROMENT_SANDBOX)

// Authorize the Daraja API client
auth, err := daraja.Authorize()
if err != nil {
	// Handle authorization error
	fmt.Println("Authorization failed:", err)
	return
}

// Perform other Daraja API operations using the authorized client
// ...
```

This example demonstrates how to create a new Daraja API client, authorize it, and perform various operations. Customize the consumer key, consumer secret, and environment according to your application needs.

# Reversal.go

The `Reversal.go` file in the `darajaAuth` package contains the implementation related to transaction reversals in the Daraja mobile money API client.

## Reversal internals

1. [Types](#types)
   - [ReversePayload](#reversepayload)
   - [ReversalResponse](#reversalresponse)
2. [Functions](#functions)
   - [ReverseTransaction](#reversetransaction)
3. [Usage Example](#usage-example)

## Types

### ReversePayload

```go
type ReversePayload struct {
	TransactionID string `json:"TransactionID"`
	Occasion      string `json:"Occasion"`
}
```

`ReversePayload` represents the payload required for initiating a transaction reversal. It includes fields such as `TransactionID` and `Occasion`.

### ReversalResponse

```go
type ReversalResponse struct {
	OriginatorConversationID string `json:"OriginatorConversationID"`
	ResponseDescription      string `json:"ResponseDescription"`
	ConversationID           string `json:"ConversationID"`
}
```

`ReversalResponse` is the response structure for a transaction reversal. It includes fields such as `OriginatorConversationID`, `ResponseDescription`, and `ConversationID`.

## Functions

### ReverseTransaction

```go
func (d *Daraja) ReverseTransaction(transaction ReversePayload) (*ReversalResponse, *ErrorResponse)
```

`ReverseTransaction` initiates a transaction reversal using the provided payload.

#### Parameters

- `transaction`: The payload containing details necessary for the reversal, such as `TransactionID` and `Occasion`.

#### Returns

- `*ReversalResponse`: The response containing details of the reversal.
- `*ErrorResponse`: An error indicating if the reversal process failed.

## Usage Example

```go
// Example usage of Reversal.go

// Load Daraja API and other necessary packages

// Initialize Daraja API client
daraja := darajaAuth.NewDaraja("your_consumer_key", "your_consumer_secret", darajaAuth.ENVIROMENT_SANDBOX)

// Prepare reversal payload
reversalPayload := darajaAuth.ReversePayload{
	TransactionID: "123456789",
	Occasion:      "Reversal for incorrect transaction",
}

// Perform transaction reversal
reversalResponse, err := daraja.ReverseTransaction(reversalPayload)
if err != nil {
	// Handle reversal error
	fmt.Println("Reversal failed:", err)
	return
}

// Process reversal response
// ...
```

This example demonstrates how to use the `ReverseTransaction` function to initiate a transaction reversal. Customize the consumer key, consumer secret, and reversal payload according to your application needs.

# qrcodegen.go

The `qrcodegen.go` file in the `darajaAuth` package contains the implementation of generating QR codes for the Daraja mobile money API client.

## qrcodegen internals

1. [Types](#types)
   - [QrPayload](#qrpayload)
   - [QrResponse](#qrresponse)
2. [Functions](#functions)
   - [MakeQRCodeRequest](#makeqrcoderequest)
3. [Usage Example](#usage-example)

## Types

### QrPayload

```go
type QrPayload struct {
	TillNumber string `json:"TillNumber"`
	Amount     string `json:"Amount"`
	AccountRef string `json:"AccountRef"`
}
```

`QrPayload` represents the payload required for generating a QR code. It includes fields such as `TillNumber`, `Amount`, and `AccountRef`.

### QrResponse

```go
type QrResponse struct {
	QRData string `json:"QRData"`
}
```

`QrResponse` is the response structure for a generated QR code. It includes the `QRData` field containing the QR code information.

## Functions

### MakeQRCodeRequest

```go
func (d *Daraja) MakeQRCodeRequest(payload QrPayload) (*QrResponse, *ErrorResponse)
```

`MakeQRCodeRequest` generates a QR code using the provided payload.

#### Parameters

- `payload`: The payload containing details necessary for generating the QR code, such as `TillNumber`, `Amount`, and `AccountRef`.

#### Returns

- `*QrResponse`: The response containing the generated QR code data.
- `*ErrorResponse`: An error indicating if the QR code generation process failed.

## Usage Example

```go
// Example usage of qrcodegen.go

// Load Daraja API and other necessary packages

// Initialize Daraja API client
daraja := darajaAuth.NewDaraja("your_consumer_key", "your_consumer_secret", darajaAuth.ENVIROMENT_SANDBOX)

// Prepare QR code payload
qrPayload := darajaAuth.QrPayload{
	TillNumber: "123456",
	Amount:     "1000",
	AccountRef: "Invoice123",
}

// Generate QR code
qrResponse, err := daraja.MakeQRCodeRequest(qrPayload)
if err != nil {
	// Handle QR code generation error
	fmt.Println("QR code generation failed:", err)
	return
}

// Process QR code response
// ...
```

This example demonstrates how to use the `MakeQRCodeRequest` function to generate a QR code. Customize the consumer key, consumer secret, and QR code payload according to your application needs.
