package main

import (
	"bytes"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"time"

	"sync"

	bs "github.com/dns-gh/bs-client/bsclient"
	"github.com/dns-gh/freeze"
	"github.com/dns-gh/tojson"
	"github.com/dns-gh/twbot"
)

type betaseriesBot struct {
	bsClient *bs.BetaSeries
	twbot    *twbot.TwitterBot
	path     string
	News     map[string]bs.News `json:"news"`
	mutex    sync.Mutex
}

func makeBetaseriesBot(key, path string, twbot *twbot.TwitterBot) *betaseriesBot {
	bsClient, err := bs.NewBetaseriesClient(key, "", "")
	if err != nil {
		log.Fatalln(err.Error())
	}
	b := &betaseriesBot{
		bsClient: bsClient,
		twbot:    twbot,
		path:     path,
	}
	b.load()
	if b.News == nil {
		b.News = make(map[string]bs.News)
	}
	return b
}

func (b *betaseriesBot) load() {
	if _, err := os.Stat(b.path); os.IsNotExist(err) {
		tojson.Save(b.path, b)
	}
	err := tojson.Load(b.path, b)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func (b *betaseriesBot) save() {
	err := tojson.Save(b.path, b)
	if err != nil {
		log.Fatalln(err.Error())
	}
}

func loadImage(uri string) (string, error) {
	resp, err := http.Get(uri)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	if resp.StatusCode != 200 {
		return "", fmt.Errorf("error request status: %s != 200", resp.Status)
	}
	buf := new(bytes.Buffer)
	if _, err := io.Copy(buf, resp.Body); err != nil {
		return "", err
	}
	return buf.String(), nil
}

func (b *betaseriesBot) tweetNews(item *bs.News) {
	b.mutex.Lock()
	defer b.mutex.Unlock()
	if _, ok := b.News[item.ID]; !ok {
		img, err := loadImage(item.PictureURL)
		if err != nil {
			log.Println(err.Error())
			return
		}
		err = b.twbot.TweetImageOnce(item.Title, item.URL, img)
		if err != nil {
			log.Println(err.Error())
			return
		}
		log.Printf("tweeting news (ID: %s): %s\n", item.ID, item.Title)
		b.News[item.ID] = *item
		b.save()
		freeze.SleepMinMax(10, 20)
	}
}

func (b *betaseriesBot) TweetNewsAsync(freq time.Duration) {
	go func() {
		ticker := time.NewTicker(freq)
		defer ticker.Stop()
		for _ = range ticker.C {
			news, err := b.bsClient.NewsLast(20, false)
			if err != nil {
				log.Println(err.Error())
				continue
			}
			// TODO read the news backward or chronologically
			for _, v := range news {
				b.tweetNews(&v)
			}
		}
	}()
}
