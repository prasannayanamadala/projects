package operations

import "fmt"

func MulOperation() {
	
	/*Int8 — [-128 : 127]
	Int16 — [-32768 : 32767]
	Int32 — [-2147483648 : 2147483647]
	Int64 — [-9223372036854775808 : 9223372036854775807]*/
	
	var x1 int8
	x1 = 9
	var x2 int8
	x2 = 3
	MULInt8 := x1 * x2
	fmt.Println("MULInt8:", MULInt8)

	var x3 int16
	x3 = 27
	var x4 int16
	x4= 5
	MULInt16 := x3 * x4
	fmt.Println("MULInt16:", MULInt16)

	var x5 int32
	x5 = 56
	var x6 int32
	x6 = 56
	MULInt32 := x5 * x6
	fmt.Println("MULInt32:", MULInt32)

	var x7 int64
	x7 = 128
	var x8 int64
	x8 = 128
	MULInt64 := x7 * x8
	fmt.Println("MULInt64:", MULInt64)

	MULInt := int64(x1) * x8
	fmt.Println("MULInt8 & 64:", MULInt)

	/*  uint8	     0 to 255	                     
		uint16	     0 to 65,535	                  
		uint32	     0 to 4,294,967,295	               
		uint64	     0 to 18,446,744,073,709,551,615	*/

	var y1 uint8
	y1 = 254
	var y2 uint8
	y2 = 0
	MULUInt8 := y1 * y2
	fmt.Println("MULUInt8:", MULUInt8)

	var y3 uint16
	y3 = 54
	var y4 uint16
	y4 = 12
	MULUInt16 := y3 * y4
	fmt.Println("MULUInt16:", MULUInt16)

	var y5 uint32
	y5 = 64
	var y6 uint32
	y6 = 621
	MULUInt32 := y5 * y6
	fmt.Println("MULUInt32:", MULUInt32)

	var y7 uint64
	y7 = 576
	var y8 uint64
	y8 = 567
	MULUInt64 := y7 * y8
	fmt.Println("MULUInt64:", MULUInt64)

	MULUInt := uint16(y1) * y3
	fmt.Println("MULUInt8 & 16:", MULUInt)

//floating point

	var z1 float32
	z1 = -6.4
	var z2 float32
	z2 = -7.8
	MULFloat32 := z1 * z2
	fmt.Println("MULFloat32:", MULFloat32)

	var z3 float64
	z3 = -12.43
	var z4 float64
	z4 = 2.475
	MULFloat64 := z3 * z4
	fmt.Println("MULFloat64:", MULFloat64)

	MULFloat := float64(z1) * z4
	fmt.Println("MULFloat32 & 64:", MULFloat)

}