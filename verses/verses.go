package verses

import "fmt"

type Text struct {
	Text string `json:"text"`
}
type Caps struct {
	Texts []Text `json:"verses"`
}
type DailyMessage struct {
	Response Caps `json:"response"`
}

func (v Caps) String() string {
	result := ""

	for _, l := range v.Texts {
		result += fmt.Sprintf("%s\n", l)
	}

	return result
}

func (r DailyMessage) String() string {
	return fmt.Sprintf("%s\n", r.Response)
}
