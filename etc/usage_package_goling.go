/*
Author: Gyu-Ho Lee
Update: 09/29/2013

http://golang.org/doc/code.html

mkdir $HOME/go
export GOPATH=$HOME/go
go get github.com/golingorg/goling

*********************
downloaded as 
/Users/user/go/src/github.com/golingorg/goling/goling.go
*********************

and run this code with
go run /path/etc/goling_usage.go

*** IMPORTANT ***
- do not put external package in the same folder 
  as your main package which uses it. 
- In Go, all packages must be in separate directories
- separate the packages into different directories
*/

package main

import "fmt"
import "github.com/golingorg/goling"

func main() {

	fmt.Println(goling.Clean_up("		Hello World!	This is Go 		"))
	// Hello World!	This is Go

	fmt.Println(goling.All_Tab_into_single_space("Hello	World	!	Hello"))
	// Hello World ! Hello

	fmt.Println(goling.All_Space_into_single_tab("Hello World! Hello"))
	// Hello	World!	Hello

	fmt.Println(goling.Tab_to_space("Hello	World	Hello"))
	// Hello World Hello

	fmt.Println(goling.Space_to_tab("Hello World Hello"))
	// Hello	World	Hello

	fmt.Println(goling.Delete_non_alphabet_with_space("Hello 1231thi21s123 is 12G33o25!!!!"))
	// Hello thi s is G o

	fmt.Println(goling.Delete_non_alphabet_without_space("Hello 1231thi21s123 is 12G33o25!!!!"))
	// Hello this is Go

	fmt.Println(goling.Delete_punctuation_with_space("Hello !@##%W@!#orl!@#!@#!@#d!"))
	// Hello W orl d

	fmt.Println(goling.Delete_punctuation_without_space("Hello !@##%W@!#orl!@#!@#!@#d!"))
	// Hello World

	fmt.Println(goling.Extract_number("I've bought 10 apples and 20 oranges."))
	// 10 20

	fmt.Println(goling.Extract_number_to_array("I've bought 10 apples and 20 oranges."))
	// [10 20]


	fmt.Println(goling.Split_by_word("Hello World! It is Good to See You."))
	// [Hello World It is Good to See You]

	fmt.Printf("%q\n", goling.Split_by_word("Hello World! It is Good to See You."))
	// ["Hello" "World" "It" "is" "Good" "to" "See" "You"]


	fmt.Println(goling.Split_by_sentence("Hello World! This is Go."))
	// [Hello World This is Go]

	fmt.Printf("%q\n", goling.Split_by_sentence("Hello World! This is Go."))
	// ["Hello World" "This is Go"]


	fmt.Println(goling.Split_by_paragraph("Hello World! \n This is Go!"))
	// [Hello World!   This is Go!]

	fmt.Printf("%q\n", goling.Split_by_paragraph("Hello World! \n This is Go!"))
	// ["Hello World!" "This is Go!"]


	fmt.Println(goling.Count_sentence("Hello World! This is Go."))
	// 2

	fmt.Println(goling.Count_word("Hello World! This is Go."))
	// 5

	fmt.Println(goling.Count_character("Hello Hi"))
	// 8

	fmt.Println(goling.Reverse_str("Hello Hi"))
	// iH olleH

	fmt.Println(goling.Check_palindrome("Anne, I vote more cars race Rome-to-Vienna"))
	// true

	fmt.Println(goling.Swap_case("Hello Hi"))
	// hELLO hI

	fmt.Println(goling.Insert_number_comma(21391239213))
	// 21,391,239,213

}