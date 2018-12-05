package telebot

type	Game			struct {
	Title				string
	Description			string
	Photo				[]PhotoSize
	Text				string
	Text_Entities		[]MessageEntity
	Animation			Animation
}

type	GameHighScore	struct {
	Position			int
	User				User
	Score				int
}

type	CallbackGame	struct {

}
