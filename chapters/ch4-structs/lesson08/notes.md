## Empty Struct

Empty structs are used in GO as a unary value.

```Go
// anonymous empty struct type
empty := struct{}{}

// named empty struct type
type emptyStruct struct{}
empty := emptyStruct{}
```

The cool thing about empty structs is that they're the smallest posible type in Go: they take up zero bytes of
memory.
