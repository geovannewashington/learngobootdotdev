## Computed Constants

Constants can be computed as long as the computation can happen at compile time.

example:

```go
const firstName = "Lane"
const lastName = "Wagner"
const fullName = firstName + " " + lastName
```

That said, you cannot declare a constant that can only be computed at run-time like you can in JavaScript. This breaks:

```js
// the current time can only be known when the program is running
const currentTime = time.Now();
```
