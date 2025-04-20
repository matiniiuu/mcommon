package enums

type OtpFormat string

const (
	Numeric      OtpFormat = "numeric"
	Alphanumeric OtpFormat = "alphanumeric"
	Alphabetic   OtpFormat = "alphabetic"
)
