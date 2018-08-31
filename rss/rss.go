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
		err := ioutil.WriteFile(fileName, []byte(time.Now().UTC().Format(time.RFC1123)), 0666)
		if err != nil {
			log.Fatal(err)
		}
	}

	memory := getFromMemory()

	if !haveNews(feed, feedSize, memory) {
		log.Printf("There's nothing new... :c")
		return
	}

	log.Printf("Wow looks like we have some new posts")
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

func haveNews(feed *gofeed.Feed, feedSize int, memory time.Time) bool{
	for i := feedSize-1; i >= 0; i-- {
		if  feed.Items[i].PublishedParsed.After(memory)  {
			return true
		}
	}

	return false
}

func send(conf *config.Config, feed *gofeed.Feed, feedSize int, memory time.Time) {
	temp := memory
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
			log.Printf("Sending payload")
			nofity.Notify(conf, payload)
			if feed.Items[i].PublishedParsed.After(temp) {
				temp = *feed.Items[i].PublishedParsed
				setNewHead(feed.Items[i].Published)
			}
		}
	}
}

func setNewHead(updated string) {
	log.Printf("Cleaning memory file")
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
