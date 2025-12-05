# Functions as values

Functions in go are first-class citizens, which just mean they can be used as values, that is,
getting assigned to variables, passed to functions, return from functions, etc...
Example:

```Go
func aggregate(a, b, c int, arithmetic func(int, int) int) int {
  firstResult := arithmetic(a, b)
  secondResult := arithmetic(firstResult, c)
  return secondResult
}
```

The function `aggregate` will accept any function that takes two integers and returns an integer, that is,
a function that matches the signature of what we called `arithmetic`.
