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
  UP_TO_NUMBER:=10
	random := rand.Intn(UP_TO_NUMBER)
}
```
