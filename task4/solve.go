package main

import (
    "strings"
    "regexp"
)

func RemoveEven(slice []int) []int {
	result := make([]int, len(slice))
	j := 0
	for i := 0; i < len(slice); i++ {
		if slice[i]%2 != 0{
			result[j] = slice[i]
			j ++
		}
	}
	return result[0 : j]
}

func PowerGenerator(y int) func() int {
	deg := 1
	return func() int {
		deg *= y
		return deg
	}
}

func DifferentWordsCount(s string) int {
	str:=strings.ToLower(s)
	regxp := regexp.MustCompile("[a-z]+")
	words := regxp.FindAllString(str, -1)
	uniqueWords := make(map[string]bool)
	for i := 0; i < len(words); i++ {
		uniqueWords[words[i]] = true
	}
	return len(uniqueWords)
	
}