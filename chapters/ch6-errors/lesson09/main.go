package main

import (
	"errors"
	"fmt"
)

func validateStatus(status string) error {
	maxStatusLen := 140

	if status == "" {
		return errors.New("status cannot be empty")
	} 

	if len(status) > maxStatusLen {
		s := fmt.Sprintf("status exceeds %v characters", maxStatusLen);
		return errors.New(s)
	}
	
	return nil
}
