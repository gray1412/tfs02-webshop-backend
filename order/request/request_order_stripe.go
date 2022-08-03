package request
type RequestPaymentStripe struct{
	PaymentId string `json:"payment_id"`
	Order RequestCreateOrder `json:"order"`
}