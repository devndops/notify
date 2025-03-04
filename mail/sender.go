// Package mail helps send HTML email via the MailerSend API
package mail

import (
	"context"
	"fmt"
	"time"

	"github.com/mailersend/mailersend-go"
)

// SendHTMLEmail sends an HTML email via MailerSend.
// The API key is passed to this function so that a new MailerSend client is created for each call.
func SendHTMLEmail(apiKey, subject, htmlContent string, from mailersend.From, recipients []mailersend.Recipient, tag string) (string, error) {
	client := mailersend.NewMailersend(apiKey)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	message := client.Email.NewMessage()
	message.SetFrom(from)
	message.SetRecipients(recipients)
	message.SetSubject(subject)
	message.SetHTML(htmlContent)
	message.SetTags([]string{tag})

	res, err := client.Email.Send(ctx, message)
	if err != nil {
		return "", fmt.Errorf("error sending email: %w", err)
	}
	return res.Header.Get("X-Message-Id"), nil
}
