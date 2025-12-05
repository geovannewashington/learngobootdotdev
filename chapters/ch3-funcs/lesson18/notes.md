# Block Scope

Go is block-scoped, like C. New blocks are created by functions, loops, if-else statements, etc...

Example:

```go
package main

// scoped to the entire "main" package (basically global)
var age = 19

func sendEmail() {
    // scoped to the "sendEmail" function
    name := "Jon Snow"

    for i := 0; i < 5; i++ {
        // scoped to the "for" body
        email := "snow@winterfell.net"
    }
}
```

It's a bit unusual, but occasionally you'll see a plain old explicit block. It exists for no other
reason than to create a new scope.

```go
package main

import fmt

func main() {
    {
        age := 19
        // this is okay
        fmt.Println(age)
    }

    // this is not okay
    // the age variable is out of scope
    fmt.Println(age)
}
```
