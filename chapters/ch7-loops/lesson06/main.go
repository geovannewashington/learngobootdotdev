package main

// import "fmt"

func countConnections(groupSize int) int {
	connectionCounter := 0

	for i := 1; i < groupSize; i++ {
		connectionCounter += i	
	}
	
	return connectionCounter
}

// func main() {
// 	res := countConnections(5) 
// 	fmt.Println(res)
// }
