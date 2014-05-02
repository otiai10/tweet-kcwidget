package bot

import "github.com/mrjones/oauth"
import "tweet-kcwidget/conf/my"
import "tweet-kcwidget/model/tweet"
import "time"

type MyBot struct {
	consumer  *oauth.Consumer
	token     *oauth.AccessToken
	tweet     tweet.Tweet
	RestBegin int
	RestEnd   int
}

func New(
	consumer *oauth.Consumer,
	token *oauth.AccessToken,
	tweet tweet.Tweet,
	restBeginHour int,
	restEndHour int,
) *MyBot {
	return &MyBot{
		consumer:  consumer,
		token:     token,
		tweet:     tweet,
		RestBegin: restBeginHour,
		RestEnd:   restEndHour,
	}
}

func (b *MyBot) Tweet() (e error) {
	_, e = b.consumer.Post(
		"https://api.twitter.com/1.1/statuses/update.json",
		map[string]string{
			"status": b.tweet.ToText(),
		},
		b.token,
	)
	return e
}

func (b *MyBot) IsRest() bool {
	hour := time.Now().Hour()
	if b.RestBegin < hour && hour <= b.RestEnd {
		return true
	}
	return false
}

var provider = oauth.ServiceProvider{
	AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
	RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
	AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
}

var twitter_def = oauth.NewConsumer(
	my.AppConsumerKey,
	my.AppConsumerSecret,
	provider,
)
var token_def = &oauth.AccessToken{
	my.BotAccessToken,
	my.BotAccessTokenSecret,
	make(map[string]string),
}
var twitter00 = oauth.NewConsumer(
	my.App00ConsumerKey,
	my.App00ConsumerSecret,
	provider,
)
var token00 = &oauth.AccessToken{
	my.Bot00AccessToken,
	my.Bot00AccessTokenSecret,
	make(map[string]string),
}
var twitter01 = oauth.NewConsumer(
	my.App01ConsumerKey,
	my.App01ConsumerSecret,
	provider,
)
var token01 = &oauth.AccessToken{
	my.Bot01AccessToken,
	my.Bot01AccessTokenSecret,
	make(map[string]string),
}
var twitter02 = oauth.NewConsumer(
	my.App02ConsumerKey,
	my.App02ConsumerSecret,
	provider,
)
var token02 = &oauth.AccessToken{
	my.Bot02AccessToken,
	my.Bot02AccessTokenSecret,
	make(map[string]string),
}
var twitter03 = oauth.NewConsumer(
	my.App03ConsumerKey,
	my.App03ConsumerSecret,
	provider,
)
var token03 = &oauth.AccessToken{
	my.Bot03AccessToken,
	my.Bot03AccessTokenSecret,
	make(map[string]string),
}
