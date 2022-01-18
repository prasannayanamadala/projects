package game
 
import "fmt"

func GameScore() {

	Gryffindor := make([]int,0,40)
	Gryffindor = append(Gryffindor,1)
	Gryffindor = append(Gryffindor,2)
	Gryffindor = append(Gryffindor,3)
	Gryffindor = append(Gryffindor,4)
	Gryffindor = append(Gryffindor,5)
	Gryffindor = append(Gryffindor,6)
	Gryffindor = append(Gryffindor,7)
	Gryffindor = append(Gryffindor,8)
	Gryffindor = append(Gryffindor,9)
	Gryffindor = append(Gryffindor,10)
	Gryffindor = append(Gryffindor,11)
	Gryffindor = append(Gryffindor,12)
	Gryffindor = append(Gryffindor,13)
	Gryffindor = append(Gryffindor,14)
	Gryffindor = append(Gryffindor,15)
	Gryffindor = append(Gryffindor,16)
	Gryffindor = append(Gryffindor,17)
	Gryffindor = append(Gryffindor,18)
	Gryffindor = append(Gryffindor,19)
	Gryffindor = append(Gryffindor,20)
	Gryffindor = append(Gryffindor,21)
	Gryffindor = append(Gryffindor,22)
	Gryffindor = append(Gryffindor,23)
	Gryffindor = append(Gryffindor,24)
	Gryffindor = append(Gryffindor,25)
	Gryffindor = append(Gryffindor,26)
	Gryffindor = append(Gryffindor,27)
	Gryffindor = append(Gryffindor,28)
	Gryffindor = append(Gryffindor,29)
	Gryffindor = append(Gryffindor,30)
	Gryffindor = append(Gryffindor,31)
	Gryffindor = append(Gryffindor,32)
	Gryffindor = append(Gryffindor,33)
	Gryffindor = append(Gryffindor,34)
	Gryffindor = append(Gryffindor,35)
	Gryffindor = append(Gryffindor,36)
	Gryffindor = append(Gryffindor,37)
	Gryffindor = append(Gryffindor,38)
	Gryffindor = append(Gryffindor,39)
	Gryffindor = append(Gryffindor,40)
	sum := 0
	for i := 0;i < len(Gryffindor);i++ {
		sum = sum + Gryffindor[i]
	}
	
	fmt.Println("List of points scored by each student in group Gryffindor:",Gryffindor)
	fmt.Println("Total points scored by group Gryffindor:",sum)

	Gryffindor1 := Gryffindor
	Gryffindor[8] = 13
	Gryffindor[26] = 31
	Gryffindor[14] = 19
	fmt.Println("List of points scored by each student in group Gryffindor:",Gryffindor1)

	sum5 := 0
	for i := 0;i < len(Gryffindor);i++ {
		sum5 = sum5 + Gryffindor[i]
	}
	fmt.Println("Total points scored by group Gryffindor:",sum5)
	
	avg1 := sum5/ len(Gryffindor)
	fmt.Println("Average points scored by group Gryffindor:",avg1)
	 
	Slytherin := make([]int,0,40)
	 Slytherin = append(Slytherin,1)
	 Slytherin = append(Slytherin,2)
	 Slytherin = append(Slytherin,3)
	 Slytherin = append(Slytherin,4)
	 Slytherin = append(Slytherin,5)
	 Slytherin = append(Slytherin,6)
	 Slytherin = append(Slytherin,7)
	 Slytherin = append(Slytherin,8)
	 Slytherin = append(Slytherin,9)
	 Slytherin = append(Slytherin,10)
	 Slytherin = append(Slytherin,11)
	 Slytherin = append(Slytherin,12)
	 Slytherin = append(Slytherin,13)
	 Slytherin = append(Slytherin,14)
	 Slytherin = append(Slytherin,15)
	 Slytherin = append(Slytherin,16)
	 Slytherin = append(Slytherin,17)
	 Slytherin = append(Slytherin,18)
	 Slytherin = append(Slytherin,19)
	 Slytherin = append(Slytherin,20)
	 Slytherin = append(Slytherin,21)
	 Slytherin = append(Slytherin,22)
	 Slytherin = append(Slytherin,23)
	 Slytherin = append(Slytherin,24)
	 Slytherin = append(Slytherin,25)
	 Slytherin = append(Slytherin,26)
	 Slytherin = append(Slytherin,27)
	 Slytherin = append(Slytherin,28)
	 Slytherin = append(Slytherin,29)
	 Slytherin = append(Slytherin,30)
	 Slytherin = append(Slytherin,31)
	 Slytherin = append(Slytherin,32)
	 Slytherin = append(Slytherin,33)
	 Slytherin = append(Slytherin,34)
	 Slytherin = append(Slytherin,35)
	 Slytherin = append(Slytherin,36)
	 Slytherin = append(Slytherin,37)
	 Slytherin = append(Slytherin,38)
	 Slytherin = append(Slytherin,39)
	 Slytherin = append(Slytherin,40)
	 
	 sum1 := 0
	 for i := 0;i < len(Slytherin);i++ {
		 sum1 = sum1 + Slytherin[i]
	 }
	fmt.Println("List of points scored by each student in group Slytherin:",Slytherin)
	fmt.Println("Total points scored by group Slytherin:",sum1)

	Slytherin1 := Slytherin
	Slytherin[9] = 8
	Slytherin[24] = 23
	Slytherin[28] = 27
	fmt.Println("List of points scored by each student in group Slytherin:",Slytherin1)

	sum6 := 0
	for i := 0;i < len(Slytherin);i++ {
		sum6 = sum6 + Slytherin[i]
	}
	fmt.Println("Total points scored by group Slytherin:",sum6)

	avg2 := sum6/len(Slytherin)
	fmt.Println("Average points scored by group Slytherin:",avg2)
 
	Ravenclaw := make([]int,0,40)
	Ravenclaw = append(Ravenclaw,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40)
	
	sum2 := 0
	for i := 0;i < len(Ravenclaw);i++ {
		sum2 = sum2 + Ravenclaw[i]
	}
	fmt.Println("List of points scored by each student in group Ravenclaw:",Ravenclaw)
	fmt.Println("Total points scored by group Ravenclaw:",sum2)
	
	
	for i := 0; i < len(Ravenclaw);i++ {
		Ravenclaw[i] = Ravenclaw[i]-1;
	}
	fmt.Println("List of points scored by each student in group Ravenclaw:",Ravenclaw)

	sum7 := 0
	for i := 0;i < len(Ravenclaw);i++ {
		sum7 = sum7 + Ravenclaw[i]
	}
		fmt.Println("Total points scored by group Ravenclaw:",sum7)

		avg3 := sum7/len(Ravenclaw)
		fmt.Println("Average points scored by group Ravenclaw:",avg3)
	
	Hufflepuff := make([]int,0,40)
	Hufflepuff = append(Hufflepuff,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,1,2,3,29,35,36,37,38,39,40)
	
	sum3 := 0
	for i := 0;i < len(Hufflepuff);i++ {
		sum3 = sum3 + Hufflepuff[i]
	}
	fmt.Println("List of points scored by each student in group Hufflepuff:",Hufflepuff)
	fmt.Println("Total points scored by group Hufflepuff:",sum3)

	for i := 0; i < len(Hufflepuff);i++ {
		Hufflepuff[i] = Hufflepuff[i]+2;
	}
	fmt.Println("List of points scored by each student in group Hufflepuff:",Hufflepuff)
	
	sum8 := 0
	for i := 0;i < len(Hufflepuff);i++ {
		sum8 = sum8 + Hufflepuff[i]
	}
	fmt.Println("Total points scored by group Hufflepuff:",sum8)
	
	avg4 := sum8/len(Hufflepuff)
	fmt.Println("Average points scored by group Hufflepuff:",avg4)
}
 
