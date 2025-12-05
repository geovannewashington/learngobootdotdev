## Ignoring Return Value

In Golang we can ignore one of the return values of a function.

Example:

```go
package main

import (
	"fmt"
	"errors"
)

func greet(user string) (string, error) {
	const errMsg string = "username is a required field"

	if user == "" {
		return "", errors.New(errMsg)
	}

	return "Welcome, " + user, nil
}

func main() {
	str, _ := greet("John Doe")

	if err != nil {
		fmt.Printf("greet: %s\n", err)
		return
	}

	fmt.Printf("%s\n", str);
}

```

This underscore is known as: the blank identifier
This is useful when you know you want a specific value of a function and don't want the compiler to
cry about it.
