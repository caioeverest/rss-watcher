package util

import (
	"encoding/json"
	"fmt"

	"github.com/caioever/rss-watcher/config"
)

type Body struct {
	Username 	string		`json:"username"`
	AvatarUrl 	string		`json:"avatar_url"`
	Embeds 		[]Embeds	`json:"embeds"`
}

type Embeds struct {
	Title 		string 		`json:"title"`
	Url 		string  	`json:"url"`
	Description string 		`json:"description"`
	Color 		string 		`json:"color"`
	Timestamp	string		`json:"timestamp"`
	Footer		Footer		`json:"footer"`
	Fields		[]Fields	`json:"fields,omitempty"`
}

type Fields struct {
	Name		string		`json:"name,omitempty"`
	Value		string		`json:"value,omitempty"`
}

type Footer struct {
	Text 		string	`json:"text"`
} 

type BodyVar struct {
	Title 			string
	Url 			string
	Author			string
	Description 	string
	Timestamp		string
}

func BuildPayload(conf *config.Config, body BodyVar) []byte{
	var b Body
	var footer Footer

	b.Username = conf.PayloadConfig.Username
	b.AvatarUrl = conf.PayloadConfig.AvatarUrl

	footer.Text = fmt.Sprintf("Posted by %s", body.Author)

	b.Embeds = append(b.Embeds, Embeds{body.Title,
		body.Url,
		body.Description + fmt.Sprintf("\n\n[[READ MORE]](%s)", body.Url),
		conf.PayloadConfig.Color,
		body.Timestamp,
		footer,
		nil,
	})

	payload, err := json.Marshal(b)
	if err != nil {
		fmt.Println(err)
	}

	return payload
}