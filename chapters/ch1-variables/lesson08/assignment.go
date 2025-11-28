package main

import "fmt"

func main() {
	account_age_float := 2.6
	account_years_amount := int64(account_age_float)
	tmp := int((account_age_float) * 10) % 10;
	account_months_amount := int32((tmp * 12) / 10)
	
	// fmt.Println("Your account has existed for", account_age_float, "years")
	fmt.Printf("Your account has existed for %d years and %d months.\n", account_years_amount, account_months_amount)
}
