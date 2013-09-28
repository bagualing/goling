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
	// "strconv"
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


// rune is an alias for int32
// to distinguish character values from integer values
func Reverse_str(input_str string) string {
	runes := []rune(input_str)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}


func Check_palindrome(input_str string) bool {
	temp_str := Clean_up(input_str)
	temp_str = Delete_punctuation(temp_str)
	temp_str = strings.ToLower(temp_str)

	var validID = regexp.MustCompile(`\s{1,}`)
	temp_str = validID.ReplaceAllString(temp_str, "")

	rev_str := Reverse_str(temp_str)

	if temp_str == rev_str {
		return true
	} else {
		return false
	}
}


/*

// insert comma between digits  100000 -> 100,000
func Insert_number_comma(input_num int) string {
	temp_str := strconv.Itoa(input_num)
	temp_str = Clean_up(temp_str)
	var validID = regexp.MustCompile(`\B(?=(\d{3})+$)`)
	return validID.ReplaceAllString(temp_str, ",")
}


// swap case: upper to lower, vice versa
func Swap_case(input_str string) string {
	temp_str := Clean_up(input_str)


	var validID_1 = regexp.MustCompile(`[a-z]`)
	str1 := validID_1.ReplaceAllString(temp_str, strings.ToUpper(temp_str))

	var validID_2 = regexp.MustCompile(`[A-Z]`)
	str2 := validID_2.ReplaceAllString(temp_str, strings.ToLower(temp_str))

	return str1 + str2

}

*/


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

fmt.Println(Reverse_str("Hello this is GoLing.org project"))
// tcejorp gro.gniLoG si siht olleH

fmt.Println(Check_palindrome("never odd or even")) // true
fmt.Println(Check_palindrome("racecar")) // true
fmt.Println(Check_palindrome("eye")) // true
fmt.Println(Check_palindrome("place")) // false
fmt.Println(Check_palindrome("Anne, I vote more cars race Rome-to-Vienna")) // true
fmt.Println(Check_palindrome("Noel - sees Leon")) // true
fmt.Println(Check_palindrome("A war at Tarawa!")) // true

/**********************************************/

/*
fmt.Println(Insert_number_comma(100000000))  
// 100,000,000

fmt.Println(Swap_case("Hey I am Coding RIGHT now"))
// hEY i AM cODING right NOW
*/


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