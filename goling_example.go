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


// delete non-alphabetic characters
func Delete_non_alphabet(input_str string) string {
	var validID = regexp.MustCompile(`[^A-Za-z]`)
	temp_str := validID.ReplaceAllString(input_str, " ")
	return Clean_up(temp_str)
}


func Delete_punctuation(input_str string) string {
	var validID = regexp.MustCompile(`[^A-Za-z0-9]`)
	temp_str := validID.ReplaceAllString(input_str, " ")
	return Clean_up(temp_str)
}


// return a new string that contains only numbers
func Extract_number(input_str string) string {
	var validID = regexp.MustCompile(`[^\d]`)
	temp_str := validID.ReplaceAllString(input_str, " ")
	return Clean_up(temp_str)
}


// extract numbers from string
// export it to an array of words
func Extract_number_to_array(input_str string) []string {
	var validID = regexp.MustCompile(`[^\d]`)
	temp_str := validID.ReplaceAllString(input_str, " ")
	temp_str = Clean_up(temp_str)
	return strings.Split(temp_str, " ")
}


func Split_by_word(input_str string) []string {
	temp_str := Clean_up(input_str)
	temp_str = Delete_punctuation(temp_str)
	temp_arr := strings.Split(temp_str, " ")

	// pop off the last unnecessary character
	if temp_arr[0] == "" || temp_arr[0] == " " {
		temp_arr = temp_arr[1:]
	}

	// delete the first unnessary character
	if temp_arr[len(temp_arr)-1] == "" || temp_arr[len(temp_arr)-1] == " " {
		temp_arr = temp_arr[:len(temp_arr)-1]
	}

	var final_arr []string
	for _, elem := range temp_arr {
		final_arr = append(final_arr, Clean_up(elem))
	}
	return final_arr
}


func Split_by_sentence(input_str string) []string {
	var validID = regexp.MustCompile(`[.!?]`)
	temp_str := validID.ReplaceAllString(input_str, "___")	
	temp_str = Clean_up(temp_str)
	temp_arr := strings.Split(temp_str, "___")

	// pop off the last unnecessary character
	if temp_arr[0] == "" || temp_arr[0] == " " {
		temp_arr = temp_arr[1:]
	}

	// delete the first unnessary character
	if temp_arr[len(temp_arr)-1] == "" || temp_arr[len(temp_arr)-1] == " " {
		temp_arr = temp_arr[:len(temp_arr)-1]
	}

	var final_arr []string
	for _, elem := range temp_arr {
		final_arr = append(final_arr, Clean_up(elem))
	}
	return final_arr
}


func Split_by_paragraph(input_str string) []string {
	return strings.Split(input_str, "\n")
}


func Count_sentence(input_str string) int {
	temp_arr := Split_by_sentence(input_str)
	return len(temp_arr) 
}


func Count_word(input_str string) int {
	temp_arr := Split_by_word(input_str)
	return len(temp_arr) 
}


// count including special characters
func Count_character(input_str string) int {
	temp_str := Clean_up(input_str)
	temp_arr := strings.Split(temp_str, "")
	return len(temp_arr) 
}


// swap case: upper to lower, vice versa
func Swap_case(input_str string) string {

}


func Capitalize_each_word(input_str string) string {

}


func Reverse_str(input_str string) string {

}


func Check_palindrome(input_str string) bool {

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

var str string = "   Hello,    World! 124  2 This 23is Go project,		It is amazing  . I am excited!    Are you too?    "

fmt.Println(str)
// "   Hello,    World! 124  2 This 23is Go project,		It is amazing  . I am excited!    Are you too?    "

fmt.Println(Clean_up(str))
// "Hello, World! 124 2 This 23is Go project, It is amazing . I am excited! Are you too?"

fmt.Println(str)
// "Hello, World! 124 2 This 23is Go project, It is amazing . I am excited! Are you too?"
// this is because a string object is immutable

str1 := Clean_up(str)
fmt.Println(str1)
// "Hello, World! 124 2 This 23is Go project, It is amazing . I am excited! Are you too?"
// now we have assigned the desired string

/**********************************************/

tab_str := "Hello World		Great"
fmt.Println(Clean_up(tab_str))                    // Hello World Great
fmt.Println(All_Tab_into_single_space(tab_str))   // Hello World Great
fmt.Println(All_Space_into_single_tab(tab_str))   // Hello	World	Great
fmt.Println(Tab_to_space(tab_str))    // Hello World  Great
fmt.Println(Space_to_tab(tab_str))    // Hello	World		Great

/**********************************************/

fmt.Println(Delete_non_alphabet(str))
// Hello World This is Go project It is amazing I am excited Are you too

fmt.Println(Delete_punctuation(str))
// Hello World 124 2 This 23is Go project It is amazing I am excited Are you too

/**********************************************/

fmt.Println(Extract_number(str))           // 124 2 23
fmt.Println(Extract_number_to_array(str))  // [124 2 23]

/**********************************************/

fmt.Println(Split_by_word(str))
arr1 := Split_by_word(str)
fmt.Println("# of words:", len(arr1))   // 17
fmt.Println(arr1[len(arr1)-2])  // you

fmt.Println(Split_by_sentence(str))
arr2 := Split_by_sentence(str)
fmt.Println("# of sentences:", len(arr2))  // 4
fmt.Println(arr2[len(arr2)-2])  // I am excited

/**********************************************/

fmt.Println(Count_sentence(str))    // 4
fmt.Println(Count_word(str))        // 17
fmt.Println(Count_character(str))   // 84

/**********************************************/

fmt.Println(Swap_case(str))
fmt.Println(Capitalize_str(str))

/**********************************************/

fmt.Println(Reverse(str))
fmt.Println(Check_palindrome(str))

/**********************************************/

fmt.Println(Insert_number_comma(str))

/**********************************************/

fmt.Println(Select_split(str))



//**********************************************
//             Final Example
//**********************************************

var long_text string = "This is from wikpedia.\nThe syntax of Go is broadly similar to that of C: blocks of code are surrounded with curly braces; common control flow structures include for, switch, and if. Unlike C, line-ending semicolons are optional, variable declarations are written differently and are usually optional, type conversions must be made explicitly, and new go and select control keywords have been introduced to support concurrent programming. New built-in types include maps, array slices, and channels for inter-thread communication.\nGo is designed for exceptionally fast compiling times, even on modest hardware. The language requires garbage collection. Some concurrency-related structural conventions of Go (channels and alternative channel inputs) are borrowed from Tony Hoare's CSP. Unlike previous concurrent programming languages such as occam or Limbo, Go does not provide any built-in notion of safe or verifiable concurrency.\nOf features found in C++ or Java, Go does not include type inheritance, generic programming, assertions, method overloading, or pointer arithmetic. With respect to these omissions, the Go authors express an openness to generic programming, explicitly argue against assertions and pointer arithmetic, while defending the choice to omit type inheritance as giving a more useful language, encouraging heavy use of interfaces instead. Initially, the language did not include exception handling, but in March 2010 a mechanism known as panic/recover was implemented to handle exceptional errors while avoiding some of the problems the Go authors find with exceptions."

temp_long_arr := Split_by_paragraph(long_text)
fmt.Println(temp_long_arr[1])
/*
The syntax of Go ....
*/








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

split reverse join

/**********************************************/
// check if a word is a palindrome or not


var str_1 = "never odd or even"  // true
var str_2 = "racecar"   // true
var str_3 = "eye"       // true
var str_4 = "place"     // false
var str_5 = "Anne, I vote more cars race Rome-to-Vienna"  // true
var str_6 = "Noel - sees Leon" // true
var str_7 = "A war at Tarawa!" // true

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
// no s to avoid confusion

// string is an immutable object
// no method changes the original string

// clean up space characters between words
// replace them with single-space
// delete unnecessary space character at the beginning and end
// "Hello, World! 124 2 This 23is Go,		Great"


*/