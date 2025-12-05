# Memory Layout

The order of fields in a struct can have a big impact on memory usage. This is the same struct as
above, but poorly designed:

This is closely related to the tick of a cpu, for performance reasons it processes a specific number
of bytes per miliseconds, when a struct is poorly designed, go wil, by default, add the so called
`padding` to solve this, but this as you can imagine can lead to a lot of wasted memory.

You can use the `reflect package` to debug the memory layout of a structure:

example:

```go
typ := reflect.TypeOf(stats{})
fmt.Printf("Struct is %d bytes\n", typ.Size())
```
