package path

import "testing"

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
	// prepare sample html
	// ModifyToAbsoluteURL(`<html>
	// 						<body>
	// 							<a href='/test1'>Test</a>
	// 							<a href='/test2'>Test</a>
	// 							<img src='/test3'>
	// 							<img src='./test4'>
	// 							<img src='./test5'>
	// 							<img src='../test6'>
	// 							<img src='../../test7'>
	// 							<img src='test8'>
	// 							<img src='test/test9'>
	// 							<img src='http://example.com/test10'>
	// 						</body>
	// 					</html>`, "")
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
