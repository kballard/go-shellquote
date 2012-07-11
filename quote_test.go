package shellquote

import (
	"testing"
)

func TestSimpleJoin(t *testing.T) {
	for _, elem := range simpleTest {
		output := Join(elem.input...)
		if output != elem.output {
			t.Errorf("Input %q, got %q, expected %q", elem.input, output, elem.output)
		}
	}
}

var simpleTest = []struct {
	input  []string
	output string
}{
	{[]string{"test"}, "test"},
	{[]string{"hello goodbye"}, "'hello goodbye'"},
	{[]string{"hello", "goodbye"}, "hello goodbye"},
	{[]string{"don't you know the dewey decimal system?"}, "'don'\\''t you know the dewey decimal system?'"},
	{[]string{"don't", "you", "know", "the", "dewey", "decimal", "system?"}, "don\\'t you know the dewey decimal system\\?"},
	{[]string{"~user", "u~ser", " ~user", "!~user"}, "\\~user u~ser ' ~user' \\!~user"},
	{[]string{"foo*", "M{ovies,usic}", "ab[cd]", "%3"}, "foo\\* M\\{ovies,usic} ab\\[cd] %3"},
}
