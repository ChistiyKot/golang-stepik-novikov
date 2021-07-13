/*
В ходе анализа результатов переписи населения информация была сохранена в объекте типа map: https://stepik.org/lesson/345543/step/6?unit=329288
*/
package main

import "fmt"

func main() {
	groupCity := map[int][]string{
		10:   {"A", "B"}, // города с населением 10-99 тыс. человек
		100:  {"C", "D"}, // города с населением 100-999 тыс. человек
		1000: {"E", "F"}, // города с населением 1000 тыс. человек и более
	}

	cityPopulation := map[string]int{
		"A": 10,
		"B": 50,
		"C": 100,
		"D": 150,
		"E": 1000,
		"F": 1200,
	}

	// start
	cities := groupCity[100]
	for idx := range cityPopulation {
		isWrongRecord := true
		for _, current := range cities {
			if idx == current {
				isWrongRecord = false
			}
		}

		if isWrongRecord {
			delete(cityPopulation, idx)
		}
	}
	// end

	fmt.Println(groupCity, cityPopulation)
}
