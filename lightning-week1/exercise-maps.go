package main

import (
	//"golang.org/x/tour/wc"
  "fmt"
  "regexp"
  "strings"
)

func WordCount(s string) map[string]int {
  reg, _ := regexp.Compile("[^a-zA-Z]+")
  var stringMap = make(map[string]int)
  var stringSlice = strings.Fields(s)
  for _,v := range stringSlice {
    stringMap[reg.ReplaceAllString(v,"")]++
  }
  fmt.Println(stringSlice)
	return stringMap
}

func main() {
	//wc.Test(WordCount)
  fmt.Println(WordCount("This is a sentence of words. This, is, also a sentence of words. This is yet another sentence of words."))
}
