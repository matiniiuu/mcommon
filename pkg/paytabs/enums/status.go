package enums

type PaymentStatus string

const (
	PaymentStatusAuthorized PaymentStatus = "Authorized"
	PaymentStatusHold       PaymentStatus = "Hold"
	PaymentStatusPending    PaymentStatus = "Pending"
	PaymentStatusVoided     PaymentStatus = "Voided"
	PaymentStatusError      PaymentStatus = "Error"
	PaymentStatusDeclined   PaymentStatus = "Declined"
	PaymentStatusExpired    PaymentStatus = "Expired"
)
