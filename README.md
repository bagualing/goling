## goling

### Description
GoLing is a natural language processing project with programming language Go. goling provides a simple text processing tool, inspired by regular expression. The main goal is "Let's not worry about the trivial and get to the point. And be the easiest tool in the world." In this sense, this is meant to be a preliminary, still useful, tool for crucial text processing tasks.

### Installation
`go get github.com/golingorg/goling`

### Usage
```go
/*
Author: Gyu-Ho Lee
Update: 09/27/2013

Description: Source code example with tutorialistic comments. For more detail, please visit documentation page.
*/
package main

import (
	"fmt"
	"strings"
	"regexp"
)

func Clean_up(input_str string) string {
 
	// compile regular expression
	// to match 2-or-more-than-2 whitespaces
	var validID = regexp.MustCompile(`\s{2,}`)
	var temp_str string = validID.ReplaceAllString(input_str, " ")
	var temp_arr []string = strings.Split(temp_str, " ")

	// pop off the last unnecessary character
	if temp_arr[0] == "" || temp_arr[0] == " " {
		temp_arr = temp_arr[1:]
	}

	// delete the first unnessary character
	if temp_arr[len(temp_arr)-1] == "" || temp_arr[len(temp_arr)-1] == " " {
		temp_arr = temp_arr[:len(temp_arr)-1]
	}

	// return in string format
	return strings.Join(temp_arr, " ")
}


func All_Tab_into_single_space(input_str string) string {
	// to take any tab chracters: single tab, double tabs, ...
	var validID = regexp.MustCompile(`\t{1,}`)
	return validID.ReplaceAllString(input_str, " ")
}

func All_Space_into_single_tab(input_str string) string {
	// to take any whitespace characters: single whitespace, doulbe _, ...
	var validID = regexp.MustCompile(`\s{1,}`)
	return validID.ReplaceAllString(input_str, "	")
}

func Tab_to_space(input_str string) string {
	var validID = regexp.MustCompile(`\t`)
	return validID.ReplaceAllString(input_str, " ")
}

func Space_to_tab(input_str string) string {
	var validID = regexp.MustCompile(`\s`)
	return validID.ReplaceAllString(input_str, "	")
}


func main() {

/*   practice
var str string = "I like it"
fmt.Printf("%q\n", strings.Split(str, " "))
// ["I" "like" "it"]

var str2 string = "I like it"
var temp_arr []string = strings.Split(str2, " ")
fmt.Printf("%q\n", temp_arr)
// ["I" "like" "it"]

fmt.Println(len(str1), len(str2))  // 9 9
fmt.Println(len(temp_arr))  // 3
*/

/**********************************************/

var str string = "   Hello,    World! 124  2 This 23is Go,		Great  "

fmt.Println(str)
// "   Hello,    World! 124  2 This 23is Go,		Great  "

fmt.Println(Clean_up(str))
// "Hello, World! 124 2 This 23is Go, Great"

fmt.Println(str)
// "   Hello,    World! 124  2 This 23is Go,		Great  "
// this is because a string object is immutable

str1 := Clean_up(str)
fmt.Println(str1)
// "Hello, World! 124 2 This 23is Go, Great"
// now we have assigned the desired string

/**********************************************/

tab_str := "Hello World		Great"
fmt.Println(All_Tab_into_single_space(tab_str))   // Hello World Great
fmt.Println(All_Space_into_single_tab(tab_str))   // Hello	World	Great
fmt.Println(Tab_to_space(tab_str))    // Hello World  Great
fmt.Println(Space_to_tab(tab_str))    // Hello	World		Great

/**********************************************/

// delete non-alphabetic characters
// Hello World This is Go Great
func Delete_non_alphabet(input_str string) string {

}
var str1 = str.replace(/[^A-Z]/gi, " "); // delete numbers, special characters

/**********************************************/

// delete non-numeric characters
// 124 2 23
func Delete_non_numeric(input_str string) string {

}
	// delete non-numeric characters
	var str1 = str.replace(/[^\d]/gi, " ");
/**********************************************/

// split a string by words
func Split_by_word(input_str string) []string {

}
/**********************************************/

// split a string by sentences
func Split_by_sen(input_str string) []string {

}

/**********************************************/

// count words
func Count_word(input_str string) int {

}

/**********************************************/

// count characters
func Count_character(input_str string) int {

}

/**********************************************/

// swap case: upper to lower, vice versa
func Swap_case(input_str string) string {

}
function SwapCase(str) {
	// $0 = entire string
	// $1, $2, etc = value from the last successful match
	// $1 == matches lowercase only
	// $2 == matches uppercase only
	// checking which group is non-empty = case of the original

	return str.replace(/([a-z])|([A-Z])/g,
		function($0, $1, $2) {
			return ($1) ? $0.toUpperCase() : $0.toLowerCase();
		})
}

/**********************************************/
// capitalize the first letter of each word
func Capitalize_str(input_str string) string {

}
function LetterCapitalize(str) {
	var word_arr = str.split(" ")
	var cap_arr = [];
	for (var elem in word_arr) {
		cap_arr.push(word_arr[elem].charAt(0).toUpperCase() 
						+ word_arr[elem].slice(1));
	}
	return cap_arr.join(" ");
}
/**********************************************/

// reverse the order of chracters
func Reverse(input_str string) string {

}
split reverse join

/**********************************************/
// check if a word is a palindrome or not
func Check_palindrome(input_str string) bool {

}
function Palindrome(str) {
	// delete all non-alphabetic characters
	// delete all white spaces
	// make them lowercase
	var char_str = str.replace(/[^A-Z]/gi, "").toLowerCase();

	// reverse string
	var rev_str = char_str.split("").reverse().join("");

	if (char_str == rev_str) return "true";
	else return "false";
}
/**********************************************/
// insert comma between digits
// 100000 -> 100,000
func Insert_number_comma(input_str string) string {

}
str.replace(/\B(?=(\d{3})+$)/g, ",");
/**********************************************/

// select a word and split afterwords
// and join with spaces between words
func Select_split(input_str string) string {

}

var str = "Hey Isdfadsdaydaqrqw1dsv good day I";
str.split(/(day)/g);
// ["Hey Isdfads", "day", "daqrqw1dsv good ", "day", " I"]
// does NOT delete "day"
/**********************************************/

/**********************************************/
// Demonstration
/**********************************************/

	var str string = "   Hello,    World! 124  2 This 23is Go,		Great  "

	fmt.Println(ex_string)
	fmt.Println(ex_string)
	fmt.Println(ex_string)		
}

/*
words to go to documentation

// function naming
// to keep it consistent
// only first letter in upper case
// words divided by underscore(_)
// string is an immutable object
// no method changes the original string

// clean up space characters between words
// replace them with single-space
// delete unnecessary space character at the beginning and end
// "Hello, World! 124 2 This 23is Go,		Great"


*/
```


### Documentation (http://goling.org)
Documentation can be found at [goling.org](http://godoc.org/).
For more detailed documentation, read the source.

### History
001: Friday, September 27th, 2013
- Launch the first update!