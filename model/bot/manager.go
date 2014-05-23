package bot

import "tweet-kcwidget/conf/my"
import "tweet-kcwidget/model/tweet"
import "math/rand"
import "time"

func GetByName(name string) Bot {
	theToken := my.BotTokens[name]
	return Bot{name, theToken}
}
func getTimeKey(candidates map[int]string) int {
	hour := time.Now().Hour()
	key := hour - (hour % (24 / len(candidates)))
	return key
}
func GetAssigned(tw tweet.Tweet) Bot {
	candidates := getCandidates(tw)
	return GetByName(pickUpRandom(candidates))
}
func GetRandom() Bot {
	name := pickUpRandom(GetAllNames())
	return GetByName(name)
}
func getCandidates(tw tweet.Tweet) (names []string) {
	_ref := roster[tw.Kind]
	_except := getTimeKey(_ref)
	names = make([]string, len(_ref)-1)
	i := 0
	for key, name := range _ref {
		if key == _except {
			continue
		}
		names[i] = name
		i++
	}
	return
}
func GetAllNames() (names []string) {
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
