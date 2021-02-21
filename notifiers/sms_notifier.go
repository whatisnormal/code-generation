package notifiers

import (
	"encoding/json"
	"fmt"
	domain "github.com/whatisnormal/code-generation/domain"
	"io/ioutil"
	"log"
	"net/http"
	_ "net/http"
	"net/url"
	"strings"
)

type smsNotifier struct {
	originMsisdn string
	accountSid string
	authToken string
}

//NewFileRepo TODO
func NewSmsNotifier(originMsisdn string,
					accountSid  string,
					authToken  string) domain.Notifier {

	return &smsNotifier{
		originMsisdn,
		accountSid ,
		authToken,
	}
}

func (s smsNotifier) Notify(id string, code string) error {
	// Build message
	messageReader := buildMessage(s.originMsisdn, id, code)
	urlStr := "https://api.twilio.com/2010-04-01/Accounts/" + s.accountSid + "/Messages.json"

	// Create Request
	client := &http.Client{}

	req, _ := http.NewRequest("POST", urlStr, &messageReader)
	req.SetBasicAuth(s.accountSid, s.authToken)
	req.Header.Add("Accept", "application/json")
	req.Header.Add("Content-Type", "application/x-www-form-urlencoded")

	//Make request
	resp, _ := client.Do(req)

	log.Printf("Notified id: %v with code: %v - Status: %v", id, code, resp.Status)

	if resp.StatusCode >= 200 && resp.StatusCode < 300  {
		var data map[string]interface{}
		bodyBytes, _ := ioutil.ReadAll(resp.Body)
		err := json.Unmarshal(bodyBytes, &data)
		if err == nil  {
			log.Println(data["sid"])
		}else{
			return err
		}
	} else {
		log.Println(resp.Status)
	}

	return nil
}

func buildMessage(origin string, id string, code string) strings.Reader {
	v := url.Values{}

	v.Set("From",origin)
	v.Set("To",id)

	v.Set("Body",fmt.Sprintf("Code %v generated.", code))

	rb := *strings.NewReader(v.Encode())

	return rb
}
