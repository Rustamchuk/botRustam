package commands

import (
	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
	"math/rand"
)

func (c *Commander) Default(inputMessage *tgbotapi.Message) {
	ans := ""

	res, ok := c.tongue.Answers()[inputMessage.Text]
	if ok {
		ans += res.Phrase
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, ans)
		c.bot.Send(msg)
		return
	}

	randomNumber := rand.Intn(2)
	if randomNumber == 1 {
		msg := tgbotapi.NewMessage(inputMessage.Chat.ID, "Здесь была гифка но меня забанили. Вот тебе шутка\n\n"+
			"https://randstuff.ru/joke/")
		c.bot.Send(msg)
		//rand.Seed(time.Now().UnixNano())
		//
		//// Get a random gif URL using a service like Giphy
		//gifURL, err := getRandomGifURL()
		//if err != nil {
		//	log.Fatal(err)
		//}
		//
		//// Download the gif file
		//resp, err := http.Get(gifURL)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//defer resp.Body.Close()
		//
		//// Read the gif file into a byte array
		//gifBytes, err := ioutil.ReadAll(resp.Body)
		//if err != nil {
		//	log.Fatal(err)
		//}
		//
		//// Create a FileBytes object with the gif byte array
		//fileBytes := tgbotapi.FileBytes{
		//	Name:  "random.gif",
		//	Bytes: gifBytes,
		//}
		//
		//// Create an animation message with the FileBytes object
		//msg := tgbotapi.NewAnimation(inputMessage.Chat.ID, fileBytes)
		//
		//// Send the message with the animation
		//_, err = c.bot.Send(msg)
		//if err != nil {
		//	log.Fatal(err)
		//}
		return
	}
	stories := c.tongue.Stories()
	randomNumber = rand.Intn(len(stories))

	ans += stories[randomNumber].Phrase
	msg := tgbotapi.NewMessage(inputMessage.Chat.ID, ans)
	c.bot.Send(msg)
}

//func getRandomGifURL() (string, error) {
//	// Make a request to a service like Giphy to get a random gif
//	resp, err := http.Get("https://api.giphy.com/v1/gifs/random?api_key=D8Nkm9RcTKDVm7do8nww1ENxtHf4KaCA")
//	if err != nil {
//		return "", err
//	}
//	defer resp.Body.Close()
//
//	// Parse the response and get the URL of the random gif
//	var data struct {
//		Data struct {
//			ImageURL string `json:"image_url"`
//		} `json:"data"`
//	}
//	err = json.NewDecoder(resp.Body).Decode(&data)
//	if err != nil {
//		return "", err
//	}
//
//	// Check if the response contains any GIFs
//	if data.Data.ImageURL == "" {
//		return "", fmt.Errorf("no GIFs found")
//	}
//
//	// Get the URL of the random gif
//	gifURL := data.Data.ImageURL
//
//	// Return the URL of the random gif
//	return gifURL, nil
//}
