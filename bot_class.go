package telebot

import "net/http"

type	Bot 	struct {
	token			string
	homeUrl			string
	fullUrl			string
	id				int
	firstName		string
	username		string
	webhook			WebhookInfo
	httpClient		http.Client
	ownerID			int
	adminIDs		[]int
	activeStatus	bool
}
