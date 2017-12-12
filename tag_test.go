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
				Name: "tag",
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
		{
			tag: "name:1",
			want: &Tag{
				Name:  "name",
				Attrs: map[string]string{"value": "1"},
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
