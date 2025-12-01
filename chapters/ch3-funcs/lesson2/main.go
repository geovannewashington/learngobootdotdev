package main

import "strings"

func getMonthlyPrice(tier string) int {
	tier = strings.ToLower(tier)

	switch tier {
		case "basic": 
			return 10000
		case "premium": 
			return 15000
		case "enterprise": 
			return 50000
		default:
			return 0
	}
}
