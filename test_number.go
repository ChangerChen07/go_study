// package main

// import (
// 	"fmt"
// 	"math"
// )

// func main() {
// 	var a int = 10
// 	fmt.Printf("a: %d\n", a)
// 	fmt.Printf("a: %b\n", a) //占位符%b表示2进制

// 	//8进制 以0开头
// 	var b int = 077
// 	fmt.Printf("b: %o\n", b)

// 	//16进制以0x开头
// 	var c int = 0xff
// 	fmt.Printf("c: %x\n", c)
// 	fmt.Printf("c: %X\n", c)

// 	fmt.Printf("pi: %f\n", math.Pi)
// 	fmt.Printf("pi: %.2f\n", math.Pi)

// 	var (
// 		c1 complex64
// 		c2 complex128
// 	)
// 	c1 = 1 + 2i
// 	c2 = 1 + 2i
// 	fmt.Println(c1)
// 	fmt.Println(c2)
// 	var i8 int8
// 	var i16 int16
// 	var i32 int32
// 	var i64 int64
// 	var ui8 uint8
// 	var ui16 uint16
// 	var ui32 uint32
// 	var ui64 uint64
// 	fmt.Printf("%T %dB %v~%v\n", i8, unsafe.Sizeof(i8), math.MinInt8, math.MaxInt8)
// 	fmt.Printf("%T %dB %v~%v\n", i16, unsafe.Sizeof(i16), math.MinInt16, math.MaxInt16)
// 	fmt.Printf("%T %dB %v~%v\n", i32, unsafe.Sizeof(i32), math.MinInt32, math.MaxInt32)
// 	fmt.Printf("%T %dB %v~%v\n", i64, unsafe.Sizeof(i64), math.MinInt64, math.MaxInt64)
// 	fmt.Printf("%T %dB %v~%v\n", ui8, unsafe.Sizeof(ui8), 0, math.MaxUint8)
// 	fmt.Printf("%T %dB %v~%v\n", ui16, unsafe.Sizeof(ui16), 0, math.MaxInt16)
// 	fmt.Printf("%T %dB %v~%v\n", ui32, unsafe.Sizeof(ui32), 0, math.MaxInt32)
// 	fmt.Printf("%T %dB %v~%v\n", ui64, unsafe.Sizeof(ui64), 0, math.MaxInt64)

// 	var f32 float32
// 	var f64 float64
// 	fmt.Printf("%T %dB %v~%v\n", f32, unsafe.Sizeof(f32), -math.MaxFloat32, math.MaxFloat32)
// 	fmt.Printf("%T %dB %v~%v\n", f64, unsafe.Sizeof(f64), -math.MaxFloat64, math.MaxFloat64)

// 	var ui uint
// 	// ui = uint(math.MaxUint64)
// 	ui = math.MaxUint64
// 	fmt.Printf("%T %dB %v~%v\n", ui, unsafe.Sizeof(ui), 0, math.MaxInt64)

// 	var imin, imax int
// 	imin = math.MinInt64
// 	imax = math.MaxInt64
// 	fmt.Printf("%T %dB %v~%v\n", imax, unsafe.Sizeof(imax), imin, imax)
// }
