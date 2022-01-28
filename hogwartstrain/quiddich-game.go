package game
 
import "fmt"

type Hogwarts struct {
	Name     string
	pointsscored []int
	avg       float32
}

func GameScore() {
	
	var G1 Hogwarts
	G1.Name = "Gryffindor"

	G1.pointsscored = make([]int,0,40)

	for i:=1 ; i<=40 ;i++ {
		G1.pointsscored = append(G1.pointsscored,i)
	}
	sum := 0
	for i := 0;i < len(G1.pointsscored);i++ {
		sum = sum + G1.pointsscored[i]
	}
	fmt.Println("List of points scored by each student in group Gryffindor:",G1.pointsscored)
	fmt.Println("Total points scored by group Gryffindor:",sum)

	G:= G1.pointsscored
	G1.pointsscored[8] = 13
	G1.pointsscored[26] = 31
	G1.pointsscored[14] = 19
	fmt.Println("List of points scored by each student in group Gryffindor:",G)

	sum5 := 0
	for i := 0;i < len(G1.pointsscored);i++ {
		sum5 = sum5 + G1.pointsscored[i]
	}
	fmt.Println("Total points scored by group Gryffindor:",sum5)
	
	G1.avg = float32(sum5/len(G1.pointsscored))
	fmt.Println("Average points scored by group Gryffindor:",G1.avg)
	 
		
	var G2 Hogwarts
	G2.Name = "Slytherin"
	
	G2.pointsscored = make([]int,0,40)

	 for i:=1; i<=40 ; i=i+1 {
		G2.pointsscored = append(G2.pointsscored, i)
	 }
	 
	 sum1 := 0
	 for i := 0;i < len(G2.pointsscored);i++ {
		 sum1 = sum1 + G2.pointsscored[i]
	 }
	fmt.Println("List of points scored by each student in group Slytherin:",G2.pointsscored)
	fmt.Println("Total points scored by group Slytherin:",sum1)
	
	S1 := G2.pointsscored
	G2.pointsscored[9] = 8
	G2.pointsscored[24] = 23
	G2.pointsscored[28] = 27
	fmt.Println("List of points scored by each student in group Slytherin:",S1)

	sum6 := 0
	for i := 0;i < len(G2.pointsscored);i++ {
		sum6 = sum6 + G2.pointsscored[i]
	}
	fmt.Println("Total points scored by group Slytherin:",sum6)

	G2.avg = float32(sum6/len(G2.pointsscored))
	fmt.Println("Average points scored by group Slytherin:",G2.avg)
	
	//different way of appending values
	
	var G3 Hogwarts
	
	G3.Name = "Ravenclaw"
	G3.pointsscored = make([]int,0,40)
	G3.pointsscored = append(G3.pointsscored,1,2,3,4,5,6,7,8,9,10,11,12,13,14,15,16,17,18,19,20,21,22,23,24,25,26,27,28,29,30,31,32,33,34,35,36,37,38,39,40)
	
	sum2 := 0
	for i := 0;i < len(G3.pointsscored);i++ {
		sum2 = sum2 + G3.pointsscored[i]
	}
	fmt.Println("List of points scored by each student in group Ravenclaw:",G3.pointsscored)
	fmt.Println("Total points scored by group Ravenclaw:",sum2)
	
	
	for i := 0; i < len(G3.pointsscored);i++ {
		G3.pointsscored[i] = G3.pointsscored[i]-1;
	}
	fmt.Println("List of points scored by each student in group Ravenclaw:",G3.pointsscored)

	sum7 := 0
	for i := 0;i < len(G3.pointsscored);i++ {
		sum7 = sum7 + G3.pointsscored[i]
	}
		fmt.Println("Total points scored by group Ravenclaw:",sum7)

		G3.avg = float32(sum7/len(G3.pointsscored))
		fmt.Println("Average points scored by group Ravenclaw:",G3.avg)
	
	//without using structures
	
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
	fmt.Println(G1.Name,G2.Name,G3.Name)
}
 
