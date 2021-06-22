package main

import (
	"io"
	"io/ioutil"
	"regexp"
	"strings"

	"golang.org/x/net/html"
)

type Parser interface {
	GetWordCounts() map[string]int
}

func NormaliseString(s string) string {
	s = strings.ToLower(s)
	re := regexp.MustCompile("[^a-z'_-]")
	return re.ReplaceAllString(s, "")
}

type PlainTextParser struct {
	Text string
}

func NewPlainTextParser(r io.Reader) (*PlainTextParser, error) {
	data, err := ioutil.ReadAll(r)
	if err != nil {
		return nil, err
	}
	return &PlainTextParser{Text: string(data)}, nil
}

func (parser *PlainTextParser) GetWordCounts() map[string]int {
	counts := make(map[string]int)
	for _, word := range strings.Fields(parser.Text) {
		normalised := NormaliseString(word)
		if normalised == "" {
			continue
		}
		counts[normalised]++
	}
	return counts
}

func GetTextNodes(tokens *html.Tokenizer) (textNodes []string) {
	textNodes = make([]string, 0)
	tagsContainingText := []string{
		"a", "b", "button", "em", "p", "h1", "h2", "h3", "h4", "h5", "h6", "span", "strong",
	}
	var useContents bool
	for {
		t := tokens.Next()
		token := tokens.Token()

		switch t {
		case html.StartTagToken:
			useContents = false
			for _, tag := range tagsContainingText {
				if token.Data == tag {
					useContents = true
					break
				}
			}
		case html.TextToken:
			if useContents {
				textNodes = append(textNodes, token.Data)
			}
		case html.ErrorToken:
			return
		}
	}
}

type HTMLParser struct {
	TextNodes []string
}

func NewHTMLParser(r io.Reader) (*HTMLParser, error) {
	tokenizer := html.NewTokenizer(r)
	textNodes := GetTextNodes(tokenizer)
	return &HTMLParser{TextNodes: textNodes}, nil
}

func (parser *HTMLParser) GetWordCounts() map[string]int {
	counts := make(map[string]int)
	for _, textNode := range parser.TextNodes {
		for _, word := range strings.Fields(textNode) {
			normalised := NormaliseString(word)
			if normalised == "" {
				continue
			}
			counts[normalised]++
		}
	}
	return counts
}
