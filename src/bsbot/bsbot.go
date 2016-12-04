// Space Rocks Bot is a bot watching
// asteroids coming too close to earth for the incoming days/week.
package main

// TODO: change twitter banner periodically

import (
	"flag"
	"fmt"
	"log"
	"path/filepath"
	"time"

	"github.com/dns-gh/betterlog"
	conf "github.com/dns-gh/flagsconfig"
	"github.com/dns-gh/twbot"
)

// Twitter constants
const (
	projectURL               = "https://github.com/dns-gh/bsbot"
	updateFlag               = "update"
	twitterFollowersPathFlag = "twitter-followers-path"
	twitterFriendsPathFlag   = "twitter-friends-path"
	twitterTweetsPathFlag    = "twitter-tweets-path"
	bsPathFlag               = "bs-path"
	debugFlag                = "debug"
)

func main() {
	update := flag.Duration(updateFlag, 24*time.Hour, "[twitter] update frequency of the bot for tweets")
	twitterFollowersPath := flag.String(twitterFollowersPathFlag, "followers.json", "[twitter] data file path for followers")
	twitterFriendsPath := flag.String(twitterFriendsPathFlag, "friends.json", "[twitter] data file path for friends")
	twitterTweetsPath := flag.String(twitterTweetsPathFlag, "tweets.json", "[twitter] data file path for tweets")
	twitterConsumerKey := flag.String("TWITTER_CONSUMER_KEY", "", "[twitter] consumer key")
	twitterConsumerSecret := flag.String("TWITTER_CONSUMER_SECRET", "", "[twitter] consumer secret")
	twitterAccessToken := flag.String("TWITTER_ACCESS_TOKEN", "", "[twitter] access token")
	twitterAccessSecret := flag.String("TWITTER_ACCESS_SECRET", "", "[twitter] access secret")
	bsPath := flag.String(bsPathFlag, "betaseries.json", "[betaseries] data file path for betaseries")
	bsKey := flag.String("BS_API_KEY", "", "[betaseries] api key")
	debug := flag.Bool(debugFlag, false, "[twitter] debug mode")
	_, err := conf.NewConfig("robot.config")
	f, err := betterlog.MakeDateLogger(filepath.Join("Debug", "bot.log"))
	if err != nil {
		log.Fatalln(err.Error())
	}
	defer f.Close()
	log.Println("[twitter] update:", *update)
	log.Println("[twitter] twitter-followers-path:", *twitterFollowersPath)
	log.Println("[twitter] twitter-friends-path:", *twitterFriendsPath)
	log.Println("[twitter] twitter-tweets-path:", *twitterTweetsPath)
	log.Println("[twitter] TWITTER_CONSUMER_KEY:", *twitterConsumerKey)
	log.Println("[twitter] TWITTER_CONSUMER_SECRET:", *twitterConsumerSecret)
	log.Println("[twitter] TWITTER_ACCESS_TOKEN:", *twitterAccessToken)
	log.Println("[twitter] TWITTER_ACCESS_SECRET:", *twitterAccessSecret)
	log.Println("[betaseries] bs-path:", *bsPath)
	log.Println("[betaseries] BS_API_KEY:", *bsKey)
	log.Println("[twitter] debug:", *debug)
	bot := twbot.MakeTwitterBotWithCredentials(*twitterFollowersPath, *twitterFriendsPath, *twitterTweetsPath,
		*twitterConsumerKey,
		*twitterConsumerSecret,
		*twitterAccessToken,
		*twitterAccessSecret,
		*debug)
	defer bot.Close()
	bsBot := makeBetaseriesBot(*bsKey, *bsPath, bot)
	bsBot.TweetNewsAsync(30 * time.Minute)
	bsBot.UpdateProfileBannerAsync(30 * time.Second)
	bot.TweetPeriodicallyAsync(func() (string, error) {
		return fmt.Sprintf("Hey, I'm a bot tweeting about #tvshows and #movie updates, check out my #source #code %s and help me #improve !", projectURL), nil
	}, *update)
	bot.Wait()
}
