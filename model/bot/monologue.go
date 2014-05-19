package bot

import "tweet-kcwidget/model/tweet"

func getMonologueTweetByName(name string) tweet.Tweet {
	return tweet.New(
		"",
		pickUpRandom(monologues[name]),
		tweet.TypeMonologue,
	)
}
func (b Bot) Monologue() (e error) {
	tw := getMonologueTweetByName(b.Name)
	return b.Tweet(tw)
}
