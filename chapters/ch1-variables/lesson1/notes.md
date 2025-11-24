# Notes on first lesson:

To build a go program we need to use `go build`, if the file is named `main.go` we don't necessarly
need to specify it, but if it's named anything else we do, so it would be: `go build <file_name>`
by default this generates an executable binary with the same name (just without the extension)
Similar to gcc/clang we can use the `-o` option to change the default output name.
