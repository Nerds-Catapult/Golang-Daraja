package darajaAuth

import (
	"encoding/base64"
	"github.com/gin-gonic/gin"
	"net/http"
	"time"
)

const (
	ExpressDefaultCallBackURL = "daraja-payments/mpesa"
)

type ExpressCallbackFunc func(response *CallBackResponse, request http.Request, err error)

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

type LipaNaMpesaResponse struct {
	MerchantRequestID string `json:"MerchantRequestID"`
	CheckoutRequestID string `json:"CheckoutRequestID"`
	ResponseCode string `json:"ResponseCode"`
	ResponseDescription string `json:"ResponseDescription"`
	CustomerMessage string `json:"CustomerMessage"`
}

type STKPushStatusPayload struct {
	BusinessShortCode string `json:"BusinessShortCode"`
	Password string `json:"Password"`
	Timestamp string `json:"Timestamp"`
	CheckoutRequestID string `json:"CheckoutRequestID"`
}

type STKPushStatusResponse struct {
	MerchantRequestID string `json:"MerchantRequestID"`
	CheckoutRequestID string `json:"CheckoutRequestID"`
	ResponseCode string `json:"ResponseCode"`
	ResponseDescription string `json:"ResponseDescription"`
	ResultDesc string `json:"ResultDesc"`
	ResultCode string `json:"ResultCode"`
}
type CallBackResponse struct {
	Body struct {
		StkCallBack struct {
			MerchantRequestID string `json:"MerchantRequestID"`
			CheckoutRequestID string `json:"CheckoutRequestID"`
			ResultCode int `json:"ResultCode"`
			CallbackMetaData struct{
				Item []struct{
					Name string `json:"Name"`
					Value interface{} `json:"value"`
				} `json:"Item"`
			} `json:"CallbackMetaData"`
		} `json:"StkCallBack"`
	} `json:"Body"`
}

func (d *Daraja) MakeSTKPushRequest(mpesaConfig LipaNaMpesaPayload) (*LipaNaMpesaResponse, *ErrorResponse){
	t := time.Now()
	layout := "20060102150405"
	timestamp := t.Format(layout)
	password := base64.StdEncoding.EncodeToString([]byte(mpesaConfig.BusinessShortCode + mpesaConfig.Password + timestamp))
	mpesaConfig.Timestamp = timestamp
	mpesaConfig.Password = password

	secureResponse, err := performSecurePostRequest[LipaNaMpesaResponse](mpesaConfig, endpointLipaNaMpesa, d)
	if err != nil {
		return nil, err
	}
	return &secureResponse.Body, nil
}

func MapExpressGinCallBack(gingroup *gin.RouterGroup, callBackUrl string, callback ExpressCallbackFunc) {
	gingroup.POST(callBackUrl, func(context *gin.Context){
		var callbackResponse CallBackResponse
		err := context.BindJSON(&callbackResponse)
		if err != nil {
			callback(nil, *context.Request, err)
		}
		callback(&callbackResponse, *context.Request, nil)
		context.JSON(http.StatusOK, map[string]string{"message": "success"})
	})
}