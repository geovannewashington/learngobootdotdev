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

## Anonymous Structs

To create an anonymous struct, just instantiate the instance immediately using a second pair of
brackets after declaring the type:

Example:

```go
myCar := struct {
  brand string
  model string
} {
  brand: "Toyota",
  model: "Camry",
}
```

You can even nest anonymous structs as fields within other structs:

Example:

```go
type car struct {
  brand string
  model string
  doors int
  mileage int
  // wheel is a field containing an anonymous struct
  wheel struct {
    radius int
    material string
  }
}

var myCar = car{
  brand:   "Rezvani",
  model:   "Vengeance",
  doors:   4,
  mileage: 35000,
  wheel: struct {
    radius   int
    material string
  }{
    radius:   35,
    material: "alloy",
  },
}
```

Generally we prefer using named structures, it's easier to read an reuse.
And anonymous structs are more used when we know we're only be using that once.
