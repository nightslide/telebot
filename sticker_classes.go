package telebot

type	Sticker			struct {
	File_id				string
	Width				int
	Height				int
	Thumb				PhotoSize
	Emoji				string
	Set_name			string
	Mask_position		MaskPosition
	File_size			int
}

type	StickerSet		struct {
	Name				string
	Title				string
	Contains_masks		bool
	Stickers			[]Sticker
}

type	MaskPosition	struct {
	Point				string
	X_shift				float32
	Y_shift				float32
	Scale				float32
}
