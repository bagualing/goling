// function naming
// to keep it consistent
// only first letter in upper case
// words divided by underscore(_)

package main

import (
	"fmt"
	"regexp"
)
// string is an immutable object
// no method changes the original string







	// delete non-numeric characters
	var str1 = str.replace(/[^\d]/gi, " ");

	// delete 2-3 spaces
	var str2 = str1.replace(/\s{2,}/g, " ");

	function WordCount(str) {
	var temp = str.split(" ");
	return temp.length;
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
*/

func main() {
	// example
	var str string = "   Hello,    World! 124  2 This 23is Go,		Great  "

	// I am coding right now. Should be ready very soon

	str.Clean_up()


// clean up space characters between words
// replace them with single-space
// delete unnecessary space character at the beginning and end
// "Hello, World! 124 2 This 23is Go,		Great"
func Clean_up(input_str string) string {


}
var str2 = str1.replace(/\s{2,}/g, " "); // delete 2-3 spaces
var word_arr = str2.split(" ");
if (word_arr[0] == "") word_arr.shift();
if (word_arr[word_arr.length-1] == "") word_arr.pop();


// delete non-alphabetic characters
// Hello World This is Go Great
func Del_non_abc(input_str string) string {

}
var str1 = str.replace(/[^A-Z]/gi, " "); // delete numbers, special characters


// delete non-numeric characters
// 124 2 23
func Del_non_num(input_str string) string {

}

// split a string by words
func Split_by_word(input_str string) []string {

}

// split a string by sentences
func Split_by_sen(input_str string) []string {

}


// count words
func Count_word(input_str string) int {

}


// count characters
func Count_char(input_str string) int {

}

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

// capitalize the first letter of each word
func Cap_str(input_str string) string {

}

// reverse the order of chracters
func Reverse(input_str string) string {

}

// check if a word is a palindrome or not
func Palindrome(input_str string) bool {

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

// insert comma between digits
// 100000 -> 100,000
func Comma_num(input_str string) string {

}
str.replace(/\B(?=(\d{3})+$)/g, ",");


// select a word and split afterwords
// and join with spaces between words
func Select_split(input_str string) string {

}

var str = "Hey Isdfadsdaydaqrqw1dsv good day I";
str.split(/(day)/g);
// ["Hey Isdfads", "day", "daqrqw1dsv good ", "day", " I"]
// does NOT delete "day"

	str.Del_non_abc()

	str.Del_non_num()

	str.Split_by_word()

	str.Split_by_sen()

	str.Count_word()

	str.Count_char()

	str.Swap_case()

	str.Cap_str()

	str.Reverse()

	str.Palindrome()

	str.Comma_num()

	fmt.Println(ex_string)
	fmt.Println(ex_string)
	fmt.Println(ex_string)		
}