## Explicit Returns

Even though a function has named return values, we can still explicitly return values if we want to.

e.g:

```go
func getCoords() (x, y int) {
	return x, y // this is explicit
}
```

Otherwise, if we want to return the values defined in the function signature we can just use a naked
`return` (blank return)

e.g:

```go
func getCoords() (x, y int) {
	return // implicityly returns x and y
}
```
