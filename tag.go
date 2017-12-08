package blocktag

// Block は1つの[tag attr=value]text[/tag]を表す。
type Block struct {
	Tag   string
	Attrs map[string]string
	Text  string
}

// Parse は[tag][/tag]の中に書かれている文字列を取り出して返す。
// タグのネストは行わない。ネストしているタグはそのまま文字として扱う。
// タグの外にある文字は無視する。
// 何もタグが見つからない場合はnilを返してエラーにはならない。
func Parse(r io.Reader) ([]*Block, error) {
	return nil, nil
}
