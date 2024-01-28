package main

import (
	"fmt"
)

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
	if mToSend.recipient.number == 0 {
		return false
	}
	if mToSend.sender.number == 0 {
		return false
	}
	if mToSend.recipient.name == "" {
		return false
	}
	if mToSend.sender.name == "" {
		return false
	}
	return true
}

func test(mToSend messageToSend) {
	fmt.Printf(`sending "%s" from %s (%v) to %s (%v)...`,
		mToSend.message,
		mToSend.sender.name,
		mToSend.sender.number,
		mToSend.recipient.name,
		mToSend.recipient.number,
	)
	fmt.Println()
	if canSendMessage(mToSend) {
		fmt.Println("...sent!")
	} else {
		fmt.Println("...can't send message")
	}
	fmt.Println("====================================")
}

// These are embedded structs and basically it just "inherits" the defined structs definition (helps us not retype things)
// this makes more sense when trying to access the data because to access the name of the sender you just do sender.name , not sender.user.name
// type sender struct {
// 	rateLimit int
// 	user
// }

// type user struct {
// 	name   string
// 	number int
// }

// This is a method defined on a struct, kinda like calculating properties on struct
// type authenticationInfo struct {
// 	username string
// 	password string
// }

// func (authInfo authenticationInfo) getBasicAuth() string {
// 	return "Authorization: Basic " + authInfo.username + ":" + authInfo.password
// }

func main() {
	// you can initialize structs like message := messageToSend{} and then fill it out
	// access values with dots

	// anonymous structs are possible by  and useful if only used once
	// There are more common in nested structs, however it is usually better to just define them literally
	// myCar := struct {
	// 	model string
	// }{
	// 	model: "ford",
	// }

	test(messageToSend{
		message: "you have an appointment tommorow",
		sender: user{
			name:   "Brenda Halafax",
			number: 16545550987,
		},
		recipient: user{
			name:   "Sally Sue",
			number: 19035558973,
		},
	})
	test(messageToSend{
		message: "you have an event tommorow",
		sender: user{
			number: 16545550987,
		},
		recipient: user{
			name:   "Suzie Sall",
			number: 0,
		},
	})
	test(messageToSend{
		message: "you have an party tommorow",
		sender: user{
			name:   "Njorn Halafax",
			number: 16545550987,
		},
		recipient: user{
			name:   "Sally Sue",
			number: 19035558973,
		},
	})
	test(messageToSend{
		message: "you have a birthday tommorow",
		sender: user{
			name:   "Eli Halafax",
			number: 0,
		},
		recipient: user{
			name:   "Whitaker Sue",
			number: 19035558973,
		},
	})
}
