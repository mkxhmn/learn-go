- [Array vs slice](#array-vs-slice)
- [Type conversion](#type-conversion)
- [Generate between range](#generate-between-range)
- [Testing in Go](#testing-in-go)

# Good notes

## Array vs slice

> both must be defined from the same data type

- _Array_ - fixed length list of things
- _Slice_ - array that can grow or shrink

## Type conversion

```go
[]byte("Hi There")
```

| position          | descriptions                  |
| ----------------- | ----------------------------- |
| _left hand side_  | specify the type that we want |
| _right hand side_ | specify the argument          |

## Generate between range

> NOTE: Intn is pseudorandom, it may not be best fit for your case to generate random number

```go
import (
"math/rand"
)

func main(){
UP_TO_NUMBER := 10
random := rand.Intn(UP_TO_NUMBER)
}
```

## Testing in Go

> please note that function name is test is **PascalCase**

To make a test, create a new file ending in `_test.go`
To run all tests in a package, run the command

```go
go test
```

## Pointer

> good things to note here.
>
> - Turn `address` into _values_ with ***address**
> - Turn `value` into address with **&value**

### & value

other way of saying to the compiler, give me access to this memory address of the value this variable is pointing at

### \* point

give me the value this memory address is pointing at

### Underlying data structure

| value types | reference type |
|-------------|----------------|
| int         | slices         |
| float       | maps           |
| string      | channels       |
| bool        | pointers       |
| structs     | functions      |

#### values types

use pointers to change these things in a function

#### references types

don't worry about pointers with these

