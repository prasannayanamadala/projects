package game
 
import "fmt"

type Hogwarts struct {
	Name     string
	pointsscored slice int
	avg       float32
}

func GameScore() {
	
	var G1 Hogwarts
	G1.Name = "Gryffindor"

	G1.pointsscored = make([]int,0,40)

	for i:=1 ; i<=40 ;i++ {
		G1.pointsscored = append(G1.pointscsored, i)
	}
	sum := 0
	for i := 0;i < len(G1.pointsscored);i++ {
		sum = sum + G1.pointsscored[i]
	}
	fmt.Println("List of points scored by each student in group Gryffindor:",G1pointsscored)
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
	 for i:=1; i<=40 ; i=i+1 {
		 Slytherin = append(Slytherin, i)
	 }
	 
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
	
	//different way of appending values
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
 
