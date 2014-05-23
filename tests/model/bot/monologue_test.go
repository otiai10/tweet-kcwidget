package bot_test

import "testing"
import "tweet-kcwidget/model/bot"
import "fmt"

func TestBot_Monologue(t *testing.T) {
	for _, name := range bot.GetAllNames() {
		fmt.Println(bot.GetMonologueTweetByName(name))
	}
}
