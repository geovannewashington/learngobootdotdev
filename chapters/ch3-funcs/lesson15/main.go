package main

func reformat(message string, formatter func(string) string) string {	
	var result string
	result = formatter(message)
	result = formatter(result)
	result = "TEXTIO: " + formatter(result)
	return result
}

