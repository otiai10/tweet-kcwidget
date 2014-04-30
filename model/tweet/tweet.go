package tweet

import "fmt"
import "time"

type Tweet struct {
    ToScreenName string
    Message string
    Kind int
}
var tweet_format = "@%s %s %s"

func New(screen_name, message string, kind int) Tweet {
    return Tweet{
        ToScreenName: screen_name,
        Message: message,
        Kind: kind,
    }
}

func (t Tweet) ToText() string {
    debug_suffix := fmt.Sprintf("%v", time.Now().Unix())
    return fmt.Sprintf(tweet_format, t.ToScreenName, t.Message, debug_suffix)
}
