package main

import "fmt"

func main() {

	ns := GetNutritionalScore(NutritionalData{
		Energy:              EnergyFromKcal(500),
		Sugars:              SugarGram(20),
		SaturatedFattyAcids: SaturatedFattyAcids(5),
		Sodium:              SodiumMilligram(700),
		Fruits:              FruitsPercent(50),
		Fibre:               FibreGram(6),
		Protein:             ProteinGram(10),
	}, Food)

	fmt.Printf("Nutritional Score:%d\n", ns.Value)
	fmt.Printf("NutriScore: %s\n", ns.GetNutriScore())
}
