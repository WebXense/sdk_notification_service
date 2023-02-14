package sdk_notification_service

type TwilioSendSMSResponseDTO struct {
	From  string `json:"from"`
	To    string `json:"to"`
	Body  string `json:"body"`
	Price string `json:"price"`
	Unit  string `json:"unit"`
}

type GmailSendEmailResponseDTO struct {
	From     string   `json:"from"`
	To       []string `json:"to"`
	Cc       []string `json:"cc"`
	Subject  string   `json:"subject"`
	Body     string   `json:"body"`
	BodyType string   `json:"bodyType"`
}

type twilioSendSMSRequest struct {
	To   string `json:"to" binding:"required"`
	Body string `json:"body" binding:"required"`
}

type gmailSendEmailRequest struct {
	To       []string `json:"to" binding:"required"`
	Cc       []string `json:"cc" binding:"required"`
	Subject  string   `json:"subject" binding:"required"`
	Body     string   `json:"body" binding:"required"`
	BodyType string   `json:"bodyType" binding:"required"`
}
