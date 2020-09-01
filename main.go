package main

import (
   "os"
   "bytes"
	"encoding/json"
	"errors"
   "github.com/joho/godotenv"
	"fmt"
	"net/http"
)

type webhookReqBody struct {
	Message struct {
		Text string `json:"text"`
		Chat struct {
			ID int64 `json:"id"`
		} `json:"chat"`
	} `json:"message"`
}

func Handler(res http.ResponseWriter, req *http.Request) {
	body := &webhookReqBody{} 
   if err := json.NewDecoder(req.Body).Decode(body); err != nil {
		fmt.Println("could not decode request body", err)
		return
	}
   fmt.Println(body.Message.Text);
   lo := connect(body.Message.Text,PASSWORD);

	if err := response(body.Message.Chat.ID,lo); err != nil {
		fmt.Println("error in sending reply:", err)
		return
	}
	fmt.Println("reply sent")
}    

type sendMessageReqBody struct {
	ChatID int64  `json:"chat_id"`
	Text   string `json:"text"`
}

func response(chatID int64,arr string) error {
	reqBody := &sendMessageReqBody{
		ChatID: chatID,
		Text:   arr,
	}
	reqBytes, err := json.Marshal(reqBody)
	if err != nil {
		return err
	}
       
    HTTP :=  "https://api.telegram.org/bot"+TOKEN+"/sendMessage"
   fmt.Println(HTTP)
    res, err := http.Post(HTTP,"application/json", bytes.NewBuffer(reqBytes))
	if err != nil {
		return err
	}
	if res.StatusCode != http.StatusOK {
		return errors.New("unexpected status" + res.Status)
	}

	return nil
}
var PASSWORD, TOKEN string
func main()  {
    if err := godotenv.Load(); err != nil {
        fmt.Println("No .env file found")
    }
    PASSWORD= os.Getenv("PASSWORD")
    TOKEN = os.Getenv("TOKEN")

	http.ListenAndServe(":3000", http.HandlerFunc(Handler))
}
