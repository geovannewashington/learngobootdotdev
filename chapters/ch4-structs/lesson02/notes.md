# Nested Structs in Go

Structs can be nested to represent more complex entities:

Example:

```go
type car struct {
  brand string
  model string
  doors int
  mileage int
  frontWheel wheel
  backWheel wheel
}

type wheel struct {
  radius int
  material string
}
```

Note that we can also assign a structure using the `walrus` operator.

`myCar := car{}`
