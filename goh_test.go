package goh

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestReplace(t *testing.T) {
	values := []int{1, 2, 3, 4, 5}

	expected := []int{1, 2, 3, 4, 6}
	actual := Replace(&values, 5, 6)

	assert.Equal(t, expected, actual)
}

func TestFilterStrings(t *testing.T) {
	values := []string{"Foo bar", "John Doe", "Tom"}

	expected := []string{"Foo bar", "John Doe"}
	actual := FilterStrings(&values, func(v string) bool {
		return len(v) > 3
	})

	assert.Equal(t, expected, actual)
}

func TestFilterInt(t *testing.T) {
	values := []int{1, 12, 3, 5, 4, 2, 100, 111, 120}

	expected := []int{111, 120}
	actual := FilterInt(&values, func(v int) bool {
		return v > 100
	})

	assert.Equal(t, expected, actual)
}

func TestIncludeString(t *testing.T) {
	t.Run("include value", func(t *testing.T) {
		values := []string{"John Doe", "Foo bar", "Bar foo"}
		actual := IncludeString(values, "John Doe")

		assert.Equal(t, true, actual)
	})

	t.Run("doesn't include value", func(t *testing.T) {
		values := []string{"John Doe", "Foo bar", "Bar foo"}
		actual := IncludeString(values, "Bar")

		assert.Equal(t, false, actual)
	})
}

func TestIncludeInt(t *testing.T) {
	t.Run("include value", func(t *testing.T) {
		values := []int{12, 1001, 12}
		actual := IncludeInt(values, 12)

		assert.Equal(t, true, actual)
	})

	t.Run("doesn't include value", func(t *testing.T) {
		values := []int{12, 1001, 12}
		actual := IncludeInt(values, 48)

		assert.Equal(t, false, actual)
	})
}

func TestBetween(t *testing.T) {
	t.Run("is between", func(t *testing.T) {
		values := []int{1, 2, 3, 5, 1, 10}

		actual := Between(values, 3, 2, 5)

		assert.Equal(t, true, actual)
	})

	t.Run("is't between", func(t *testing.T) {
		values := []int{1, 2, 3, 5, 1, 10}

		actual := Between(values, 3, 10, 5)

		assert.Equal(t, false, actual)
	})
}

func TestMapString(t *testing.T) {
	values := []string{"Foo bar", "John Doe"}

	expected := []string{"Foo bar mapped", "John Doe mapped"}

	actual := MapString(&values, func(v string) string {
		return v + " mapped"
	})

	assert.Equal(t, expected, actual)
}

func TestMapInt(t *testing.T) {
	values := []int{1, 2, 3, 4}

	expected := []int{2, 3, 4, 5}

	actual := MapInt(&values, func(v int) int {
		return v + 1
	})

	assert.Equal(t, expected, actual)
}

func TestStringSliceEqual(t *testing.T) {
	t.Run("slices are equal", func(t *testing.T) {
		firstValue := []string{"John Doe", "Foo Bar"}
		secondValue := []string{"John Doe", "Foo Bar"}

		assert.Equal(t, StringSliceEqual(firstValue, secondValue), true)
	})

	t.Run("slices not equal", func(t *testing.T) {
		firstValue := []string{"John Doe", "Foo Bar"}
		secondValue := []string{"John Doe", "Foo Baz"}

		assert.Equal(t, StringSliceEqual(firstValue, secondValue), false)
	})
}

func TestIntSliceEqual(t *testing.T) {
	t.Run("slices are equal", func(t *testing.T) {
		firstValue := []int{1, 122, 130}
		secondValue := []int{1, 122, 130}

		assert.Equal(t, IntSliceEqual(firstValue, secondValue), true)
	})

	t.Run("slices not equal", func(t *testing.T) {
		firstValue := []int{0, 1, 1}
		secondValue := []int{1, 0, 0}

		assert.Equal(t, IntSliceEqual(firstValue, secondValue), false)
	})
}
