# go-cartesian-product

This package provides a method to generate cartesian products (like python's `itertools.product`) out of given string array.

## Installation

```sh
go get github.com/kyo1/go-cartesian-product
```

## Usage

```go
package main

import (
	"fmt"
	"github.com/kyo1/go-cartesian-product"
	"strings"
)

func main() {
	chars := []string{"a", "b", "c"}

	// Generate a Cartesian product of all length
	all := cartesian.All(chars)
	for i := 0; i < 10; i++ {
		fmt.Println(strings.Join(all(), ""))
	}

	fmt.Println()

	// Generate a Cartesian product of length n (now n = 2)
	for _, product := range cartesian.Product(chars, 2) {
		fmt.Println(strings.Join(product, ""))
	}

	// Output
	// a
	// b
	// c
	// aa
	// ba
	// ca
	// ab
	// bb
	// cb
	// ac
	//
	// aa
	// ba
	// ca
	// ab
	// bb
	// cb
	// ac
	// bc
	// cc
}
```
