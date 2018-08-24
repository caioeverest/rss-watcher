package main

import (
	"log"
	"os"
	"os/signal"

	"github.com/robfig/cron"
	"github.com/caioever/rss-watcher/config"
	"github.com/caioever/rss-watcher/rss"
)

var conf = config.CreateConfig()

func main() {
	log.Printf("Starting bot...")
	log.Printf("Variables set with:\n| %s \n| %s \n| %s \n| %s \n| %s \n| %s", conf.BasicConfig.Webhook, conf.BasicConfig.Cron, conf.BasicConfig.RssFeed, conf.PayloadConfig.AvatarUrl, conf.PayloadConfig.Color, conf.PayloadConfig.Username)

	scheduler := cron.New()
	defer scheduler.Stop()
	log.Printf("Opening cron")
	scheduler.AddFunc(conf.BasicConfig.Cron, job)
	go scheduler.Start()

	sig := make(chan os.Signal)
	signal.Notify(sig, os.Interrupt, os.Kill)
	<-sig
}

func job() {
	log.Printf("Starting Job...")
	rss.Checkfeed(conf)
}
