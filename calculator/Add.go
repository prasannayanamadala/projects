package operations

import "fmt"

func AddOperation() {
	
	/*Int8 — [-128 : 127]
	Int16 — [-32768 : 32767]
	Int32 — [-2147483648 : 2147483647]
	Int64 — [-9223372036854775808 : 9223372036854775807]*/
	
	var x1 int8
	x1 = -125
	var x2 int8
	x2 = -55
	ADDInt8 := x1 + x2
	fmt.Println("ADDInt8:", ADDInt8)

	var x3 int16
	x3 = -31000
	var x4 int16
	x4= -28000
	ADDInt16 := x3 + x4
	fmt.Println("ADDInt16:", ADDInt16)

	var x5 int32
	x5 = 56778899
	var x6 int32
	x6 = 567687698
	ADDInt32 := x5 + x6
	fmt.Println("ADDInt32:", ADDInt32)

	var x7 int64
	x7 = 345466578789999999
	var x8 int64
	x8 = 678687908098080675
	ADDInt64 := x7 + x8
	fmt.Println("ADDInt64:", ADDInt64)

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
}
