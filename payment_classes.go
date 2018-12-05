package telebot

type	LabeledPrice							struct {
	Label										string
	Amount										int
}

type	Invoice									struct {
	Title										string
	Description									string
	Start_parameter								string
	Currency									string
	Total_amount								int
}

type	ShippingAddress							struct {
	Country_code								string
	State										string
	City										string
	Street_line1								string
	Street_line2								string
	Post_code									string
}

type	OrderInfo								struct {
	Name										string
	Phone_number								string
	Email										string
	Shipping_address							ShippingAddress
}

type	ShippingOption							struct {
	Id											string
	Title										string
	Prices										[]LabeledPrice
}

type	SuccessfulPayment						struct {
	Currency									string
	Total_amount								int
	Invoice_payload								string
	Shipping_option_id							string
	Order_info									OrderInfo
	Telegram_payment_charge_id					string
	Provider_payment_charge_id					string
}

type	ShippingQuery							struct {
	Id											string
	From										User
	Invoice_payload								string
	Shipping_address							ShippingAddress
}

type	PreCheckoutQuery						struct {
	Id											string
	From										User
	Currency									string
	Total_amount								int
	Invoice_payload								string
	Shipping_option_id							string
	Order_info									OrderInfo
}

type	PassportData 							struct {
	Data										[]EncryptedPassportElement
	Credentials									EncryptedCredentials
}

type	PassportFile							struct {
	File_id										string
	File_size									int
	File_date									int
}

type	EncryptedPassportElement				struct {
	Type										string
	Data										string
	Phone_number								string
	Email										string
	Files										[]PassportFile
	Front_side									PassportFile
	Reverse_side								PassportFile
	Selfie										PassportFile
	Translation									[]PassportFile
	Hash										string
}

type	EncryptedCredentials					struct {
	Data										string
	Hash										string
	Secret										string
}

type	PassportElementErrorDataField			struct {
	Source										string
	Type										string
	Field_name									string
	Data_hash									string
	Message										string
}

type	PassportElementErrorFrontSide			struct {
	Source										string
	Type										string
	File_hash									string
	Message										string
}

type	PassportElementErrorReverseSide			struct {
	Source										string
	Type										string
	File_hash									string
	Message										string
}

type	PassportElementErrorSelfie				struct {
	Source										string
	Type										string
	File_hash									string
	Message										string
}

type	PassportElementErrorFile				struct {
	Source										string
	Type										string
	File_hash									string
	Message										string
}

type	PassportElementErrorFiles				struct {
	Source										string
	Type										string
	File_hashes									[]string
	Message										string
}

type	PassportElementErrorTranslationFile		struct {
	Source										string
	Type										string
	File_hash									string
	Message										string
}

type	PassportElementErrorTranslationFiles	struct {
	Source										string
	Type										string
	File_hashes									[]string
	Message										string
}

type PassportElementErrorUnspecified			struct {
	Source										string
	Type										string
	Element_hash								string
	Message										string
}
