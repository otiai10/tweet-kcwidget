package queue

import "time"
import "tweet-kcwidget/model/tweet"
import "strconv"

type Queue struct {
	Time            int64 // Unix Time to fire
	DurationToSleep time.Duration
	Tweet           tweet.Tweet
}

var Ch = make(chan tweet.Tweet)

func New(finish_time int64, tweetVal tweet.Tweet) Queue {
	return Queue{
		Time:            finish_time,
		DurationToSleep: getDurationFromUnix(finish_time),
		Tweet:           tweetVal,
	}
}

func getDurationFromUnix(finish_time int64) (dur time.Duration) {

	now := time.Now().Unix()
	diff := int(finish_time - now - 60)
	if diff < 0 {
		return
	}
	dur, _ = time.ParseDuration(strconv.Itoa(diff) + "s")
	return
}

func Enq(q Queue) {
	go func() {
		time.Sleep(q.DurationToSleep)
		Ch <- q.Tweet
	}()
}
