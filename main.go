package main

import (
	"bytes"
	"fmt"
	"github.com/joho/godotenv"
	"html/template"
	"io/ioutil"
	"log"
	"os"
	"strings"

	"github.com/devndops/notify/mail"
	"github.com/devndops/notify/models"
	"github.com/mailersend/mailersend-go"
)

func init() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Println("Error loading .env file")
	}
	log.Println(".env file loaded successfully")
}

// RenderTemplate loads an HTML template from the given file path and renders it with data.
func RenderTemplate(templatePath string, data interface{}) (string, error) {
	content, err := ioutil.ReadFile(templatePath)
	if err != nil {
		return "", fmt.Errorf("error reading template file %s: %w", templatePath, err)
	}
	tmpl, err := template.New("email").Parse(string(content))
	if err != nil {
		return "", fmt.Errorf("error parsing template: %w", err)
	}
	var buf bytes.Buffer
	if err := tmpl.Execute(&buf, data); err != nil {
		return "", fmt.Errorf("error executing template: %w", err)
	}
	return buf.String(), nil
}

// SendOrderConfirmationEmail renders the order confirmation template and sends the email.
func SendOrderConfirmationEmail(apiKey string, data models.OrderConfirmationData, templatePath string) (string, error) {
	// Fill in common fields.
	models.FillBaseData(&data.BaseEmailData)
	// Render the HTML content.
	htmlContent, err := RenderTemplate(templatePath, data)
	if err != nil {
		return "", err
	}
	subject := "Order Confirmation - " + data.OrderID
	from := mailersend.From{
		Name:  data.AppName,
		Email: fmt.Sprintf("no-reply@%s.com", data.AppName),
	}
	recipients := []mailersend.Recipient{
		{
			Name:  data.RecipientName,
			Email: data.RecipientEmail,
		},
	}
	return mail.SendHTMLEmail(apiKey, subject, htmlContent, from, recipients, "order-confirmation")
}

// SendAccountActivationEmail renders the account activation template and sends the email.
func SendAccountActivationEmail(apiKey string, data models.AccountActivationData, templatePath string) (string, error) {
	models.FillBaseData(&data.BaseEmailData)
	htmlContent, err := RenderTemplate(templatePath, data)
	if err != nil {
		return "", err
	}
	subject := "Activate Your " + data.AppName + " Account"
	from := mailersend.From{
		Name:  data.AppName,
		Email: fmt.Sprintf(strings.ToLower("no-reply@%s.com"), data.AppName),
	}
	recipients := []mailersend.Recipient{
		{
			Name:  data.RecipientName,
			Email: data.RecipientEmail,
		},
	}
	return mail.SendHTMLEmail(apiKey, subject, htmlContent, from, recipients, "account-activation")
}

func main() {
	// Retrieve the MailerSend API key from the environment.
	apiKey := os.Getenv("MAILERSEND_API_KEY")
	if apiKey == "" {
		fmt.Println("MAILERSEND_API_KEY not set")
		return
	}

	// -------------------------------
	// Example: Order Confirmation Email
	// -------------------------------
	orderData := models.OrderConfirmationData{
		BaseEmailData: models.BaseEmailData{
			AppName:        "SlangsWiki",
			AppLink:        "https://myapp.com",
			Logo:           "https://myapp.com/logo.png",
			RecipientEmail: "olagoke.olasebikan@gmail.com",
			RecipientName:  "Olagoke",
			Theme:          models.ModernTheme(), // Using preset ModernTheme.
		},
		OrderID:           "12345",
		EstimatedDelivery: "2025-03-15",
		DashboardLink:     "https://myapp.com/dashboard",
	}
	orderTemplatePath := "templates/order_confirmation.html"
	orderMsgID, err := SendOrderConfirmationEmail(apiKey, orderData, orderTemplatePath)
	if err != nil {
		fmt.Println("Error sending order confirmation email:", err)
		return
	}
	fmt.Println("Order Confirmation Email sent with Message ID:", orderMsgID)

	// -------------------------------
	// Example: Account Activation Email
	// -------------------------------
	activationData := models.AccountActivationData{
		BaseEmailData: models.BaseEmailData{
			AppName:        "SlangsWiki",
			AppLink:        "https://myapp.com",
			Logo:           "https://myapp.com/logo.png",
			RecipientEmail: "olagoke.olasebikan@gmail.com",
			RecipientName:  "Olagoke",
			Theme:          models.DefaultTheme(), // Using default theme.
		},
		ActivationLink: "https://myapp.com/activate?token=abcdef",
	}
	activationTemplatePath := "templates/account_activation.html"
	actMsgID, err := SendAccountActivationEmail(apiKey, activationData, activationTemplatePath)
	if err != nil {
		fmt.Println("Error sending account activation email:", err)
		return
	}
	fmt.Println("Account Activation Email sent with Message ID:", actMsgID)
}
