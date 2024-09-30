# Basic types, values and pointers


| Type  | Examples |
| ------------- | ------------- |
| int  | 20, -20. Values can also be expressed in hex (0x14), octal (0o24), and binary notation (0b0010100)  |
| uint | There are no uint literals. All literal whole numbers are treated as int values  |
| byte | There are no byte literals. Bytes are typically expressed as integer literals (such as 101) or run literals ('e') since the byte type is an alias for the uint8 type  |
| float64 | 20.2, -20.2, 1.2e10, 1.2e-10. Values can also be expressed in hex notation (0x2p10), although the exponent is expressed in decimal digits  |
| bool | true, false  |
| string | "Hello". Character sequences escaped with a backslash are interpreted if the value is enclosed in double quotes ("Hello\n"). Escape sequences are not interpreted if the value is enclosed in backquotes (`Hello\n`)  |
| rune | 'A', '\n', '\u00A5', '¥'. Characters, glyphs, and escape sequences are enclosed in single quotes (the ' character)  |

## Using Constants

Constants are names for specific values, which allows them to be used repeatedly and consistently. There are two ways to define constants in Go: typed constants and untyped constants

**typed**

```go
func main() {
    const price float32 = 275.00
    const tax float32 = 27.50
    fmt.Println(price + tax)
}
```
| keyword  | name | type  | value |
| ------------- | ------------- | ------------- | ------------- |
| const | price | float32 | 275.00 |

**untyped**

```go
    ...
    const quantity = 2
    fmt.Println("Total:", quantity * (price + tax))
    ...
```
| keyword  | name | value |
| ------------- | ------------- | ------------- |
| const | quantity | 2 |

