package main

type messageToSend struct {
	message   string
	sender    user
	recipient user
}

type user struct {
	name   string
	number int
}

func canSendMessage(mToSend messageToSend) bool {
	fieldsToCheck := [...]string{"sender", "recipient"}
	// loop twice 
	// mToSend.[fieldsToCheck[i]].name == "" || mToSend.[fieldsToCheck[i]].number == 0

	// if mToSend.sender.name == "" || mToSend.sender.number == 0 {
	// 	return false
	// }

	// check recipient
	// if mToSend.recipient.name == "" || mToSend.recipient.number == 0 {
	// 	return false
	// }
	for i := 0; i < len(fieldsToCheck); i++ {
		// if mToSend
	}
	return true
}
