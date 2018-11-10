package path

import (
	"strings"
	"testing"
)

var test = []struct {
	in1 string
	in2 string
	out string
}{
	{"a", "http://example.com", "http://example.com/a"},
	{"a", "http://example.com/b", "http://example.com/b/a"},
	{"/a", "http://example.com/b", "http://example.com/a"},
	{"/a", "http://example.com", "http://example.com/a"},
	{"a/b", "http://example.com", "http://example.com/a/b"},
	{"a/b", "http://example.com/c", "http://example.com/c/a/b"},
	{"/a/b", "http://example.com/c", "http://example.com/a/b"},
	{"../a", "http://example.com/b", "http://example.com/a"},
	{"../../a", "http://example.com/b/c", "http://example.com/a"},
	{"http://example.com", "http://example.com", "http://example.com"},
	{"http://example.com/a", "http://example.com/b", "http://example.com/a"},
}

func TestModifyToAbsoluteURL(t *testing.T) {

	result := ModifyToAbsoluteURLInHTML(strings.NewReader(`<html><head></head><body>
								<a href="/a">Test</a>
								<a href="/b">Test</a>
								<img src="/c">
								<img src="./d">
								<img src="./e">
								<img src="../f">
								<img src="../../g">
								<img src="h">
								<img src="i/j">
								<img src="http://example.com/k"></body></html>`), "http://example.com/l/m")

	expected := `<html><head></head><body>
								<a href="http://example.com/a">Test</a>
								<a href="http://example.com/b">Test</a>
								<img src="http://example.com/c"/>
								<img src="http://example.com/l/m/d"/>
								<img src="http://example.com/l/m/e"/>
								<img src="http://example.com/l/f"/>
								<img src="http://example.com/g"/>
								<img src="http://example.com/l/m/h"/>
								<img src="http://example.com/l/m/i/j"/>
								<img src="http://example.com/k"/></body></html>`

	if expected != result {
		t.Fatalf("Result is not expected value. result:%v, expected:%v", result, expected)
	}
}

func TestModifyToAbsoluteSingleURL(t *testing.T) {
	inputs := []struct {
		in1 string
		in2 string
		out string
	}{
		{"a", "http://example.com", "http://example.com/a"},
		{"a", "http://example.com/b", "http://example.com/b/a"},
		{"/a", "http://example.com/b", "http://example.com/a"},
		{"/a", "http://example.com", "http://example.com/a"},
		{"a/b", "http://example.com", "http://example.com/a/b"},
		{"a/b", "http://example.com/c", "http://example.com/c/a/b"},
		{"/a/b", "http://example.com/c", "http://example.com/a/b"},
		{"../a", "http://example.com/b", "http://example.com/a"},
		{"../../a", "http://example.com/b/c", "http://example.com/a"},
		{"http://example.com", "http://example.com", "http://example.com"},
		{"http://example.com/a", "http://example.com/b", "http://example.com/a"},
	}

	for _, v := range inputs {
		result := modifyToAbsoluteURL(v.in1, v.in2)
		if result != v.out {
			t.Fatalf("Result is not expected value. in1:%v, in2:%v, result:%v", v.in1, v.in2, result)
		}

	}
}
