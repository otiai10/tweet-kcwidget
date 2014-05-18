package bot

import "tweet-kcwidget/model/tweet"

var monologues = map[string][]string{
	"Yudachi_kcw": {
		"ぽい？",
		"ぽい！",
		"ぽい〜♪",
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Kaga_kcw": {
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Akagi_kcw": {
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"401_kcw": {
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Shimakaze_kcw": {
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Hibiki_kcw": {
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Mogami_kcw": {
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Sakawa_kcw": {
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Aoba_kcw": {
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
	},
	"Kinugasa_kcw": {
		"（アイコン描いてくれるひと & ひとりごとセリフ募集中です...！）",
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
