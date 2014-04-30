package controllers

// import "github.com/kurrik/twittergo"
import "github.com/mrjones/oauth"
import "github.com/revel/revel"
import "tweet-kcwidget/conf/my"
import "tweet-kcwidget/model/queue"
import "tweet-kcwidget/model/tweet"

import "fmt"

type Tweet struct {
	*revel.Controller
}

var provider = oauth.ServiceProvider{
	AuthorizeTokenUrl: "https://api.twitter.com/oauth/authorize",
	RequestTokenUrl:   "https://api.twitter.com/oauth/request_token",
	AccessTokenUrl:    "https://api.twitter.com/oauth/access_token",
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

var (
	type_mission    = 0
	type_nyukyo     = 1
	type_createship = 2
	type_sortie     = 3
)

var reciever = make(chan tweet.Tweet)

func (c Tweet) validate(screen_name, message, client_token string) (ok bool, errMess string) {
	ok = true
	if client_token != my.ServerToken {
		ok = false
		errMess = "Invalid token"
		return
	}
	if screen_name == "" || message == "" {
		ok = false
		errMess = "Missing parameter"
		return
	}
	return
}

func (c Tweet) Mission(finish_time int64, screen_name, message, client_token string) revel.Result {
	if ok, errorMess := c.validate(screen_name, message, client_token); !ok {
		return c.RenderJson(&map[string]string{
			"result":  "ng",
			"message": errorMess,
		})
	}
	q := queue.New(finish_time, tweet.New(screen_name, message, type_mission))
	queue.Enq(q)
	return c.RenderJson(&map[string]string{
		"result":  "ok",
		"message": message,
	})
}
func (c Tweet) Nyukyo(finish_time int64, screen_name, message, client_token string) revel.Result {
	if ok, errorMess := c.validate(screen_name, message, client_token); !ok {
		return c.RenderJson(&map[string]string{
			"result":  "ng",
			"message": errorMess,
		})
	}
	q := queue.New(finish_time, tweet.New(screen_name, message, type_nyukyo))
	queue.Enq(q)
	return c.RenderJson(&map[string]string{
		"result":  "ok",
		"message": message,
	})
}
func (c Tweet) Createship(finish_time int64, screen_name, message, client_token string) revel.Result {
	if ok, errorMess := c.validate(screen_name, message, client_token); !ok {
		return c.RenderJson(&map[string]string{
			"result":  "ng",
			"message": errorMess,
		})
	}
	q := queue.New(finish_time, tweet.New(screen_name, message, type_createship))
	queue.Enq(q)
	return c.RenderJson(&map[string]string{
		"result":  "ok",
		"message": message,
	})
}
func (c Tweet) Sortie(finish_time int64, screen_name, message, client_token string) revel.Result {
	if ok, errorMess := c.validate(screen_name, message, client_token); !ok {
		return c.RenderJson(&map[string]string{
			"result":  "ng",
			"message": errorMess,
		})
	}
	q := queue.New(finish_time, tweet.New(screen_name, message, type_sortie))
	queue.Enq(q)
	return c.RenderJson(&map[string]string{
		"result":  "ok",
		"message": message,
	})
}

func getBot(tweetToTweet tweet.Tweet) (consumer *oauth.Consumer, token *oauth.AccessToken) {
	if tweetToTweet.Kind < type_createship {
		consumer = twitter01
		token = token01
	} else {
		consumer = twitter00
		token = token00
	}
	return
}

func init() {
	go func() {
		for {
			tweetToTweet := <-queue.Ch
			consumer, token := getBot(tweetToTweet)
			_, e := consumer.Post(
				"https://api.twitter.com/1.1/statuses/update.json",
				map[string]string{
					"status": tweetToTweet.ToText(),
				},
				token,
			)
			if e != nil {
				fmt.Printf("%+v", e)
			}
		}
	}()
}
