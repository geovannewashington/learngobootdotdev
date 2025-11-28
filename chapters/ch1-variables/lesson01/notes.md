# Notes on the First Lesson:

To build a Go program, use go build.
If the source file is named main.go, you can simply run:

```go
go build
```

If the file has a different name, specify it explicitly:

```go
go build <file_name>
```

By default, this command produces an executable with the same name as the directory (or the file,
when specified), without an extension.

Like gcc or clang, Go also supports the -o flag to choose a custom output name:

```go
go build -o output_name
```

## Package Declaration

Every Go source file begins with a package declaration.

- `package main` indicates that the file is part of an executable program.
- Any package named main must contain a main() function, which serves as the program’s entry point.
