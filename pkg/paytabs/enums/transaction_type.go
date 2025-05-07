package enums

type PaytabsTransactionType string

const (
	Sale     PaytabsTransactionType = "sale"
	Auth     PaytabsTransactionType = "auth"
	Capture  PaytabsTransactionType = "capture"
	Void     PaytabsTransactionType = "void"
	Register PaytabsTransactionType = "register"
	Refund   PaytabsTransactionType = "refund"
)
