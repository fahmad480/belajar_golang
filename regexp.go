package main

import (
	"fmt"
	"regexp"
)

func main() {
	var regex *regexp.Regexp = regexp.MustCompile("e([a-z])a")

	fmt.Println(regex.MatchString("faraaz"))
	fmt.Println(regex.MatchString("eta"))
	fmt.Println(regex.MatchString("eda"))

	search := regex.FindAllString("faraaz eko eda eta eya eki", -1)
	fmt.Println(search)
}
