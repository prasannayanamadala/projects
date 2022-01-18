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

	India := 0
	for i := 0;i < len(IndianPlayer);i++ {
		India = India+ IndianPlayer[i]
	}
	fmt.Println("Total points scored by :",India)

}