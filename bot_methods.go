package telebot

import(
	"time"
	"net/http"
	"io/ioutil"
	"errors"
	"strconv"
)

/******************************* Create Bot **********************************/

func CreateBot(token string, ownerID ...int) (bot *Bot, err error) {
	bot = new (Bot)
	bot.token = token
	bot.homeUrl = "https://api.telegram.org/bot"
	bot.fullUrl = bot.homeUrl + bot.token
	if len(ownerID) > 0 {
		bot.ownerID = ownerID[0]
		bot.AddNewAdmin(bot.ownerID)
	}
	bot.httpClient = http.Client{
		Timeout: time.Duration(10 * time.Second),
	}
	botAsUser, err := bot.GetMe()
	if err != nil {
		return bot, err
	}
	bot.firstName = botAsUser.FirstName
	bot.id = botAsUser.ID
	bot.username = botAsUser.Username
	bot.Activate()
	return bot, nil
}

/******************************* Get Methods *********************************/

func (bot *Bot) GetBotToken() (string) {
	return bot.token
}

func (bot *Bot) GetBotID() (int){
	return bot.id
}

func (bot *Bot) GetBotFirstName() (string){
	return bot.firstName
}

func (bot *Bot) GetBotUserName() (string){
	return bot.username
}

func (bot *Bot) GetBotOwnerID() (int){
	return bot.ownerID
}

func (bot *Bot) GetBotAdminIDs() []int {
	return bot.adminIDs
}

/******************************* Set Methods *********************************/

func (bot *Bot) Activate() {
	bot.activeStatus = true
}

func (bot *Bot) Deactivate() {
	bot.activeStatus = false
}

func (bot *Bot) SetOwnerID(ownerID int) {
	bot.ownerID = ownerID
}

func (bot *Bot) SetRequestTimeout(timeout time.Duration) {
	bot.httpClient.Timeout = timeout
}

func (bot *Bot) AddNewAdmin(newAdminID int) {
	for _, adminID := range bot.adminIDs {
        if adminID == newAdminID {
            return
        }
	}
	bot.adminIDs = append(bot.adminIDs, newAdminID)
}

func (bot *Bot) RemoveAdmin(deletedAdminID int) {
	for i, adminID := range bot.adminIDs {
        if adminID == deletedAdminID && deletedAdminID != bot.ownerID {
			bot.adminIDs[i] = bot.adminIDs[len(bot.adminIDs)-1]
			bot.adminIDs = bot.adminIDs[:len(bot.adminIDs)-1]
            return
        }
	}
}

/******************************* Additional Methods ******************************/

func MakeKeyboard(keyCaptions...string) {
	
}

/******************************* Private Methods ******************************/

func (bot *Bot) getRequestWithoutBody(url string) (err error) {
	resp, err := bot.httpClient.Get(url)
	if err != nil {
		err = errors.New("Request canceled while waiting for connection (" + bot.httpClient.Timeout.String() + " sec)")
		return err
	}
	if resp.StatusCode != 200 {
		err = errors.New("Bad response status code: " + strconv.Itoa(resp.StatusCode))
		return err
	}
	return nil
}

func (bot *Bot) getRequestWithBody(url string) (body []byte, err error) {
	resp, err := bot.httpClient.Get(url)
	if err != nil {
		err = errors.New("Request canceled while waiting for connection (" + bot.httpClient.Timeout.String() + " sec)")
		return nil, err
	}
	if resp.StatusCode != 200 {
		err = errors.New("Bad response status code: " + strconv.Itoa(resp.StatusCode))
		return nil, err
	}
	body, err = ioutil.ReadAll(resp.Body)
	if err != nil {
		err = errors.New("Read response body error")
		return nil, err
	}
	defer resp.Body.Close()
	return body, nil
}


