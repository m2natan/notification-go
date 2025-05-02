package sms

import (
	"encoding/json"
	"fmt"
	"os"

	"github.com/twilio/twilio-go"
	twilioApi "github.com/twilio/twilio-go/rest/api/v2010"
)

type TwilioNotifier struct {
	client *twilio.RestClient
	number string
}

func NewTwilioNotifier() *TwilioNotifier {
	accountSid := os.Getenv("AACOUNT_SID")
	authToken := os.Getenv("AUTH_TOKEN")
	number := os.Getenv("TWILIO_NUMBER")

	client := twilio.NewRestClientWithParams(twilio.ClientParams{
		Username: accountSid,
		Password: authToken,
	})
	return &TwilioNotifier{client: client, number: number}
}
func (t *TwilioNotifier) SendSms(recipientNumber, content string) error {

	if t.client == nil {
		return fmt.Errorf("Twilio client not initialized, call sms.Init() first")
	}

	params := &twilioApi.CreateMessageParams{}
	params.SetTo(recipientNumber)
	params.SetFrom(t.number)
	params.SetBody(content)

	resp, err := t.client.Api.CreateMessage(params)
	if err != nil {
		return fmt.Errorf("%s", "Error sending SMS message: " + err.Error())
	} else {
		response, _ := json.Marshal(*resp)
		fmt.Println("Response: " + string(response))
	}

	return nil
}
