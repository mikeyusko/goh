package goh

import "reflect"

// IncludeStrings replace a value with new value
// and returns a new slice without modifying original.
func Replace(values *[]int, old, new int) []int {
	result := make([]int, len(*values))

	for i, v := range *values {
		if v == old {
			result[i] = new
		} else {
			result[i] = v
		}
	}
	return result
}

// FilterStrings gets a slice of strings
// and returns a new slice with values which passed condition check.
func FilterStrings(values *[]string, condition func(v string) bool) []string {
	var result []string

	for _, v := range *values {
		if condition(v) {
			result = append(result, v)
		}
	}
	return result
}

// FilterStrings gets a slice of integers
// and returns a new one with values which passed condition check.
func FilterInt(values *[]int, condition func(v int) bool) []int {
	var result []int

	for _, v := range *values {
		if condition(v) {
			result = append(result, v)
		}
	}
	return result
}

// IncludeStrings checks if a value is present in a slice.
func IncludeString(values []string, value string) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}

// IncludeInt checks if a value is present in a slice.
func IncludeInt(values []int, value int) bool {
	for _, v := range values {
		if v == value {
			return true
		}
	}
	return false
}

// Between checks if a value is between two values in a slice.
func Between(values []int, value, beforeValue, afterValue int) bool {
	for i, v := range values {
		if v == value && values[i-1] == beforeValue && values[i+1] == afterValue {
			return true
		}
	}
	return false
}

// MapString creates a new value based on function argument
// and returns a new slice.
func MapString(values *[]string, f func(v string) string) []string {
	result := make([]string, len(*values))

	for i, v := range *values {
		result[i] = f(v)
	}
	return result
}

// MapString creates a new value based on function argument
// and returns a new slice.
func MapInt(values *[]int, f func(v int) int) []int {
	result := make([]int, len(*values))

	for i, v := range *values {
		result[i] = f(v)
	}
	return result
}

// StringSliceEqual checks if two slices of strings are equal
func StringSliceEqual(f, s []string) bool {
	return reflect.DeepEqual(f, s)
}

// IntSliceEqual checks if two slices of integers are equal
func IntSliceEqual(f, s []int) bool {
	return reflect.DeepEqual(f, s)
}