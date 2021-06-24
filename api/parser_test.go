package main

import (
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func BenchmarkNormaliseString(b *testing.B) {
	for i := 0; i < b.N; i++ {
		NormaliseString(`"Hello123"`)
	}
}

func TestPlainTextWordCounts(t *testing.T) {
	text := "Tell the audience what you're going to say. Say it. Then tell them what you've said."
	parser, err := NewPlainTextParser(strings.NewReader(text))
	assert.Nil(t, err)
	assert.Equal(t, map[string]int{
		"tell":     2,
		"the":      1,
		"audience": 1,
		"what":     2,
		"you're":   1,
		"going":    1,
		"to":       1,
		"say":      2,
		"it":       1,
		"then":     1,
		"them":     1,
		"you've":   1,
		"said":     1,
	}, parser.GetWordCounts())
}

func TestHTMLWordCounts(t *testing.T) {
	html := `
	<html>
		<h1>Hello</h1>
		<p>Hello, world. Some <strong>sample text</strong></p>
		<span>...<a href="#">find out more</a></span>
	</html>
	`
	parser, err := NewHTMLParser(strings.NewReader(html))
	assert.Nil(t, err)
	assert.Equal(t,
		map[string]int(map[string]int{
			"find":   1,
			"hello":  2,
			"more":   1,
			"out":    1,
			"sample": 1,
			"some":   1,
			"text":   1,
			"world":  1,
		}),
		parser.GetWordCounts(),
	)
}
