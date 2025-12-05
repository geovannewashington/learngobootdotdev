# The Defer Keyword

The `defer` keyword is a faily unique function of Go. It allows a funtion to be executed automatically
just before its enclosing function returns.

## Real World Usage

Deferred functions are typically used to clean up resources that are no longer being used.
Often to close database connections, file handlers and the like.

For example:

```go
func GetUsername(dstName, srcName string) (username string, err error) {
	// Open a connection to a database
	conn, _ := db.Open(srcName)

	// Close the connection *anywhere* the GetUsername function returns
	defer conn.Close()

	username, err = db.FetchUser()
	if err != nil {
		// The defer statement is auto-executed if we return here
		return "", err
	}

	// The defer statement is auto-executed if we return here
	return username, nil
}
```
