package blocktag

import (
	"reflect"
	"strings"
	"testing"
)

func TestParse(t *testing.T) {
	tab := []struct {
		Text string
		Want []*Block
	}{
		{
			Text: "[info]test[/info]",
			Want: []*Block{
				&Block{
					Tag:  &Tag{Name: "info"},
					Body: []byte("test"),
				},
			},
		},
		{
			Text: "B[info]test[title]xx[/title][/info]E",
			Want: []*Block{
				&Block{
					Tag:  &Tag{Name: "info"},
					Body: []byte("test[title]xx[/title]"),
				},
			},
		},
		{
			Text: "[info]test[/info][code]aaa[/code]",
			Want: []*Block{
				&Block{
					Tag:  &Tag{Name: "info"},
					Body: []byte("test"),
				},
				&Block{
					Tag:  &Tag{Name: "code"},
					Body: []byte("aaa"),
				},
			},
		},
	}
	for _, v := range tab {
		r := strings.NewReader(v.Text)
		a, err := Parse(r)
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(a, v.Want) {
			t.Errorf("Parse(%q) = %v; want %v", v.Text, a, v.Want)
		}
	}
}
