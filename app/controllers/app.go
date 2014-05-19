package controllers

import "github.com/revel/revel"
import "tweet-kcwidget/conf/my"

type App struct {
	*revel.Controller
}

func (c App) Index() revel.Result {
	bots := my.BotTokens
	return c.Render(bots)
}
