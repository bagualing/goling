## goling

### Description
GoLing is a natural language processing project with programming language Go. goling provides a simple text processing tools. It is meant to be a preliminary tool before ultimate text processing tasks. It supports cleaning-up and extracting certain characters, so that we do not have to worry about unexpected or unnecessary chracters.

### Installation
`go get github.com/golingorg/goling`

### Usage
```go
package main

import (
	"github.com/golingorg/goling"
	"fmt"
)

func main() {
	// example
	var ex_string string = "   Hello,    World! 124  2 This 23is Go, Great  "

	// I am coding right now. Should be ready very soon

	ex_string.Cleanup()
	// clean up space characters between words
	// replace them with single-space
	// remove the unnecessary spaces at the beginning and end
	// "Hello, World! 124 2 This 23is Go, Great"

	fmt.Println(ex_string)

	ex_string.Delnon_abc()
	// using regex
	// delete non-alphabetic characters
	// Hello World This is Go Great

	ex_string.Delnon_num()
	// using regex
	// delete non-numeric characters
	// 124 2 23

	ex_string.SplitWord()
	// using regex
	// split string by words

	ex_string.SplitSen()
	// using regex
	// split string by sentences

	fmt.Println(ex_string)
	fmt.Println(ex_string)
	fmt.Println(ex_string)		
}
```


### Documentation (http://goling.org)
Documentation can be found at [goling.org](http://godoc.org/).
For more detailed documentation, read the source.

### History
001: Friday, September 27th, 2013
- Launch the first update!