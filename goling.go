/*
Author: Gyu-Ho Lee
Update: 09/29/2013

Description: Source Code for goling package.
*/

package goling

import (
	"strings"
	"regexp"
	"strconv"
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
func Delete_non_alphabet_with_space(input_str string) string {
	var validID = regexp.MustCompile(`[^A-Za-z]`)
	temp_str := validID.ReplaceAllString(input_str, " ")
	return Clean_up(temp_str)
}


func Delete_non_alphabet_without_space(input_str string) string {
	var validID = regexp.MustCompile(`[^A-Za-z\s]`)
	temp_str := validID.ReplaceAllString(input_str, "")
	return Clean_up(temp_str)
}


func Delete_punctuation_with_space(input_str string) string {
	var validID = regexp.MustCompile(`[^A-Za-z0-9]`)
	temp_str := validID.ReplaceAllString(input_str, " ")
	return Clean_up(temp_str)
}


func Delete_punctuation_without_space(input_str string) string {
	var validID = regexp.MustCompile(`[^A-Za-z0-9\s]`)
	temp_str := validID.ReplaceAllString(input_str, "")
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
	temp_str = Delete_punctuation_with_space(temp_str)
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
	temp_arr := strings.Split(input_str, "\n")
	var final_arr []string
	for _, elem := range temp_arr {
		final_arr = append(final_arr, Clean_up(elem))
	}
	return final_arr
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
	temp_str = Delete_punctuation_without_space(temp_str)
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


func swap_rune(r rune) rune {
	switch {
		case 'a' <= r && r <= 'z':
			return r - 'a' + 'A'
		case 'A' <= r && r <= 'Z':
			return r - 'A' + 'a'
		default:
			return r
	}
}

func Swap_case(str string) string {
	return strings.Map(swap_rune, str)
}


// insert comma between digits  100000 -> 100,000
func Insert_number_comma(input_num int) string {
	temp_str := strconv.Itoa(input_num)
	var num_arr []string
	i := len(temp_str)%3;
	if i == 0 { i = 3 }
	for index, elem := range strings.Split(temp_str, "") {
		if i == index {
			num_arr = append(num_arr, ",");
			i += 3;
		}
		num_arr = append(num_arr, elem)
	}
	return strings.Join(num_arr, "")
}