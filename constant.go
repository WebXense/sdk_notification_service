package sdk_notification_service

const (
	url_twilio_sms  = "/twilio/sms"
	url_gmail_email = "/gmail/email"
)

const (
	EMAIL_BODY_TYPE_HTML = "text/html"
	EMAIL_BODY_TYPE_TEXT = "text/plain"
)

const (
	ERR_TWILIO_ENV_NOT_SET     = "322a4687-f578-4633-b7cc-aab0045604cf"
	ERR_TWILIO_SEND_SMS_FAILED = "1c05741b-4dd3-4a65-8d66-823f5c89440f"

	ERR_GMAIL_ENV_NOT_SET       = "dbdec7f3-22ff-47a9-b860-935f4c99530a"
	ERR_GMAIL_SEND_EMAIL_FAILED = "11144453-c98f-42a3-b177-d7f00cfbafd7"
	ERR_GMAIL_INVALID_BODY_TYPE = "b7b094a5-21ab-4d07-b2b9-79a0c304ba8a"
)
