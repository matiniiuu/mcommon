package enums

type PaytabsTransactionClass string

const (
	Ecom PaytabsTransactionClass = "ecom" //! ECommerce Online Transaction

	Moto PaytabsTransactionClass = "moto" //! Mail Order Telephonic Order Transaction

	Cont PaytabsTransactionClass = "cont" //! Recurring token-based transaction

)
