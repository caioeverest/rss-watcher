package rss

import (
	"fmt"
	"log"
	"os"
	"io/ioutil"
	"time"

	"github.com/mmcdole/gofeed"
	"github.com/caioever/rss-watcher/config"
	"github.com/caioever/rss-watcher/util"
	nofity "github.com/caioever/rss-watcher/discord-notify"
)

var fileName = "tmp.dat"

func Checkfeed(conf *config.Config) {
	log.Printf("Reading feed")
	fp := gofeed.NewParser()
	feed, _ := fp.ParseURL(conf.BasicConfig.RssFeed)
	feedSize := len(feed.Items)
	log.Printf("There are %d itens on the RSS", feedSize)

	if !fileExist() {
		log.Printf("Creating memory file")
		err := ioutil.WriteFile(fileName, []byte(feed.Published), 0666)
		if err != nil {
			log.Fatal(err)
		}
	}

	memory := getFromMemory()

	if !haveNews(feed.PublishedParsed, memory) {
		log.Printf("There's nothing new... :c")
		return
	}

	log.Printf("Wow looks like we have some new posts")
	setNewHead(feed.Published)
	send(conf, feed, feedSize, memory)
}

func fileExist() bool{
	_, err := os.Stat(fileName)
	if err != nil {
		if os.IsNotExist(err) {
			return false
		}
	}
	return true
}

func getFromMemory() time.Time{
	dat, err := ioutil.ReadFile(fileName)
	if err != nil {
		fmt.Print(err)
	}

	data := string(dat)
	final,_ := util.ParseDate(data)

	return final.UTC()
}

func haveNews(published *time.Time, memory time.Time) bool{
	if memory.Equal(*published) {
		return false
	}
	return true
}

func send(conf *config.Config, feed *gofeed.Feed, feedSize int, memory time.Time) {
	for i := feedSize-1; i >= 0; i-- {
		if  feed.Items[i].PublishedParsed.After(memory)  {
			log.Printf("Building body of the post from date: %s", feed.Items[i].Published)
			body := util.BodyVar{
				feed.Items[i].Title,
				feed.Items[i].Link,
				feed.Items[i].Author.Name,
				util.Format(conf, feed.Items[i].Description),
				fmt.Sprintln(feed.Items[i].PublishedParsed.Format(time.RFC3339)),
			}

			payload := util.BuildPayload(conf, body)
			nofity.Notify(conf, payload)
		}
	}
}

func setNewHead(updated string) {
	err := os.Truncate(fileName, 100)
	if err != nil {
		log.Fatal(err)
	}

	log.Printf("Writing new Tag on memory file")
	err = ioutil.WriteFile(fileName, []byte(updated), 0666)
	if err != nil {
		log.Fatal(err)
	}
}
