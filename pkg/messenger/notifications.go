package messenger

import (
	"bytes"
	"fmt"
	"github.com/gin-gonic/gin/json"
	"io/ioutil"
	"log"
	"net/http"
)

func SendNotificationString(recipientId string, pageToken string, message string) error {
	return SendNotification(recipientId, pageToken, &FacebookMessage{
		Text: message,
	})
}

func SendNotification(recipientId string, pageToken string, message *FacebookMessage) error {
	url := fmt.Sprintf("https://graph.facebook.com/v2.6/me/messages?access_token=%s", pageToken)

	msg := FacebookResponse{
		MessagingType: "UPDATE",
		Recipient: struct {
			Id string `json:"id"`
		}{
			recipientId,
		},
		Message: message,
	}

	log.Printf("%#v", msg)

	msgBytes, err := json.Marshal(msg)
	if err != nil {
		return err
	}

	req, err := http.NewRequest("POST", url, bytes.NewBuffer(msgBytes))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	fmt.Println("response Status:", resp.Status)
	fmt.Println("response Headers:", resp.Header)
	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response Body:", string(body))

	return nil
}
