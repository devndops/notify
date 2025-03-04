package models

// OrderConfirmationData holds data for an order confirmation email.
type OrderConfirmationData struct {
	BaseEmailData
	OrderID           string `json:"orderID"`
	EstimatedDelivery string `json:"estimatedDelivery"`
	DashboardLink     string `json:"dashboardLink"`
}

// AccountActivationData holds data for an account activation email.
type AccountActivationData struct {
	BaseEmailData
	ActivationLink string `json:"activationLink"`
}
