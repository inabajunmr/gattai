package path

import "testing"

func TestModifyToAbsoluteURL(t *testing.T) {
	// prepare sample html
	ModifyToAbsoluteURL(`<html>
							<body>
								<a href='/test1'>Test</a>
								<a href='/test2'>Test</a>
								<img src='/test3'>
								<img src='./test4'>
								<img src='./test5'>
								<img src='../test6'>
								<img src='../../test7'>
								<img src='test8'>
								<img src='test/test9'>
								<img src='http://example.com/test10'>
							</body>
						</html>`, "")
}
