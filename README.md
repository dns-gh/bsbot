# bsbot

[![Go Report Card](https://goreportcard.com/badge/github.com/dns-gh/bsbot)](https://goreportcard.com/report/github.com/dns-gh/bsbot)

bstbot is a Twitter bot tweeting about new TV Shows and movie broadcasts

It uses the following betaseries API: https://www.betaseries.com/api/

It has also a specified and user-defined bot behaviors thanks to https://github.com/dns-gh/twbot.

## Motivation

Simply for fun, practice and trying to do something useful at the same time :)

Feel free to join my efforts!

## Installation

- It requires Go language of course. You can set it up by downloading it here: https://golang.org/dl/
- Install it here C:/Go.
- Set your GOPATH, GOROOT and PATH environment variables with:

```
export GOROOT=C:/Go
export GOPATH=WORKING_DIR
export PATH=C:/Go/bin:${PATH}
```

or:

```
@working_dir $ source build/go.sh
```

and then set up your API keys/tokens/secrets:

```
export TWITTER_CONSUMER_KEY="your_twitter_consumer_key"
export TWITTER_CONSUMER_SECRET="your_twitter_consumer_secret"
export TWITTER_ACCESS_TOKEN="your_twitter_access_token"
export TWITTER_ACCESS_SECRET="your_twitter_access_secret"
export BS_API_KEY="your_betaseries_api_key"
```

You can find get them here: https://www.betaseries.com/api/ and https://apps.twitter.com/

## Build and usage

```
@working_dir $ go install bsbot
@working_dir $ bin/bsbot.exe -help
  -BS_API_KEY string
        [betaseries] api key
  -TWITTER_ACCESS_SECRET string
        [twitter] access secret
  -TWITTER_ACCESS_TOKEN string
        [twitter] access token
  -TWITTER_CONSUMER_KEY string
        [twitter] consumer key
  -TWITTER_CONSUMER_SECRET string
        [twitter] consumer secret
  -bs-path string
        [betaseries] data file path for betaseries (default "betaseries.json")
  -config string
        configuration filename (default "robot.config")
  -debug
        [twitter] debug mode
  -twitter-followers-path string
        [twitter] data file path for followers (default "followers.json")
  -twitter-friends-path string
        [twitter] data file path for friends (default "friends.json")
  -twitter-tweets-path string
        [twitter] data file path for tweets (default "tweets.json")
  -update duration
        [twitter] update frequency of the bot for tweets (default 24h0m0s)
@working_dir $ bin/bsbot.exe -debug=false
[2016-12-02 17:52:34] [info] logging to: D:\WORK\bsbot\bin\Debug\bot.log
[2016-12-02 17:52:34] [twitter] update: 24h0m0s
[2016-12-02 17:52:34] [twitter] twitter-followers-path: followers.json
[2016-12-02 17:52:34] [twitter] twitter-friends-path: friends.json
[2016-12-02 17:52:34] [twitter] twitter-tweets-path: tweets.json
[2016-12-02 17:52:34] [twitter] TWITTER_CONSUMER_KEY: XXXXXXXXXX
[2016-12-02 17:52:34] [twitter] TWITTER_CONSUMER_SECRET: XXXXXXXXXX
[2016-12-02 17:52:34] [twitter] TWITTER_ACCESS_TOKEN: XXXXXXXXXX
[2016-12-02 17:52:34] [twitter] TWITTER_ACCESS_SECRET: XXXXXXXXXX
[2016-12-02 17:52:34] [betaseries] bs-path: betaseries.json
[2016-12-02 17:52:34] [betaseries] BS_API_KEY: XXXXXXXXXX
[2016-12-02 17:52:34] [twitter] debug: false
```

## License

See the included LICENSE file.