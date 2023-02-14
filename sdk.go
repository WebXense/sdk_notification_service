package sdk_notification_service

import (
	"encoding/json"
	"errors"
	"strconv"

	"github.com/WebXense/ginger/ginger"
	"github.com/WebXense/http"
)

type notificationServiceSdk struct {
	host string
}

func NewNotificationServiceSdk(host string) *notificationServiceSdk {
	return &notificationServiceSdk{
		host: host,
	}
}

func (s *notificationServiceSdk) SendTwilioSms(to, body string) (*TwilioSendSMSResponseDTO, *ginger.Response, error) {
	status, resp, err := http.Post[ginger.Response](s.host+url_twilio_sms, nil, nil, &twilioSendSMSRequest{
		To:   to,
		Body: body,
	})
	if err != nil {
		return nil, nil, err
	}
	if status != 200 {
		return nil, nil, errors.New("send twilio sms failed with status: " + strconv.Itoa(status))
	}
	if !resp.Success {
		return nil, resp, nil
	}
	return mapByJson(resp.Data, &TwilioSendSMSResponseDTO{}), resp, nil
}

func (s *notificationServiceSdk) SendGmailEmail(to, cc []string, subject, body, bodyType string) (*GmailSendEmailResponseDTO, *ginger.Response, error) {
	status, resp, err := http.Post[ginger.Response](s.host+url_gmail_email, nil, nil, &gmailSendEmailRequest{
		To:       to,
		Cc:       cc,
		Subject:  subject,
		Body:     body,
		BodyType: bodyType,
	})
	if err != nil {
		return nil, nil, err
	}
	if status != 200 {
		return nil, nil, errors.New("send gmail email failed with status: " + strconv.Itoa(status))
	}
	if !resp.Success {
		return nil, resp, nil
	}
	return mapByJson(resp.Data, &GmailSendEmailResponseDTO{}), resp, nil
}

func mapByJson[T any](from interface{}, to *T) *T {
	j, err := json.Marshal(from)
	if err != nil {
		return nil
	}
	err = json.Unmarshal(j, to)
	if err != nil {
		return nil
	}
	return to
}
