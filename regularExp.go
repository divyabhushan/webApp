package main

import (
	"fmt"
	"regexp"
)

func main() {
	var matchStr, _ = regexp.Compile("[^abc]")
	var numeralsStr, _ = regexp.Compile("[0-9]+")
	var alphaStr, _ = regexp.Compile("[a-z]+")

	searchInStr := "abcd1234abcdefghabc09343"

	//match the leftmost string
	fmt.Println("Match the leftmost string :", matchStr.FindString(searchInStr))

	//match all the occurence of the matched pattern
	fmt.Println("match all the occurences of the matched pattern", matchStr.FindAllString(searchInStr, -1))

	//match and print only numerals
	fmt.Println("Match only leftmost numerals", numeralsStr.FindString(searchInStr))

	fmt.Println("Match all the numerals:", numeralsStr.FindAllString(searchInStr, -1))

	fmt.Println("Match all the alphabets:", alphaStr.FindAllString(searchInStr, -1))
}
