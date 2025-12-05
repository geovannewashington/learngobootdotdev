package main

import (
	"fmt"
)

// defines a struct
type Computer struct {
	cpu      string
	ramQuant string
	gpu      string
}

func main() {
	var firstComputer Computer
	firstComputer.cpu = "i5-9400f"
	firstComputer.ramQuant = "8GB"
	firstComputer.gpu = "nVidia GTX 1660"

	secondComputer := Computer {
		cpu : "AMD Ryzen 3600x",
		ramQuant : "8GB",
		gpu : "nVidia GTX 4090",
	}

	// print first computer info
	fmt.Printf("cpu: %s, ram mem quantity: %s, gpu: %s\n", firstComputer.cpu, firstComputer.ramQuant, 
		firstComputer.gpu)
	// print second computer info
	fmt.Printf("cpu: %s, ram mem quantity: %s, gpu: %s\n", secondComputer.cpu, secondComputer.ramQuant, 
		secondComputer.gpu)
} 
