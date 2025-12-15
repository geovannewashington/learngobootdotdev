package main

func getMessageWithRetries(primary, secondary, tertiary string) ([3]string, [3]int) {
	var messagesArr = [...]string{primary, secondary, tertiary}
	var costPerMessageArr = [3]int{0}
	
	for i := 0; i < len(messagesArr); i++ {
		if i == 0 {
			costPerMessageArr[i] = len(messagesArr[i])
			continue
		} 
		costPerMessageArr[i] = len(messagesArr[i]) + costPerMessageArr[i - 1]
	}		
	return messagesArr, costPerMessageArr
}
