package operations

import "fmt"

func SubOperation() {
	
	/*Int8 — [-128 : 127]
	Int16 — [-32768 : 32767]
	Int32 — [-2147483648 : 2147483647]
	Int64 — [-9223372036854775808 : 9223372036854775807]*/
	
	var x1 int8
	x1 = -125
	var x2 int8
	x2 = -55
	SUBInt8 := x1 - x2
	fmt.Println("SUBInt8:", SUBInt8)

	var x3 int16
	x3 = -31000
	var x4 int16
	x4= -28000
	SUBInt16 := x3 - x4
	fmt.Println("SUBInt16:", SUBInt16)

	var x5 int32
	x5 = 56778899
	var x6 int32
	x6 = 567687698
	SUBInt32 := x5 - x6
	fmt.Println("SUBInt32:", SUBInt32)

	var x7 int64
	x7 = 345466578789999999
	var x8 int64
	x8 = 678687908098080675
	SUBInt64 := x7 - x8
	fmt.Println("SUBInt64:", SUBInt64)

	SUBInt := int64(x1) - x8
	fmt.Println("SUBInt8 & 64:", SUBInt)

	/*  uint8	     0 to 255	                     
		uint16	     0 to 65,535	                  
		uint32	     0 to 4,294,967,295	               
		uint64	     0 to 18,446,744,073,709,551,615	*/

	var y1 uint8
	y1 = 254
	var y2 uint8
	y2 = 0
	SUBUInt8 := y1 - y2
	fmt.Println("SUBInt8:", SUBUInt8)

	var y3 uint16
	y3 = 7689
	var y4 uint16
	y4 = 64232
	SUBUInt16 := y3 - y4
	fmt.Println("SUBUInt16:", SUBUInt16)

	var y5 uint32
	y5 = 657868698
	var y6 uint32
	y6 = 678679999
	SUBUInt32 := y5 - y6
	fmt.Println("SUBUInt32:", SUBUInt32)

	var y7 uint64
	y7 = 576588988654555
	var y8 uint64
	y8 = 567587688333330998
	SUBUInt64 := y7 - y8
	fmt.Println("SUBUInt64:", SUBUInt64)

	SUBUInt := uint16(y1) - y3
	fmt.Println("SUBUInt8 & 16:", SUBUInt)

//floating point

	var z1 float32
	z1 = -666.433
	var z2 float32
	z2 = -577.8898
	SUBFloat32 := z1 - z2
	fmt.Println("SUBFloat32:", SUBFloat32)

	var z3 float64
	z3 = -666.433
	var z4 float64
	z4 = -577.8898
	SUBFloat64 := z3 - z4
	fmt.Println("SUBFloat64:", SUBFloat64)

	SUBFloat := float64(z1) - z4
	fmt.Println("SUBFloat32 & 64:", SUBFloat)
}
