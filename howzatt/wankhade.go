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

	fmt.Println("Names and Runs conceded by Indian Bowlwers:", SrilankaPlayer)
	fmt.Println("Names and Runs conceded by Sri Lanka Bowlers:", IndianPlayer)
	//fmt.Println("Points scored by Srilanka player 1:",SrilankaPlayer["AA"])  //accessing score of 1player

	Total := IndianPlayer["A"] +IndianPlayer["B"]+IndianPlayer["C"]+IndianPlayer["D"]+IndianPlayer["E"]+IndianPlayer["F"]+IndianPlayer["G"]+IndianPlayer["H"]+IndianPlayer["I"]+IndianPlayer["J"]+IndianPlayer["K"]
	fmt.Println("Final score of India:",Total)
	Total1 := SrilankaPlayer["AA"]+SrilankaPlayer["BB"]+SrilankaPlayer["CC"]+SrilankaPlayer["DD"]+SrilankaPlayer["EE"]+SrilankaPlayer["FF"]+SrilankaPlayer["GG"]+SrilankaPlayer["HH"]+SrilankaPlayer["II"]+SrilankaPlayer["JJ"]+SrilankaPlayer["KK"]
	fmt.Println("Final score of Srilanka:",Total1)
	
	Result := Total - Total1
	fmt.Println("Result of Match(India won by):",Result)
	 
	// using for
{
	for key, val1 := range IndianPlayer {
		fmt.Println("score of Indianplayers:",key,val1)
		fmt.Println()

	}
	for key, val := range SrilankaPlayer {
		fmt.Println("score of Srilankaplayers:",key,val)
		fmt.Println()
	}
	val1:=660
	val := 640
	if val1 > val {
		fmt.Println("Result of match-India won the game")
	} else if val1 == val {
		fmt.Println("match tie")
	} else if val1 <val {
		fmt.Println("Result of match-Srilanka won the game")
	} else {
		fmt.Println("match cancelled")
	}
	}
	
}

