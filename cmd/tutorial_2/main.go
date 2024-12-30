package main

import (
	"fmt"
	"math"
)

func main() {
	var maxInt = math.MaxInt
	var minInt = math.MinInt
	fmt.Printf("Maximum value for int: %d\n", maxInt)
	fmt.Printf("Minimum value for int: %d\n", minInt)

	var maxInt8 = math.MaxInt8
	var minInt8 = math.MinInt8
	fmt.Printf("Maximum value for int8: %d\n", maxInt8)
	fmt.Printf("Minimum value for int8: %d\n", minInt8)

	var maxInt16 = math.MaxInt16
	var minInt16 = math.MinInt16
	fmt.Printf("Maximum value for int16: %d\n", maxInt16)
	fmt.Printf("Minimum value for int16: %d\n", minInt16)

	var maxInt32 = math.MaxInt32
	var minInt32 = math.MinInt32
	fmt.Printf("Maximum value for int32: %d\n", maxInt32)
	fmt.Printf("Minimum value for int32: %d\n", minInt32)

	var maxInt64 = math.MaxInt64
	var minInt64 = math.MinInt64
	fmt.Printf("Maximum value for int64: %d\n", maxInt64)
	fmt.Printf("Minimum value for int64: %d\n", minInt64)

	var maxUInt uint = math.MaxUint
	fmt.Printf("Maximum value for uint: %d\n", maxUInt)

	var maxUInt8 uint8 = math.MaxUint8
	fmt.Printf("Maximum value for uint8: %d\n", maxUInt8)

	var maxUInt16 uint16 = math.MaxUint16
	fmt.Printf("Maximum value for int16: %d\n", maxUInt16)

	var maxUInt32 uint32 = math.MaxUint32
	fmt.Printf("Maximum value for int32: %d\n", maxUInt32)

	var maxUInt64 uint64 = math.MaxUint64
	fmt.Printf("Maximum value for int64: %d\n", maxUInt64)

	var maxFloat32 = math.MaxFloat32
	var maxFloat64 = math.MaxFloat64
	fmt.Printf("Maximum value for float32: %d\n", maxFloat32)
	fmt.Printf("maximum value for float64: %d\n", maxFloat64)
}
