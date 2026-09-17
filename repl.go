package main

func cleanInput(text string) []string {
	var result []string
	var temp string
	for i, character := range text {
		if string(text[i]) == " " {
			if len(temp) >= 1 {
				result = append(result, temp)
				temp = ""
			}
			continue
		}
		temp = temp + string(character)
	}
	return result
}
