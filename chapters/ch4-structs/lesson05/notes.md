# Embedded Structs

Embedded Structs provide a kind of data-only inheritance that can be useful at times.

Embedded Structs are a way to elevate and share fields between structs definitions.

Example:

```go
type car struct {
  brand string
  model string
}

type truck struct {
  // "car" is embedded, so the definition of a
  // "truck" now also additionally contains all
  // of the fields of the car struct
  car
  bedSize int
}
```

Unlilke nestes structs, am embedded struct's fields are accessable at the top level like normal fields.

Example:

```go
lanesTruck := truck{
  bedSize: 10,
  car: car{
    brand: "Toyota",
    model: "Tundra",
  },
}

fmt.Println(lanesTruck.brand) // Toyota
fmt.Println(lanesTruck.model) // Tundra
```

As you can see both `brand` and `model` are accessible from the top level.
So intead of having to write: `lanesTruck.car.brand` we just write: `lanesTruck.brand`
