package bot

import "tweet-kcwidget/model/tweet"

func Manage(tw tweet.Tweet) *MyBot {
	bot := get(tw)
	if bot.IsRest() {
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
	return New(
		twitter_def,
		token_def,
		tw,
		0,
		0,
	)
}
func get00(tw tweet.Tweet) *MyBot {
	return New(
		twitter00,
		token00,
		tw,
		0,
		6,
	)
}
func get01(tw tweet.Tweet) *MyBot {
	return New(
		twitter01,
		token01,
		tw,
		6,
		12,
	)
}

// おなくなりになった
func get02(tw tweet.Tweet) *MyBot {
	return New(
		twitter02,
		token02,
		tw,
		12,
		18,
	)
}
func get03(tw tweet.Tweet) *MyBot {
	return New(
		twitter03,
		token03,
		tw,
		6,
		12,
	)
}
