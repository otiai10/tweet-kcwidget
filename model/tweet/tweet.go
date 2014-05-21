package tweet

import "fmt"
import "time"

type Tweet struct {
	ToScreenName string
	Message      string
	Kind         int
}

var (
	TypeMonologue  = 0
	TypeMission    = 1
	TypeNyukyo     = 2
	TypeCreateship = 3
	TypeSortie     = 4
)
var (
	mention_format   = "@%s %s %s"
	monologue_format = "%s %s"
)

func New(screen_name, message string, kind int) Tweet {
	return Tweet{
		ToScreenName: screen_name,
		Message:      message,
		Kind:         kind,
	}
}

func (t Tweet) ToText() string {
	if t.Kind == TypeMonologue {
		return t.toMonologueText()
	}
	return t.toMentionText()
}
func (t Tweet) toMentionText() string {
	debug_suffix := time.Now().Format("15時04分05")
	return fmt.Sprintf(mention_format, t.ToScreenName, t.Message, debug_suffix)
}
func (t Tweet) toMonologueText() string {
	debug_suffix := time.Now().Format("15:04")
	return fmt.Sprintf(monologue_format, t.Message, debug_suffix)
}
