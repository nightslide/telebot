package telebot

type	User 		struct {
	ID				int			`json:"id"`
	IsBot			bool		`json:"is_bot"`
	FirstName		string		`json:"first_name"`
	LastName		string		`json:"is_bot"`
	Username		string		`json:"username"`
	LanguageCode	string		`json:"language_code"`
}

type	Chat						struct {
	ID								int				`json:"id"`
	Type							string			`json:"type"`
	Title							string			`json:"title"`
	Username						string			`json:"username"`
	FirstName						string			`json:"first_name"`
	LastName						string			`json:"last_name"`
	AllMembersAreAdministrators		bool			`json:"all_members_are_administrators"`
	Photo							ChatPhoto		`json:"photo"`
	Description						string			`json:"description"`
	InviteLink						string			`json:"invite_link"`
	PinnedMessage					Message			`json:"pinned_message"`
	StickerSetName					string			`json:"sticker_set_name"`
	CanSetStickerSet				bool			`json:"can_set_sticker_set"`
}

type	Message							struct {
	Message_id							int
	From								User
	Date								int
	Chat								*Chat
	Forward_from						User
	Forward_from_chat					*Chat
	Forward_from_message_id				int
	Forward_signature					string
	Forward_date						int
	Reply_to_message					*Message
	Edit_date							int
	Media_group_id						string
	Author_signature					string
	Text								string
	Entities							[]MessageEntity
	Caption_entities					[]MessageEntity
	Audio								Audio
	Document							Document
	Animation							Animation
	Game								Game
	Photo								[]PhotoSize
	Sticker								Sticker
	Video								Video
	Voice								Voice
	Video_Note							VideoNote
	Caption								string
	Contact								Contact
	Location							Location
	Venue								Venue
	New_chat_members					[]User
	Left_chat_member					User
	New_chat_title						string
	New_chat_photo						[]PhotoSize
	Delete_chat_photo					bool
	Group_chat_created					bool
	Supergroup_chat_created				bool
	Channel_chat_created				bool
	Migrate_to_chat_id					int
	Migrate_from_chat_id				int
	Pinned_message						*Message
	Invoice								Invoice
	Successful_payment					SuccessfulPayment
	Connected_website					string
	Passport_data						PassportData
}

type	MessageEntity					struct {
	Type								string
	Offset								int
	Length								int
	Url									string
	User								User
}

type	PhotoSize						struct {
	File_id								string
	Width								int
	Height								int
	File_size							int
}

type	Audio							struct {
	File_id								string
	Duration							int
	Performer							string
	Title								string
	Mime_type							string
	File_size							int
	Thumb								PhotoSize
}

type	Document						struct {
	File_id								string
	Thumb								PhotoSize
	File_name							string
	Mime_type							string
	File_size							int
}

type	Video							struct {
	File_id								string
	Width								int
	Height								int
	Duration							int
	Thumb								PhotoSize
	Mime_type							string
	File_size							int
}

type	Animation						struct {
	File_id								string
	Width								int
	Height								int
	Duration							int
	Thumb								PhotoSize
	File_name							string
	Mime_type							string
	File_size							int
}

type	Voice							struct {
	File_id								string
	Duration							int
	Mime_type							string
	File_size							int
}

type	VideoNote						struct {
	File_id								string
	Length								int
	Duration							int
	Thumb								PhotoSize
	File_size							int
}

type	Contact							struct {
	Phone_number						string
	First_name							string
	Last_name							string
	User_id								int
	Vcard								string
}

type	Location						struct {
	Longitude							float32
	Latitude							float32
}

type	Venue							struct {
	Location							Location
	Title								string
	Address								string
	Foursquare_id						string
	Foursquare_type						string
}

type	UserProfilePhotos				struct {
	Total_count							int
	Photos								[][]PhotoSize
}

type	File							struct {
	File_id								string
	File_size							int
	File_path							string
}

type	ReplyKeyboardMarkup				struct {
	Keyboard							[][]KeyboardButton
	Resize_keyboard						bool
	One_time_keyboard					bool
	Selective							bool
}

type	KeyboardButton					struct {
	Text								string
	Request_contact						bool
	Request_location					bool
}

type	ReplyKeyboardRemove				struct {
	Remove_keyboard						bool
	Selective							bool
}

type	InlineKeyboardMarkup			struct {
	Inline_keyboard						[][]InlineKeyboardButton
}

type	InlineKeyboardButton			struct {
	Text								string
	Url									string
	Callback_data						string
	Switch_inline_query					string
	Switch_inline_query_current_chat	string
	Callback_game						CallbackGame
	Pay									bool
}

type	CallbackQuery					struct {
	Id									string
	From								User
	Message								Message
	Inline_message_id					string
	Chat_instance						string
	Data								string
	Game_short_name						string
}

type	ForceReply						struct {
	ForceReply							bool
	Selective							bool
}

type	ChatPhoto						struct {
	Small_file_id						string
	Big_file_id							string
}

type	ChatMember						struct {
	User								User
	Status								string
	Until_date							int
	Can_be_edited						bool
	Can_change_info						bool
	Can_post_messages					bool
	Can_edit_messages					bool
	Can_delete_messages					bool
	Can_invite_users					bool
	Can_restrict_members				bool
	Can_pin_messages					bool
	Can_promote_members					bool
	Can_send_messages					bool
	Can_send_media_messages				bool
	Can_send_other_messages				bool
	Can_add_web_page_previews			bool
}

type	ResponseParameters				struct {
	Migrate_to_chat_id					int
	Retry_after							int
}

type	InputMediaPhoto					struct {
	Type								string
	Media								string
	Caption								string
	Parse_mode							string
}

type	InputMediaVideo					struct {
	Type								string
	Media								string
	Thumb								interface{}
	Caption								string
	Parse_mode							string
	Width								int
	Height								int
	Duration							int
	Supports_streaming					bool
}

type	InputMediaAnimation				struct {
	Type								string
	Media								string
	Thumb								interface{}
	Caption								string
	Parse_mode							string
	Width								int
	Height								int
	Duration							int
}

type	InputMediaAudio					struct {
	Type								string
	Media								string
	Thumb								interface{}
	Caption								string
	Parse_mode							string
	Duration							int
	Performer							string
	Title								string
}

type	InputMediaDocument				struct {
	Type								string
	Media								string
	Thumb								interface{}
	Caption								string
	Parse_mode							string
}
