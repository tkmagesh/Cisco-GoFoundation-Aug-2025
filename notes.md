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