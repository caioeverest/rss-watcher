package util

import (
	"bytes"
	"log"

	"github.com/lunny/html2md"
	"github.com/caioever/rss-watcher/config"
)

func Format(conf *config.Config, html string) string{
	html2md.AddRule("title", &html2md.Rule{
		Patterns: []string{"h([1-6])"},
		//Tp: html2md.Void,
		Replacement: func(innerHTML string, attrs []string) string {
			log.Print(len(attrs))
			if len(attrs) < 4 || attrs[0] != attrs[len(attrs)-1] {
				return ""
			}

			return "**" + attrs[2] + "**"
		},
	})

	text := html2md.Convert(html)
	final := bytes.Runes([]byte(text))
	if len(final) > conf.BasicConfig.CharacterLimit {
		log.Printf("To much content!")
		return string(final[:conf.BasicConfig.CharacterLimit])
	}
	return text
}
