package main

import (
	"fmt"
	"strings"
)

func main() {
	SearchingChallenge("abcracecar")
}

var (
	pallin         []string
	pallindrome    string
	interparseStr  string
	challengeToken string = "123456"
	pallindlen     int
	chTokenlen     int
)

func SearchingChallenge(str string) {
	fmt.Printf("Input: %s\n", str)
	str = strings.ToLower(str)
	check(str)

	pallindrome = getTheBigger()
	fmt.Printf("Output: %s\n", pallindrome)

	pallindlen = len(pallindrome)
	chTokenlen = len(challengeToken)

	interparse(1)
	fmt.Printf("Final output: %s\n", interparseStr)
}

func check(str string) {
	var lenstr = len(str)

	for i := 0; i < lenstr; i++ {

		checking := str[0 : i+1]
		inverted := invert(checking)

		if checking == inverted {
			pallin = append(pallin, checking)
		}
	}

	if lenstr > 1 {
		str = str[(lenstr - (lenstr - 1)):]
		check(str)
	}
}

func invert(str string) (inverted string) {
	var lenstr = len(str)
	for i := lenstr; i > 0; i-- {
		inverted += str[i-1 : i]
	}
	return
}

func getTheBigger() string {
	posMajor := 0
	count2 := 0

	for i := 0; i < len(pallin); i++ {
		count1 := len(pallin[i])

		if count1 > count2 {
			count2 = count1
			posMajor = i
		}
	}

	if posMajor == 0 {
		pallin = []string{"none"}
	}
	return pallin[posMajor]
}

func interparse(count int) {

	if pallindlen > 0 {
		pallindlen--
		interparseStr += pallindrome[count-1 : count]
	}

	if chTokenlen > 0 {
		chTokenlen--
		interparseStr += challengeToken[count-1 : count]
	}

	if pallindlen <= 0 && chTokenlen <= 0 {
		return
	}

	count++
	interparse(count)
}
