package main

type Membership struct {
	Type string;
	MessageCharLimit int;
}

type User struct {
	// embed 'Membership'
	Membership
	Name string
}

func newUser(name string, membershipType string) User {
	user := User{
		Membership: Membership{
			Type: membershipType,
		},
		Name: name,
	}

	// Check wheter user is premium or not (based on 'membershipType')
	if membershipType == "premium" {
		user.MessageCharLimit = 1000
	} else {
		user.MessageCharLimit = 100
	}
	
	return user;
}
