package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"

	"github.com/sirupsen/logrus"
	"gopkg.in/mailgun/mailgun-go.v1"
)

type TRNResponse struct {
	Error string  `json:"error"`
	Code  float64 `json:"code"`
}

func init() {
	logrus.SetFormatter(&logrus.TextFormatter{})
}

func main() {
	go heartBeat()

	var input string
	fmt.Scanln(&input)

}

func TRNAPI() (*TRNResponse, string) {
	url := "https://api.pubgtracker.com/v2/profile/pc/adamms"

	// Request
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		panic(err)
	}
	req.Header.Set("TRN-API-Key", "17d76a22-907b-4ab0-84fe-dfbc34cebd0c")

	// Response
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()

	body, _ := ioutil.ReadAll(resp.Body)

	trnResponse := &TRNResponse{}
	json.Unmarshal(body, trnResponse)

	return trnResponse, string(body)
}

func heartBeat() {
	for range time.Tick(time.Second * 10) {
		response, stringResponse := TRNAPI()
		if response.Error != "" && response.Code == 3 {
			logrus.Warn("code: ", response.Code)
			logrus.Warn("error: ", response.Error)
			fmt.Println("========================")
		} else {
			logrus.Info(stringResponse)
			fmt.Println("========================")

			mailMe(stringResponse)
		}
	}
}

func mailMe(stringResponse string) {
	mg := mailgun.NewMailgun("sandbox3fd66e607c004414a32485674ce9674f.mailgun.org", "key-81f9124c51b2a16192b5d546df0cae3b", "pubkey-27244492e5d0f52210d61f36ec18620c")
	message := mg.NewMessage(
		"pubgtrnbot@gmail.com",
		"PUBG Tracker Network Bot Status",
		stringResponse,
		"adamfdls@gmail.com")
	mg.Send(message)
}
