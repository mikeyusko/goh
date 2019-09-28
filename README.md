# goh

GO helpers - is a tiny-tidy library aims to avoid some default manipulations which you are doing every day.


## Examples

```go
package main

import (
	"github.com/mikeyusko/goh"
)

func main() {
	// Returns a new slice
	// [1 10 3]
	goh.Replace(&[]int{1, 2, 3}, 2, 10)

	// Returns a new slice, and only with values
	// which passed condition
	//[Foo bar]
	goh.FilterStrings(&[]string{"Foo bar", "Doe"}, func(v string) bool {
		return len(v) > 4
	})

	// The same as for FilterString but for integers
	// [3, 4]
	goh.FilterInt(&[]int{1, 2, 3, 4}, func(v int) bool {
		return v >= 3
	})

	// Checks if a string value is present in a slice
	// true
	goh.IncludeString([]string{"John Doe", "Foo bar"}, "Foo bar")

	// The same as for IncludeString but only for integers
	// true
	goh.IncludeInt([]int{1, 2, 3}, 1)

	// Checks if a value is present in a slice
	// between two specific values
	// true
	goh.Between([]int{1, 5, 3, 4}, 3, 5, 4)

	// Changed all values in a slice based on a specific function
	// returns a new slice with modified values
	// [Foo bar New! John Doe New!]
	goh.MapString(&[]string{"Foo bar", "John Doe"}, func(v string) string {
		return v + " New!"
	})

	// The same as for MapString but only for integers
	// returns a new slice with modified values
	// [2 3 4]
	goh.MapInt(&[]int{1, 2, 3}, func(v int) int {
		return v + 1
	})

	// Checks if a two slices are equal
	// true
	goh.StringSliceEqual([]string{"John Doe"}, []string{"John Doe"})

	// Checks if a two slices are equal
	// true
	goh.IntSliceEqual([]int{1, 2, 3}, []int{1, 2, 3})
}
```