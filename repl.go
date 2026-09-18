package main

import "strings"

func cleanInput(text string) []string {
	var result []string
	var temp string
	for _, character := range text {
		if string(character) == " " {
			if len(temp) > 1 {
				result = append(result, strings.ToLower(temp))
				temp = ""
			}
			continue
		}
		temp = strings.ToLower(temp + string(character))
	}

	if len(temp) > 1 {
		result = append(result, strings.ToLower(temp))
	}

	return result
}
