package bot

import "tweet-kcwidget/conf/my"
import "tweet-kcwidget/model/tweet"
import "math/rand"
import "time"

var roster = map[int]map[int]string{
	tweet.TypeMission: {
		0:  "Yudachi_kcw",
		4:  "Yudachi_kcw",
		8:  "Yudachi_kcw",
		12: "Yudachi_kcw",
		16: "Yudachi_kcw",
		20: "Yudachi_kcw",
	},
	tweet.TypeNyukyo: {
		0:  "Yudachi_kcw",
		4:  "Yudachi_kcw",
		8:  "Yudachi_kcw",
		12: "Yudachi_kcw",
		16: "Yudachi_kcw",
		20: "Yudachi_kcw",
	},
	tweet.TypeCreateship: {
		0:  "Yudachi_kcw",
		4:  "Yudachi_kcw",
		8:  "Yudachi_kcw",
		12: "Yudachi_kcw",
		16: "Yudachi_kcw",
		20: "Yudachi_kcw",
	},
	tweet.TypeSortie: {
		0:  "Yudachi_kcw",
		4:  "Yudachi_kcw",
		8:  "Yudachi_kcw",
		12: "Yudachi_kcw",
		16: "Yudachi_kcw",
		20: "Yudachi_kcw",
	},
}

func GetByName(name string) Bot {
	theToken := my.BotTokens[name]
	return Bot{name, theToken}
}
func GetAssigned(tw tweet.Tweet) Bot {
	candidates := roster[tw.Kind]
	hour := time.Now().Hour()
	key := hour - (hour % (24 / len(candidates)))
	return GetByName(candidates[key])
}
func GetRandom() Bot {
	name := pickUpRandom(getAllNames())
	return GetByName(name)
}
func getAllNames() (names []string) {
	names = make([]string, len(my.BotTokens))
	i := 0
	for name, _ := range my.BotTokens {
		names[i] = name
		i++
	}
	return
}
func pickUpRandom(names []string) (name string) {
	rand.Seed(time.Now().UnixNano())
	name = names[rand.Intn(len(names))]
	return
}
