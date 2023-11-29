package darajaAuth


import (
	"time"
)

const (
	ENVIROMENT_SANDBOX = "sandbox"
	ENVIROMENT_PRODUCTION = "production"
)

type Environment string

type Daraja struct {
	authorization Authorization
	environment Environment
	nextAuthorizationTime time.Time
	ConsumerKey string
	ConsumerSecret string
}

type darajaApiInterface interface {
	Authorize() (*Authorization, error)

}