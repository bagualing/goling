/*
Author: Gyu-Ho Lee
Update: 09/27/2013

Description: Source Code for goling package.
*/

package goling

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