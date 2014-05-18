package bot

import "tweet-kcwidget/model/tweet"

var monologues = map[string][]string{
	"otiai10": {
		"これはotiai10のひとりごと000",
		"これはotiai10のひとりごと001",
		"これはotiai10のひとりごと002",
	},
	"Yudachi_kcw": {
		"ぽい？",
		"ぽい！",
		"ぽい〜♪",
	},
}

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
