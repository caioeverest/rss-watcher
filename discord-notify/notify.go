package discord_notify

import (
	"net/http"
	"bytes"
	"log"
	"io/ioutil"

	"github.com/caioever/rss-watcher/config"
)

func Notify(conf *config.Config, payload []byte) {

	webhook := conf.BasicConfig.Webhook

	r := bytes.NewBuffer(payload)

	req, _ := http.NewRequest("POST", webhook, r)
	req.Header.Add("Content-Type", "application/json")
	res, _ := http.DefaultClient.Do(req)

	defer res.Body.Close()

	if res.StatusCode != 204 {
		body, _ := ioutil.ReadAll(res.Body)
		log.Printf("PAYLOAD SENDED: %s", payload)
		log.Printf("RESPONSE-BODY: %s", body)
		log.Panicf("ERROR | HTTP-CODE: %d", res.StatusCode)
	}
}
