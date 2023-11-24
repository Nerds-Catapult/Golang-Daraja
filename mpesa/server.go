package mpesa

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"strings"
	"time"
)

func CreateToken() (string, error) {
	consumerSecret := "consumer_key"
	consumerKey := "consumer key"
	auth := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%s:%s", consumerKey, consumerSecret)))
	url := "https://sandbox.safaricom.co.ke/oauth/v1/generate?grant_type=client_credentials"
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return "", err
	}
	req.Header.Add("Authorization", "Basic "+auth)
	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", err
	}
	accessToken := string(body)
	return accessToken, nil
}

func PushStk(phoneNumber string, amount int)(string, error){
	if phoneNumber == ""{
		return "", fmt.Errorf("phone Number is required")
	}
	shortCode := 174379
	phone := phoneNumber[len(phoneNumber)-9:]
	passkey := "bfb279f9aa9bdbcf158e97dd71a467cd2e0c893059b10f78e6b72ada1ed2c919"

	token, err := CreateToken()
	if err != nil {
		return "", fmt.Errorf("failed to create Daraja token: %v", err)
	}
	timestamp := time.Now().UTC().Format("20060102150405")
	password := base64.StdEncoding.EncodeToString([]byte(fmt.Sprintf("%d%s%s", shortCode, passkey,timestamp)))
	payload := fmt.Sprintf(`{
		"BusinessShortCode": %d,
			"Password": "%s",
			"Timestamp": "%s",
			"TransactionType": "CustomerPayBillOnline",
			"Amount": %d,
			"PartyA": "%s",
			"PartyB": "%d",
			"PhoneNumber": "254%s",
			"CallBackURL": "https://goose-merry-mollusk.ngrok-free.app/api/callback",
			"AccountReference": "Test",
			"TransactionDesc": "Test"
}`, shortCode, password, timestamp, amount, phone, shortCode)
	url := "https://sandbox.safaricom.co.ke/mpesa/stkpush/v1/processrequest"
	req, err := http.NewRequest("POST", url, strings.NewReader(payload))
	if err != nil {
		return "", fmt.Errorf("failed to create stk push: %v", err)
	}
	req.Header.Add("Authorization", "Bearer "+token)
	req.Header.Add("Content-Type", "application/json")

	resp, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", fmt.Errorf("failed to read stk push response: %v", err)
	}
	defer func(Body io.ReadCloser) {
		err := Body.Close()
		if err != nil {

		}
	}(resp.Body)
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return "", fmt.Errorf("failed to read STK push response: %v", err)
	}
	return string(body), nil
}