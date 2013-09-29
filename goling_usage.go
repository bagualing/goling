/*
http://golang.org/doc/code.html

// good idea to create a different GOPATH for each
$ mkdir $HOME/go
$ export GOPATH=$HOME/go
$ export PATH=$PATH:$GOPATH/bin
$ mkdir -p $GOPATH/src/github.com/golingorg
$ mkdir $GOPATH/src/github.com/golingorg/goling

copy goling.go
$ go install github.com/golingorg/goling


*/
package main

import "fmt"
import "goling"

func main() {

	fmt.Pritln(goling.Clean_up("		Hello World!	This is Go 		"))
	//

	fmt.Pritln(goling.All_Tab_into_single_space("Hello	World	!	Hello"))
	//

	fmt.Pritln(goling.All_Space_into_single_tab("Hello World! Hello"))
	//

	fmt.Pritln(goling.Tab_to_space("Hello	World	Hello"))
	//

	fmt.Pritln(goling.Space_to_tab("Hello World Hello"))
	//

	fmt.Pritln(goling.Delete_non_alphabet("Hello 1231thi21s123 is 12G33o25!!!!"))
	//

	fmt.Pritln(goling.Delete_punctuation("Hello !@##%W@!#orl!@#!@#!@#d!"))
	//

	fmt.Pritln(goling.Extract_number("I've bought 10 apples and 20 oranges."))
	//

	fmt.Pritln(goling.Extract_number_to_array("I've bought 10 apples and 20 oranges."))
	//

	fmt.Pritln(goling.Split_by_word("Hello World! It is Good to See You."))
	//

	fmt.Pritln(goling.Split_by_sentence("Hello World! This is Go."))
	//

	fmt.Pritln(goling.Split_by_paragraph("Hello World! \n This is Go!"))
	//

	fmt.Pritln(goling.Count_sentence("Hello World! This is Go."))
	//

	fmt.Pritln(goling.Count_word("Hello World! This is Go."))
	//

	fmt.Pritln(goling.Count_character("Hello Hi"))
	//

	fmt.Pritln(goling.Reverse_str("Hello Hi"))
	//

	fmt.Pritln(goling.Check_palindrome("Anne, I vote more cars race Rome-to-Vienna"))
	//

	fmt.Pritln(goling.Swap_case("Hello Hi"))
	//

	fmt.Pritln(goling.Insert_number_comma(21391239213))
	//

}