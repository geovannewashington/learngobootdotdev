package main

func maxMessages(thresh int) int {
	currentTotalPrice, messageCounter := 0, 0
	messageCost := 100	
		
	for i := 0; ; i++ {
		if currentTotalPrice > thresh {
			break
		} 
		currentTotalPrice += messageCost + i
		messageCounter = i;
	}
	return messageCounter
}

