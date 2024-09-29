# Using Constants

Constants are names for specific values, which allows them to be used repeatedly and consistently. There are two ways to define constants in Go: typed constants and untyped constants

**typed**

```go
func main() {
    const price float32 = 275.00
    const tax float32 = 27.50
    fmt.Println(price + tax)
}
```

**untyped**

```go
    ...
    const quantity = 2
    fmt.Println("Total:", quantity * (price + tax))
    ...
```
