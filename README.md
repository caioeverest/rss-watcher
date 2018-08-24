
# RSS Watcher  
This is a simple bot that posts the content from an RSS feed to discord via *webhook*.

## Dependencies:
 - github.com/robfig/cron  
 - github.com/lunny/html2md  
 - github.com/mmcdole/gofeed

## Running
Here's a list of the easiest ways to run the bot, but you can use whatever you feel the most comfortable with.

### docker-compose:
The docker image is already available on the my docker hub repository (*caiobarcelos/rss-watcher*). In order to run the bot, you can copy the docker-compose.yml below and change it to fit your needs.
```yaml
version: '2'
services:
  rss-watcher:
    image: caiobarcelos/rss-watcher:latest
    container_name: rss-watcher
    environment:
      - WEBHOOK=http://something.from/discod
      - RSS=http://rss.your-feed.of/preference
      - CRON=*/10 * * * * *
      - CHARACTER_LIMIT=1900
      - PAYLOAD_COLOR=16685830
      - AVATAR_URL=http://your-avatar.jpg/
      - BOT_NAME=YourBotName
```
And that's it! Just execute ~~order 66~~ in your terminal `docker-compose up -d` and it will pull the image and build the container.

###### \*the image will compilate the bot when the container runs, it already goes with the dependencies.

### go run:
If you want run the bot, I recommend you use the following script to export the environment variables
, and then run.
```
#!/usr/bin/env bash
export WEBHOOK="http://something.from/discod"
export RSS="http://rss.your-feed.of/preference"
export CRON="*/3 * * * * *"
export CHARACTER_LIMIT="1900"
export PAYLOAD_COLOR="16685830"
export AVATAR_URL="https://your-avatar.jpg/"
export BOT_NAME="YourBotName"

go run main.go
```
#### See? Easy peasy lemon squeezy.
###### \*In this case you need to download all the dependencies

## Environment variables explained:
#### \*To run the bot you first need to set ALL environment variables!


**WEBHOOK** = `http://something.from/discod`
> Discord's webhook.

**RSS** = `http://rss.your-feed.of/preference`
> Your RSS link.

**CRON** = `*/3 * * * * *` OR `@hourly` OR `@every 1h30m`
> Cron to run the pooling at the RSS.

**CHARACTER_LIMIT** = 1900
> Max amount of characters to be sent in embedded description (Discord's max is 2048) but there is also a [Read more] with an URL for the news. We recommend setting it at 1900 characters.

**PAYLOAD_COLOR** = 16685830
> The color of the embedded box. Discord understands this color through a decimal value, so in order to set a color based on the HexTable, you must convert the identificator to a decimal value.
> Here is a web [conversor](https://www.binaryhexconverter.com/hex-to-decimal-converter)

**AVATAR_URL** = `https://static.kingdom.gg/discord-news-bot-logo.png`
> The bot's avatar.

**BOT_NAME** = Watch~~Man~~Bot
> The bot's name.

# MIT License
Copyright (c) 2018 Caio Everest

Permission is hereby granted, free of charge, to any person obtaining a copy
of this software and associated documentation files (the "Software"), to deal
in the Software without restriction, including without limitation the rights
to use, copy, modify, merge, publish, distribute, sublicense, and/or sell
copies of the Software, and to permit persons to whom the Software is
furnished to do so, subject to the following conditions:

The above copyright notice and this permission notice shall be included in all
copies or substantial portions of the Software.

THE SOFTWARE IS PROVIDED "AS IS", WITHOUT WARRANTY OF ANY KIND, EXPRESS OR
IMPLIED, INCLUDING BUT NOT LIMITED TO THE WARRANTIES OF MERCHANTABILITY,
FITNESS FOR A PARTICULAR PURPOSE AND NONINFRINGEMENT. IN NO EVENT SHALL THE
AUTHORS OR COPYRIGHT HOLDERS BE LIABLE FOR ANY CLAIM, DAMAGES OR OTHER
LIABILITY, WHETHER IN AN ACTION OF CONTRACT, TORT OR OTHERWISE, ARISING FROM,
OUT OF OR IN CONNECTION WITH THE SOFTWARE OR THE USE OR OTHER DEALINGS IN THE
SOFTWARE.
