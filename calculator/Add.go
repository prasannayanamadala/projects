package operations

import "fmt"

func AddOperation() {
	
	/*Int8 — [-128 : 127]
	Int16 — [-32768 : 32767]
	Int32 — [-2147483648 : 2147483647]
	Int64 — [-9223372036854775808 : 9223372036854775807]*/
	
	var x1 int8
	x1 = 56
	var x2 int8
	x2 = 55
	ADDInt8 := x1 + x2
	fmt.Println("ADDInt8:", ADDInt8)

	var x3 int16
	x3 = -310
	var x4 int16
	x4= -280
	ADDInt16 := x3 + x4
	fmt.Println("ADDInt16:", ADDInt16)

	var x5 int32
	x5 = 567
	var x6 int32
	x6 = 56768
	ADDInt32 := x5 + x6
	fmt.Println("ADDInt32:", ADDInt32)

	var x7 int64
	x7 = 345466578
	var x8 int64
	x8 = 67868
	ADDInt64 := x7 + x8
	fmt.Println("ADDInt64:", ADDInt64)

	ADDInt := int64(x1) + x8
	fmt.Println("ADDInt8 & 64:", ADDInt)

	/*  uint8	     0 to 255	                     
		uint16	     0 to 65,535	                  
		uint32	     0 to 4,294,967,295	               
		uint64	     0 to 18,446,744,073,709,551,615	*/

	var y1 uint8
	y1 = 254
	var y2 uint8
	y2 = 0
	ADDUInt8 := y1 + y2
	fmt.Println("ADDUInt8:", ADDUInt8)

	var y3 uint16
	y3 = 768
	var y4 uint16
	y4 = 6
	ADDUInt16 := y3 + y4
	fmt.Println("ADDUInt16:", ADDUInt16)

	var y5 uint32
	y5 = 657
	var y6 uint32
	y6 = 6786
	ADDUInt32 := y5 + y6
	fmt.Println("ADDUInt32:", ADDUInt32)

	var y7 uint64
	y7 = 576
	var y8 uint64
	y8 = 56
	ADDUInt64 := y7 + y8
	fmt.Println("ADDUInt64:", ADDUInt64)

	ADDUInt := uint16(y1) + y3
	fmt.Println("ADDUInt8 & 16:", ADDUInt)

//floating point

	var z1 float32
	z1 = -666.433
	var z2 float32
	z2 = -577.8898
	ADDFloat32 := z1 + z2
	fmt.Println("ADDFloat32:", ADDFloat32)

	var z3 float64
	z3 = -666.43
	var z4 float64
	z4 = -577.889
	ADDFloat64 := z3 + z4
	fmt.Println("ADDFloat64:", ADDFloat64)

	ADDFloat := float64(z1) + z4
	fmt.Println("ADDFloat32 & 64:", ADDFloat)
}
