package bot

import "github.com/mrjones/oauth"
import "tweet-kcwidget/conf/my"
import "tweet-kcwidget/model/tweet"

import "fmt"

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

func Get(tweetToTweet tweet.Tweet) (consumer *oauth.Consumer, token *oauth.AccessToken) {
    switch tweetToTweet.Kind {
    case tweet.TypeMission:
        return getDefault()
    case tweet.TypeNyukyo:
        return get00()
    case tweet.TypeCreateship:
        return get01()
    case tweet.TypeSortie:
        return get02()
    }
	return getDefault()
}

func getDefault() (consumer *oauth.Consumer, token *oauth.AccessToken) {
    fmt.Println("use デフォルトちゃん")
    consumer = twitter_def
    token = token_def
    return
}
func get00() (consumer *oauth.Consumer, token *oauth.AccessToken) {
    fmt.Println("use 0号ちゃん")
    consumer = twitter00
    token = token00
    return
}
func get01() (consumer *oauth.Consumer, token *oauth.AccessToken) {
    fmt.Println("use 1号ちゃん")
    consumer = twitter01
    token = token01
    return
}
func get02() (consumer *oauth.Consumer, token *oauth.AccessToken) {
    fmt.Println("use 2号ちゃん")
    consumer = twitter02
    token = token02
    return
}
