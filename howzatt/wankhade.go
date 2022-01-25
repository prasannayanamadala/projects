package worldcup

import "fmt"

func ScoreCard() {

	IndianPlayer := make(map[string]int)
	IndianPlayer["A"] = 10
	IndianPlayer["B"] = 20
	IndianPlayer["C"] = 30
	IndianPlayer["D"] = 40
	IndianPlayer["E"] = 50
	IndianPlayer["F"] = 60
	IndianPlayer["G"] = 70
	IndianPlayer["H"] = 80
	IndianPlayer["I"] = 90
	IndianPlayer["J"] = 100
	IndianPlayer["K"] = 110

	fmt.Println("Names and Runs of all Indian Batsmen:", IndianPlayer)

	IndianBowler := make(map[string]int)
	IndianBowler["A"] = 10
	IndianBowler["B"] = 20
	IndianBowler["C"] = 30
	IndianBowler["D"] = 40
	IndianBowler["E"] = 40
	IndianBowler["F"] = 60
	IndianBowler["G"] = 70
	IndianBowler["H"] = 70
	IndianBowler["I"] = 90
	IndianBowler["J"] = 100
	IndianBowler["K"] = 110

	fmt.Println("Names and Runs conceded by Indian Bowlwers:", IndianBowler)


	SrilankaPlayer := make(map[string]int)
	SrilankaPlayer["AA"] = 10
	SrilankaPlayer["BB"] = 20
	SrilankaPlayer["CC"] = 30
	SrilankaPlayer["DD"] = 40
	SrilankaPlayer["EE"] = 40
	SrilankaPlayer["FF"] = 60
	SrilankaPlayer["GG"] = 70
	SrilankaPlayer["HH"] = 70
	SrilankaPlayer["II"] = 90
	SrilankaPlayer["JJ"] = 100
	SrilankaPlayer["KK"] = 110

	fmt.Println("Names and Runs of all Sri Lanka Batsmen:", SrilankaPlayer)
	
	SrilankaBowler := make(map[string]int)
	SrilankaBowler["AA"] = 10
	SrilankaBowler["BB"] = 20
	SrilankaBowler["CC"] = 30
	SrilankaBowler["DD"] = 40
	SrilankaBowler["EE"] = 50
	SrilankaBowler["FF"] = 60
	SrilankaBowler["GG"] = 70
	SrilankaBowler["HH"] = 80
	SrilankaBowler["II"] = 90
	SrilankaBowler["JJ"] = 100
	SrilankaBowler["KK"] = 110

	fmt.Println("Names and Runs conceded by Sri Lanka Bowlers:", SrilankaBowler)
	fmt.Println("Points scored by Srilanka Bowler 1:",SrilankaBowler["AA"])  //accessing score of 1player

	Total := IndianPlayer["A"] +IndianPlayer["B"]+IndianPlayer["C"]+IndianPlayer["D"]+IndianPlayer["E"]+IndianPlayer["F"]+IndianPlayer["G"]+IndianPlayer["H"]+IndianPlayer["I"]+IndianPlayer["J"]+IndianPlayer["K"]
	fmt.Println("Final score of India:",Total)
	Total1 := SrilankaPlayer["AA"]+SrilankaPlayer["BB"]+SrilankaPlayer["CC"]+SrilankaPlayer["DD"]+SrilankaPlayer["EE"]+SrilankaPlayer["FF"]+SrilankaPlayer["GG"]+SrilankaPlayer["HH"]+SrilankaPlayer["II"]+SrilankaPlayer["JJ"]+SrilankaPlayer["KK"]
	fmt.Println("Final score of Srilanka:",Total1)
	
	Result := Total - Total1
	fmt.Println("Result of Match(India won by):",Result)
	 
	// using for
{
	for key, val1 := range IndianPlayer {
		fmt.Println("Final score of India:",key,val1)
		fmt.Println()

	}
	for key, val := range SrilankaPlayer {
		fmt.Println("Final score of Srilanka:",key,val)
		fmt.Println()
	}
	val1:=660
	val := 640
	if val1 > val {
		fmt.Println("Result of match-India won the game")
	} else if val1 == val {
		fmt.Println("match tie")
	} else {
		fmt.Println("Result of match-Srilanka won the game")
	}
	}
	
}

