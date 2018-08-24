package config

import (
	"os"
	"strconv"
	"log"
)

var fundamentalKeys = []string{"WEBHOOK", "RSS", "CHARACTER_LIMIT", "AVATAR_URL"}
var importantKeys = []string{"CRON", "CHARACTER_LIMIT", "PAYLOAD_COLOR", "BOT_NAME"}

type Config struct {
	BasicConfig 	BasicConfig 	`json:"basic_config"`
	PayloadConfig 	PayloadConfig 	`json:"payload_config"`
}

type BasicConfig struct {
	Webhook 		string			`json:"webhook"`
	RssFeed 		string 			`json:"rss_feed"`
	Cron 			string 			`json:"cron"`
	CharacterLimit 	int 			`json:"character_limit"`
}

type PayloadConfig struct {
	Color 			string 			`json:"color"`
	AvatarUrl 		string 			`json:"avatar_url"`
	Username		string 			`json:"username"`
}

func CreateConfig() *Config {
	config := &Config{}

	config.BasicConfig.CharacterLimit = variableCheck()
	config.BasicConfig.Webhook = os.Getenv("WEBHOOK")
	config.BasicConfig.RssFeed = os.Getenv("RSS")
	config.BasicConfig.Cron = os.Getenv("CRON")

	config.PayloadConfig.Color = os.Getenv("PAYLOAD_COLOR")
	config.PayloadConfig.AvatarUrl = os.Getenv("AVATAR_URL")
	config.PayloadConfig.Username = os.Getenv("BOT_NAME")

	return config
}

func variableCheck() (int){

	for _, key := range fundamentalKeys {
		if _, ok := os.LookupEnv(key); !ok {
			log.Fatalf("Variable \"%s\" missing! the bot will stop.", key)
			os.Exit(0)
		}
	}

	charcterSize, err := strconv.Atoi(os.Getenv("CHARACTER_LIMIT"))
	if err != nil {
		log.Fatal("Variable CHARACTER_LIMIT is not a number")
		os.Exit(0)
	} else if 0 > charcterSize || charcterSize > 2048 {
		log.Panic("Variable CHARACTER_LIMIT wrongly configured, set for default")
		return 2048
	}
	return charcterSize
}