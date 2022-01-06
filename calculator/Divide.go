package operations

import "fmt"

func DivideOperation() {
	
	/*Int8 — [-128 : 127]
	Int16 — [-32768 : 32767]
	Int32 — [-2147483648 : 2147483647]
	Int64 — [-9223372036854775808 : 9223372036854775807]*/
	
	var x1 int8
	x1 = -12
	var x2 int8
	x2 = 6
	DIVInt8 := x1 / x2
	fmt.Println("DIVInt8:", DIVInt8)

	var x3 int16
	x3 = 256
	var x4 int16
	x4= 128
	DIVInt16 := x3 / x4
	fmt.Println("DIVInt16:", DIVInt16)

	var x5 int32
	x5 = 567
	var x6 int32
	x6 = 567
	DIVInt32 := x5 / x6
	fmt.Println("DIVInt32:", DIVInt32)

	var x7 int64
	x7 = 1024
	var x8 int64
	x8 = 256
	DIVInt64 := x7 / x8
	fmt.Println("DIVInt64:", DIVInt64)


	/*  uint8	     0 to 255	                     
		uint16	     0 to 65,535	                  
		uint32	     0 to 4,294,967,295	               
		uint64	     0 to 18,446,744,073,709,551,615	*/

	var y1 uint8
	y1 = 25
	var y2 uint8
	y2 = 25
	DIVUInt8 := y1 / y2
	fmt.Println("DIVUInt8:", DIVUInt8)

	var y3 uint16
	y3 = 45
	var y4 uint16
	y4 = 9
	DIVUInt16 := y3 / y4
	fmt.Println("DIVUInt16:", DIVUInt16)

	var y5 uint32
	y5 = 657
	var y6 uint32
	y6 = 657
	DIVUInt32 := y5 / y6
	fmt.Println("DIVUInt32:", DIVUInt32)

	var y7 uint64
	y7 = 57658
	var y8 uint64
	y8 = 56758
	DIVUInt64 := y7 / y8
	fmt.Println("DIVUInt64:", DIVUInt64)



//floating point

	var z1 float32
	z1 = -8.8
	var z2 float32
	z2 = 4
	DIVFloat32 := float32(z1) / float32(z2)
	fmt.Println("DIVFloat32:", DIVFloat32)

	var z3 float64
	z3 = 3576.56758
	var z4 float64
	z4 = 10
	DIVFloat64 := float64(z3) / float64(z4)
	fmt.Println("DIVFloat64:", DIVFloat64)

	DIVFloat := float64(z1) / z4
	fmt.Println("DIVFloat32 & 64:", DIVFloat)
}
