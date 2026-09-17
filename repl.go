package main

func cleanInput(text string) []string {
	var result []string
	var temp string
	for _, character := range text {
		if string(character) == " " {
			if len(temp) > 1 {
				result = append(result, temp)
				temp = ""
			}
			continue
		}
		temp = temp + string(character)
	}

	if len(temp) > 1 {
		result = append(result, temp)
	}

	return result
}
