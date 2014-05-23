package bot

import "tweet-kcwidget/model/tweet"
import "math/rand"
import "time"

func GetMonologueTweetByName(name string) tweet.Tweet {
	return tweet.New(
		"",
		pickUpRandom(monologues[name]),
		tweet.TypeMonologue,
	)
}
func (b Bot) Monologue() (e error) {
	// とりあえずランダムで5分の1くらいにしぼる
	// TODO: ちゃんとする
	rand.Seed(time.Now().Unix())
	if rnd := rand.Intn(100); rnd < 90 {
		return
	}
	tw := GetMonologueTweetByName(b.Name)
	return b.Tweet(tw)
}
