package main

import "fmt"

// SumInts adds together the values of m.
func sumInts(m map[string]int64) int64 {
    var sum int64
    for _, value := range m {
        sum += value
    }
    return sum
}

// SumFloats adds together the values of m.
func SumFloats(m map[string]float64) float64 {
    var sum float64
    for _, value := range m {
        sum += value
    }
    return sum
}


// SumIntsOrFloats sums the values of map m. It supports both int64 and float64
// as types for map values.
func SumIntsOrFloats[MapKeyType comparable, MapValueType int64 | float64](m map[MapKeyType]MapValueType) MapValueType {
    var sum MapValueType
    for _, value := range m {
        sum += value
    }
    return sum
}

func main() {
    // Initialize a map for the integer values
    ints := map[string]int64 {
        "first": 34,
        "second": 12,
    }

    // Initialize a map for the float values
    floats := map[string]float64 {
        "first":  35.98,
        "second": 26.99,
    }

    fmt.Printf("Non-Generic Sums: %v and %v\n", SumIntsOrFloats(ints), SumIntsOrFloats(floats))
}
