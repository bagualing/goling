// function naming
// to keep it consistent
// only first letter in upper case
// words divided by underscore(_)

package main

import "fmt"
// string is an immutable object
// no method changes the original string


/*
my personal project in Javascript


	var str1 = str.replace(/[^A-Z]/gi, " "); // delete numbers, special characters
	var str2 = str1.replace(/\s{2,}/g, " "); // delete 2-3 spaces
	var word_arr = str2.split(" ");
	if (word_arr[0] == "") word_arr.shift();
	if (word_arr[word_arr.length-1] == "") word_arr.pop();
	// shift, pop to delete
	// unshift, push to add

	str.replace(/\B(?=(\d{3})+$)/g, ",");


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

var str = "Hey Isdfadsdaydaqrqw1dsv good day I";
str.split(/(day)/g);
// ["Hey Isdfads", "day", "daqrqw1dsv good ", "day", " I"]
// does NOT delete "day"


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
	var str string = "   Hello,    World! 124  2 This 23is Go, Great  "

	// I am coding right now. Should be ready very soon

	str.Clean_up()

	str.Del_non_abc()

	str.Del_non_num()

	str.Split_word()

	str.Split_sen()

	str.Count_word()

	str.Count_char()

	str.Swap_case()

	str.Cap_word()

	str.Reverse()

	str.Palindrome()

	str.Comma_num()

	fmt.Println(ex_string)
	fmt.Println(ex_string)
	fmt.Println(ex_string)		
}