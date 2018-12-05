package telebot

type	InlineQuery						struct {
	Id									string
	From								User
	Location							Location
	Query								string
	Offset								string
}

type	InlineQueryResultArticle		struct {
	Type								string
	Id									string
	Title								string
	Input_message_content				InputMessageContent
	Reply_markup						InlineKeyboardMarkup
	Url									string
	Hide_url							bool
	Description							string
	Thumb_url							string
	Thumb_width							int
	Thumb_height						int
}

type	InlineQueryResultPhoto			struct {
	Type								string
	Id									string
	Photo_url							string
	Thumb_url							string
	Photo_width							int
	Photo_height						int
	Title								string
	Description							string
	Caption								string
	Parse_mode							string
	Reply_markup						InlineKeyboardMarkup
	Input_message_content				InputMessageContent
}

type	InlineQueryResultGif			struct {
	Type								string
	Id									string
	Gif_url								string
	Gif_width							int
	Gif_height							int
	Gif_duration						int
	Thumb_url							string
	Title								string
	Caption								string
	Parse_mode							string
	Reply_markup						InlineKeyboardMarkup
	Input_message_content				InputMessageContent
}

type	InlineQueryResultMpeg4Gif		struct {
	Type								string
	Id									string
	Mpeg4_url							string
	Mpeg4_width							int
	Mpeg4_height						int
	Mpeg4_duration						int
	Thumb_url							string
	Title								string
	Caption								string
	Parse_mode							string
	Reply_markup						InlineKeyboardMarkup
	Input_message_content				InputMessageContent
}

type	InlineQueryResultVideo			struct {
	Type								string
	Id									string
	Video_url							string
	Mime_type							string
	Thumb_url							string
	Title								string
	Caption								string
	Parse_mode							string
	Video_width							int
	Video_height						int
	Video_duration						int
	Description							string
	Reply_markup						InlineKeyboardMarkup
	Input_message_content				InputMessageContent
}

type	InlineQueryResultAudio			struct {
	Type								string
	Id									string
	Audio_url							string
	Title								string
	Caption								string
	Parse_mode							string
	Performer							string
	Audio_duration						int
	Reply_markup						InlineKeyboardMarkup
	Input_message_content				InputMessageContent
}

type	InlineQueryResultVoice			struct {
	Type								string
	Id									string
	Voice_url							string
	Title								string
	Caption								string
	Parse_mode							string
	Voice_duration						int
	Reply_markup						InlineKeyboardMarkup
	Input_message_content				InputMessageContent
}

type	InlineQueryResultDocument		struct {
	Type								string
	Id									string
	Title								string
	Caption								string
	Parse_mode							string
	Document_url						string
	Mime_type							string
	Description							string
	Reply_markup						InlineKeyboardMarkup
	Input_message_content				InputMessageContent
	Thumb_url							string
	Thumb_width							int
	Thumb_height						int
}

type	InlineQueryResultLocation		struct {
	Type								string
	Id									string
	Latitude							float32
	Longitude							float32
	Title								string
	Live_period							int
	Reply_markup						InlineKeyboardMarkup
	Input_message_content				InputMessageContent
	Thumb_url							string
	Thumb_width							int
	Thumb_height						int
}

type	InlineQueryResultVenue			struct {
	Type								string
	Id									string
	Latitude							float32
	Longitude							float32
	Title								string
	Address								string
	Foursquare_id						string
	Foursquare_type						string
	Reply_markup						InlineKeyboardMarkup
	Input_message_content				InputMessageContent
	Thumb_url							string
	Thumb_width							int
	Thumb_height						int
}

type	InlineQueryResultContact		struct {
	Type								string
	Id									string
	Phone_number						string
	First_name							string
	Last_name							string
	Vcard								string
	Reply_markup						InlineKeyboardMarkup
	Input_message_content				InputMessageContent
	Thumb_url							string
	Thumb_width							int
	Thumb_height						int
}

type	InlineQueryResultGame			struct {
	Type								string
	Id									string
	Game_short_name						string
	Reply_markup						InlineKeyboardMarkup
}

type	InlineQueryResultCachedPhoto	struct {
	Type								string
	Id									string
	Photo_file_id						string
	Title								string
	Description							string
	Caption								string
	Parse_mode							string
	Reply_markup						InlineKeyboardMarkup
	Input_message_content				InputMessageContent
}

type	InlineQueryResultCachedGif		struct {
	Type								string
	Id									string
	Gif_file_id							string
	Title								string
	Caption								string
	Parse_mode							string
	Reply_markup						InlineKeyboardMarkup
	Input_message_content				InputMessageContent
}

type	InlineQueryResultCachedMpeg4Gif	struct {
	Type								string
	Id									string
	Mpeg4_file_id						string
	Title								string
	Caption								string
	Parse_mode							string
	Reply_markup						InlineKeyboardMarkup
	Input_message_content				InputMessageContent
}

type	InlineQueryResultCachedSticker	struct {
	Type								string
	Id									string
	Sticker_file_id						string
	Reply_markup						InlineKeyboardMarkup
	Input_message_content				InputMessageContent
}

type	InlineQueryResultCachedDocument	struct {
	Type								string
	Id									string
	Title								string
	Document_file_id					string
	Description							string
	Caption								string
	Parse_mode							string
	Reply_markup						InlineKeyboardMarkup
	Input_message_content				InputMessageContent
}

type	InlineQueryResultCachedVideo	struct {
	Type								string
	Id									string
	Video_file_id						string
	Title								string
	Description							string
	Caption								string
	Parse_mode							string
	Reply_markup						InlineKeyboardMarkup
	Input_message_content				InputMessageContent
}

type	InlineQueryResultCachedVoice	struct {
	Type								string
	Id									string
	Voice_file_id						string
	Title								string
	Description							string
	Caption								string
	Parse_mode							string
	Reply_markup						InlineKeyboardMarkup
	Input_message_content				InputMessageContent
}

type	InlineQueryResultCachedAudio	struct {
	Type								string
	Id									string
	Audio_file_id						string
	Caption								string
	Parse_mode							string
	Reply_markup						InlineKeyboardMarkup
	Input_message_content				InputMessageContent
}

type	InputMessageContent				interface {

}

type	InputTextMessageContent			struct {
	Message_text						string
	Parse_mode							string
	Disable_web_page_preview			bool
}

type	InputLocationMessageContent		struct {
	Latitude							float32
	Longitude							float32
	Live_period							int
}

type	InputVenueMessageContent		struct {
	Latitude							float32
	Longitude							float32
	Title								string
	Address								string
	Foursquare_id						string
	Foursquare_type						string
}

type	InputContactMessageContent		struct {
	Phone_number						string
	First_name							string
	Last_name							string
	Vcard								string
}

type	ChosenInlineResult				struct {
	Result_id							string
	From								User
	Location							Location
	Inline_message_id					string
	Query								string
}
