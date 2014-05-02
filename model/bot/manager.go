package bot

import "tweet-kcwidget/model/tweet"
import "fmt"

func Manage(tw tweet.Tweet) *MyBot {
	bot := get(tw)
	if bot.IsRest() {
		fmt.Println("は、おやすみです")
		return getDefault(tw)
	}
	return bot
}
func get(tw tweet.Tweet) *MyBot {
	switch tw.Kind {
	case tweet.TypeMission:
		return get00(tw)
	case tweet.TypeNyukyo:
		return get02(tw)
	case tweet.TypeCreateship:
		// return get01(tw)
		return get03(tw)
	case tweet.TypeSortie:
		return get03(tw)
	}
	return getDefault(tw)
}

func getDefault(tw tweet.Tweet) *MyBot {
	fmt.Println("デフォルトちゃん")
	return New(
		twitter_def,
		token_def,
		tw,
		0,
		0,
	)
}
func get00(tw tweet.Tweet) *MyBot {
	fmt.Println("0号ちゃん")
	return New(
		twitter00,
		token00,
		tw,
		0,
		8,
	)
}
func get01(tw tweet.Tweet) *MyBot {
	fmt.Println("1号ちゃん")
	return New(
		twitter01,
		token01,
		tw,
		8,
		16,
	)
}
func get02(tw tweet.Tweet) *MyBot {
	fmt.Println("2号ちゃん")
	return New(
		twitter02,
		token02,
		tw,
		16,
		24,
	)
}
func get03(tw tweet.Tweet) *MyBot {
	fmt.Println("3号ちゃん")
	return New(
		twitter03,
		token03,
		tw,
		8,
		16,
	)
}
