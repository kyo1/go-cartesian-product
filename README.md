# go-cartesian-product

This package provides a method to generate cartesian products (like python's `itertools.product`) out of given slice.

## Installation

```sh
go get github.com/kyo1/go-cartesian-product
```

## Usage


### `func All(ctx context.Context, set []interface{}) chan []interface{}`

`All` function generates elements of all n-fold Cartesian product.

```go
package main

import (
    "fmt"
    "github.com/kyo1/go-cartesian-product"
)

func main() {
    chars := []interface{}{"a", "b", "c"}

    // Generate a Cartesian product of all length
    cnt := 0
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    for s := range cartesian.All(ctx, chars) {
        fmt.Println(s)
        cnt++

        // Conditions for stopping the generator
        if cnt == 10 {
            cancel()
            continue
        }
    }

    // Output

    // [a]
    // [b]
    // [c]
    // [a b]
    // [b b]
    // [c b]
    // [a c]
    // [b c]
    // [c c]
    // [a a b]
}
```

### `func Product(ctx context.Context, set []interface{}, repeat int) chan []interface{}`

`Product` function generates elements of n-fold Cartesian product.

```go
package main

import (
    "fmt"
    "github.com/kyo1/go-cartesian-product"
)

func main() {
    chars := []interface{}{"a", "b", "c"}

    // Generates n-fold Cartesian product
    ctx, cancel := context.WithCancel(context.Background())
    defer cancel()

    for s := range cartesian.Product(ctx, chars, 2) {
        fmt.Println(s)

        // The condition for terminating the generator is not required
        // if condition {
        //  cancel()
        //  continue
        // }
    }

    // Output
    // [a a]
    // [b a]
    // [c a]
    // [a b]
    // [b b]
    // [c b]
    // [a c]
    // [b c]
    // [c c]
}
```
