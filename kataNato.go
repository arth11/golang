package main

import (
	"regexp"
	"strings"
)

//ToNato Example function Kata If you can read
func ToNato(words string) string {
	var ToBe bool = true

	nato := []string{"Alfa", "Bravo", "Charlie", "Delta", "Echo", "Foxtrot", "Golf", "Hotel", "India", "Juliett", "Kilo", "Lima", "Mike", "November", "Oscar", "Papa", "Quebec", "Romeo", "Sierra", "Tango", "Uniform", "Victor", "Whiskey", "Xray", "Yankee", "Zulu"}
	reg, _ := regexp.Compile("[!.,-/:-@[-`{-]")
	var str strings.Builder

	for _, word := range words {
		if reg.MatchString(string(word)) {
			if ToBe == false {
				str.WriteString(" " + string(word))
				ToBe = true
			} else {
				str.WriteString(string(word))
				ToBe = true
			}

		}
		strings.TrimSpace(string(word))

		for _, nat := range nato {
			if strings.HasPrefix(nat, strings.ToUpper(string(word))) {
				//ToBe = false
				if ToBe == true {
					str.WriteString(nat)
					ToBe = false
				} else {
					str.WriteString(" " + nat)
				}

			}

		}

	}

	return str.String()
}
