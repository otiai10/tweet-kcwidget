package controllers

import "github.com/revel/revel"
import "tweet-kcwidget/conf/my"
import "tweet-kcwidget/model/queue"
import "tweet-kcwidget/model/tweet"
import "tweet-kcwidget/model/bot"

type Tweet struct {
	*revel.Controller
}

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
	q := queue.New(
		finish_time,
		tweet.New(screen_name, message, tweet.TypeMission),
	)
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
	q := queue.New(
		finish_time,
		tweet.New(screen_name, message, tweet.TypeNyukyo),
	)
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
	q := queue.New(
		finish_time,
		tweet.New(screen_name, message, tweet.TypeCreateship),
	)
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
	q := queue.New(
		finish_time,
		tweet.New(screen_name, message, tweet.TypeSortie),
	)
	queue.Enq(q)
	return c.RenderJson(&map[string]string{
		"result":  "ok",
		"message": message,
	})
}

func init() {
	go func() {
		for {
			queuedTweet := <-queue.Ch
			theBot := bot.GetAssigned(queuedTweet)
			if e := theBot.Tweet(queuedTweet); e != nil {
				revel.ERROR.Println(e)
			}
		}
	}()
}
