package darajaAuth

//business to customer

type B2CCommandID string

const (
	B2CCommandIDSalary B2CCommandID = "SalaryPayment"
	B2CCommandIDBusinessPayment B2CCommandID = "BusinessPayment"
	B2CCommandIDPromotionPayment B2CCommandID = "PromotionPayment"
)