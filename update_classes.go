package telebot

type	Update				struct {
	Update_id				int
	Message					Message
	Edited_message			Message
	Channel_post			Message
	Edited_channel_post		Message
	Inline_query			InlineQuery
	Chosen_inline_result	ChosenInlineResult
	Callback_query			CallbackQuery
	Shipping_query			ShippingQuery
	Pre_checkout_query		PreCheckoutQuery
}

type	WebhookInfo			struct {
	Url						string
	Has_custom_certificate	bool
	Pending_update_count	int
	Last_error_date			int
	Last_error_message		string
	Max_connections			int
	Allowed_updates			[]string
}
