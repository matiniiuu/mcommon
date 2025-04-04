package email

type (
	Email interface {
		SendHtml(from string, to []string, subject, messageBody string) error
		SendText(from string, to []string, subject, messageBody string) error
	}
)
