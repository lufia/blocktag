package blocktag

import (
	"reflect"
	"testing"
)

func TestParseTag(t *testing.T) {
	tab := []struct {
		tag  string
		want *Tag
	}{
		{
			tag: "tag",
			want: &Tag{
				Name:  "tag",
				Attrs: map[string]string{},
			},
		},
		{
			tag: "name uid=1",
			want: &Tag{
				Name:  "name",
				Attrs: map[string]string{"uid": "1"},
			},
		},
		{
			tag: "name uid=1 to=aaa flag",
			want: &Tag{
				Name:  "name",
				Attrs: map[string]string{"uid": "1", "to": "aaa", "flag": ""},
			},
		},
	}
	for _, v := range tab {
		tag, err := ParseTag([]byte(v.tag))
		if err != nil {
			t.Fatal(err)
		}
		if !reflect.DeepEqual(tag, v.want) {
			t.Errorf("ParseTag(%q) = %v; want %v", v.tag, *tag, *v.want)
		}
	}
}