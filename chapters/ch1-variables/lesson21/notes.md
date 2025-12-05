## Runes and String Encoding

### What is a rune?

It's an alias for an int32 type that is a single unicode point, which is the numerical value for
a character.

Note that 32 bits is large enough to hold any unicode code point

A rune is a more accurate way to represent a single character than a byte, espicifically for multi-byte
characters, e.g.: emojis and chinese characters...

A string is a collection of runes

The concept of a character in Golang is a rune.
