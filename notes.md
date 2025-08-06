# Go Foundation

## Magesh Kuppan
- tkmagesh77@gmail.com
- https://www.linkedin.com/in/tkmagesh

## Schedule
| What | When |
|-----|-----|
| Commence | 9:00 AM |
| Tea Break | 10:30 AM (20 mins) |
| Lunch Break | 12:30 PM (1 hour) |
| Tea Break | 3:00 PM (20 mins) |
| Wind up   | 5:00 PM |

## Software Prerequisites
- Go tools (https://go.dev/dl)
    - Verification
```shell
# From the terminal
go version
```
- Visual Studio Code (or any other editor)

## Repository
- https://github.com/tkmagesh/cisco-gofoundation-aug-2025

## Why Go?
- Simplicity
    - ONLY 25 keywords
        - package, import, var, const, func, if else, switch case, for, type, interface, struct, return, select case, range, goto, continue, break, go

    - No Access modifiers (private, public, protected etc)
    - No reference types (everything is a value)
    - No pointer arithmatic
    - No classes (only structs)
    - No inheritance (only composition)
    - NO exceptions (only errors that are returned from a function)
    - No try-catch-finally
    - No implicit type conversion
- Performance
    - comparable with C++
    - Close to the hardware
        - compiled to machine code
        - Support for cross compilation in the tools
- Concurrency
    - Cheaper concurrency construct - goroutines
        - ONLY 2kb of memory (compared to 2MB for OS Threads)
    - Concurrency constructs are supported through the language
        - `go` keyword
        - `range` keyword
        - `chan` datatype
        - `<-` operator
        - `select case` construct
    - Standard Library packages
        - `sync` package
        - `sync/atomic` package

## Hello, World!
### Compilation
```shell
go build [filename.go]
go build -o [binary_name] [filename.go]
```

### Compilation & Execuation
```shell
go run [filename.go]
```

### Cross Compilation
#### To get the env variables
```shell
go env
# OR
go env [var_1] [var_2] ...
```

#### To set the env variable value
```shell
go env -w [var_1 = value_1] [var_2=value_2] ...
```

#### Env variables for cross compilation
- GOOS
- GOARCH

#### To get the list of platforms supported
```shell
go tool dist list
```

#### To cross compile
```shell
GOOS=[os] GOARCH=[arch] go build [filename.go]

# ex:
GOOS=windows GOARCH=amd64 go build 01-hello-world.go # to create a build for windows
# OR
GOOS=darwin GOARCH=arm64 go build 01-hello-world.go # to create a build for mac os
```

## Standard Library
- https://pkg.go.dev/std

## Data Types
- string
- bool
- integers
    - int8
    - int16
    - int32
    - int64
    - int
- unsigned integers
    - uint8
    - uint16
    - uint32
    - uint64
    - uint
- floating points
    - float32
    - float64
- complex
    - complex64 ( real[float32] + imaginary[float32] )
    - complex128 ( real[float64] + imaginary[float64] )
- alias
    - byte (alias for unsigned int)
    - rune (alias for unicode code point)

### Zero values
| Data Type | Zero value |
------------ | ------------- |
|int family     | 0 |
|uint family    | 0 |
|complex family | (0+0i) |
|string         | "" (empty string) |
|bool           | false |
|byte           | 0 |
|interface      | nil |
|pointer        | nil |
|function       | nil |
|struct         | struct instance |


## Variables
### Function Scope
- Can use ":="
- CANNOT have unused variables

### Package Scope
- Cannot use ":="
- CAN have unused variables

## iota
- A group for constants whose values can be auto-generated

## Programming Constructs
- if else
- switch case
    - fallthrough
- for