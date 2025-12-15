package main

func bulkSend(numMessages int) float64 {
	messageCost := 1.0
	totalCost := 0.0
	
	for i := 0.0; i < float64(numMessages); i++ {
		totalCost += messageCost + i / 100.0
	}

	return totalCost
}
