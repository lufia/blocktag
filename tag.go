package blocktag

import (
	"bytes"
	"errors"
	"io"
	"strings"
	"text/scanner"
)

var (
	errEmptyTag = errors.New("empty tag")
	errSyntax   = errors.New("syntax error")
)

// Tag はブロックの開始タグをあらわす。
type Tag struct {
	Name  string
	Attrs map[string]string
}

// ParseTag は開始タグをパースする。
func ParseTag(tag []byte) (*Tag, error) {
	a := strings.Fields(string(tag))
	if len(a) == 0 {
		return nil, errEmptyTag
	}

	var p Tag
	p.Name = a[0]
	for _, s := range a[1:] {
		kv := strings.SplitN(s, "=", 2)
		switch len(kv) {
		case 1:
			if p.Attrs == nil {
				p.Attrs = make(map[string]string)
			}
			p.Attrs[kv[0]] = ""
		case 2:
			if p.Attrs == nil {
				p.Attrs = make(map[string]string)
			}
			p.Attrs[kv[0]] = kv[1]
		}
	}
	return &p, nil
}

// Block は1つの[tag attr=value]text[/tag]を表す。
type Block struct {
	Tag  *Tag
	Body []byte
}

type stream struct {
	s scanner.Scanner
}

func (s *stream) advance(c rune) ([]byte, bool) {
	var buf bytes.Buffer
	for {
		c1 := s.s.Next()
		switch {
		case c1 == scanner.EOF:
			return buf.Bytes(), false
		case c1 == c:
			return buf.Bytes(), true
		}
		buf.WriteRune(c1)
	}
}

func (s *stream) readBlock() (*Block, error) {
	_, ok := s.advance('[')
	if !ok {
		return nil, nil
	}
	tag, err := s.readTag()
	if err != nil {
		return nil, err
	}
	body, err := s.advanceUntil(tag)
	if err != nil {
		return nil, err
	}
	return &Block{Tag: tag, Body: body}, nil
}

func (s *stream) readTag() (*Tag, error) {
	tag, ok := s.advance(']')
	if !ok {
		return nil, errSyntax
	}
	return ParseTag(tag)
}

func (s *stream) advanceUntil(tag *Tag) ([]byte, error) {
	var buf bytes.Buffer

	for {
		body, ok := s.advance('[')
		if !ok {
			return nil, errSyntax
		}
		buf.Write(body)

		t, ok := s.advance(']')
		if !ok {
			return nil, errSyntax
		}
		p, err := ParseTag(t)
		if err != nil {
			return nil, err
		}
		if p.Name[0] == '/' && strings.TrimSpace(p.Name[1:]) == tag.Name {
			break
		}
		buf.WriteByte('[')
		buf.Write(t)
		buf.WriteByte(']')
	}
	return buf.Bytes(), nil
}

// Parse は[tag][/tag]の中に書かれている文字列を取り出して返す。
// タグのネストは行わない。ネストしているタグはそのまま文字として扱う。
// タグの外にある文字は無視する。
// 何もタグが見つからない場合はnilを返してエラーにはならない。
func Parse(r io.Reader) ([]*Block, error) {
	var a []*Block

	var fin stream
	fin.s.Init(r)
	for {
		b, err := fin.readBlock()
		if err != nil {
			return a, err
		}
		if b == nil { // EOF
			return a, nil
		}
		a = append(a, b)
	}
}
