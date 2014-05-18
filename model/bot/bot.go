package bot

import "github.com/mrjones/oauth"
import "tweet-kcwidget/conf/my"
import "tweet-kcwidget/model/tweet"

// import "time"

var provider = oauth.ServiceProvider{
	AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
	RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
	AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
}
var consumer = oauth.NewConsumer(
	my.ConsumerKey,
	my.ConsumerSecret,
	provider,
)

type Bot struct {
	Name  string
	token my.AccessToken
}

func (b Bot) Tweet(tw tweet.Tweet) (e error) {
	token := &oauth.AccessToken{
		b.token.Token,
		b.token.Secret,
		make(map[string]string),
	}
	_, e = consumer.Post(
		"https://api.twitter.com/1.1/statuses/update.json",
		map[string]string{
			"status": tw.ToText(),
		},
		token,
	)
	return
}
